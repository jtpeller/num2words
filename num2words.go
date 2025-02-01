// ============================================================================
// = num2words.go
// = 	Description: Converts a number to its English counterpart
// = 			Allows arbitrary precision numbers
// =	Based on: https://github.com/divan/num2words
// = 	Date: December 12, 2021
// ============================================================================

package num2words

import (
	"math/big"
)

var unique = []string{
	"zero", "one", "two", "three", "four", "five",
	"six", "seven", "eight", "nine", "ten", 
	"eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen",
}

var tens_words = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
}

// 64-bit ints go up to quintillion, big.Int is arbitrarily large
// IMPORTANT: this uses the short-scale system for large numbers
var more = []string{
	"hundred", "thousand", "million", "billion", "trillion", 
	"quadrillion", "quintillion", "sextillion", "septillion",
	"octillion", "nonillion",	"decillion", "undecillion",
	"duodecillion", "tredecillion", "quattuordecillion",
	"quindecillion", "sexdecillion", "septendecillion", 
	"octodecillion", "novemdecillion",
	"vigintillion",		// vigintillion is 10^63
}

// cent is 10^303, need to manage the gap b/w vigintillion & this
var vigin = "vigintillion"			// nothing defined b/w this & cent
var cent = "centillion"				// nothing defined after this

// 1,234,567,890,123,456,789

/**
 * Num2Words converts a word to its English counterpart.
 *  For instance, 6 converts to six. 
 *  235 can convert to two hundred thirty five if sep is false
 *  or two hundred and thirty-five if sep is true.
 * @param num 	number to convert
 * @param sep 	whether to include separators like 'and' or hyphens
 * @return 		string representation of num
 */
func Num2Words(num *big.Int, sep bool) string {
	// base cases
	if greaterEqual(num, zero()) && lessEqual(num, newint(19)) {
		return unique[num.Int64()]
	} else if greaterEqual(num, newint(-19)) && lessEqual(num, zero()) {
		return "negative " + unique[abs(num).Int64()]
	}

	// init
	ret := ""						// string to return
	ndigits := countDigits(num)		// # of digits

	// check for very large numbers
	if greater(ndigits, newint(303)) {			// num > centillion
		// break up the number
		divisor := pow(newint(10), newint(303))
		left := div(num, divisor)			// left most digits
		right := mod(num, divisor)			// right most digits < 10^303

		// if right is zero, then Num2Words(right) is just vigintillion
		if equals(right, zero()) {
			ret += Num2Words(left, sep) + " " + cent
		} else {
			ret += Num2Words(left, sep) + " " + cent + addcomma(sep) + Num2Words(right, sep)
		}
	} else if greater(ndigits, newint(63)) {	// num > vigintillion
		// break up the number
		divisor := pow(newint(10), newint(63))
		left := div(num, divisor)			// left most digits
		right := abs(mod(num, divisor))			// right most digits

		// if right is zero, then Num2Words(right) is just vigintillion
		if equals(right, zero()) {
			ret += Num2Words(left, sep) + " " + vigin
		} else {
			ret += Num2Words(left, sep) + " " + vigin + addcomma(sep) + Num2Words(right, sep)
		}
	} else {				// num < vigintillion
		ret += convertNum(num, sep)
	}

	return ret
}

// used to convert nums < 10^303 (i.e. centillion or below)
func convertNum(num *big.Int, sep bool) string {
	// inits
	pos := abs(num)
	ndigits := countDigits(num)
	ngroups := div(ndigits, newint(3)).Int64() + 1

	// separate into 3-digit groups
	groups := make([]int64, 0)		// 3-digit groups won't need big.Int
	for i := int64(0); i < ngroups; i++ {
		groups = append(groups, mod(pos, newint(1000)).Int64())
		pos.Div(pos, newint(1000))
	}

	// convert each of the groups
	strgroups := make([]string, len(groups))
	for i := 0; i < len(groups); i++ {
		strgroups[i] = convertGroup(groups[i], sep)
	}

	// assemble (and handle larger denominations than hundreds)
	str := strgroups[0]
	for i := int64(1); i < ngroups; i++ {
		if groups[i] != 0 {
			prefix := ""
			// num is centillion or more (although, it's limited anyway)
			if i >= 101 {
				prefix = strgroups[i] + " " + cent
			} else if i >= 22 {	// vigintillion < num < centillion
				prefix = strgroups[i] + " " + vigin
			} else {
				prefix = strgroups[i] + " " + more[i]
			}
			if len(str) != 0 {
				prefix += addcomma(sep)
			}
			str = prefix + str
		}
	}
	
	// set sign
	if num.Sign() == -1 {
		str = "negative " + str
	}

	return str
}

// this does the "hundreds" of whatever denom we're in
// so, if i've got 324, it would return three hundred [and] twenty[-]four
// and this would denote how many thousands or millions or whatever
func convertGroup(group int64, sep bool) string {
	// init
	hundreds := group / 100
	rem := group % 100
	out := ""			// string to return

	// convert hundreds to words
	if hundreds != 0 {
		out += unique[hundreds] + " " + more[0]
		if rem != 0 {
			out += addand(sep)
		}
	}

	// compute tens to words
	tens := rem / 10	// get how many tens
	ones := rem % 10
	if tens >= 2 {
		out += tens_words[tens]
		if ones != 0 {
			out += adddash(sep) + unique[ones]
		}
	} else if rem != 0 {
		out += unique[rem]
	}

	return out
}

func countDigits(num *big.Int) *big.Int {
	a := copy(num)		// create deep copy of number so we don't modify the original
	count := newint(0)
	for !equals(a, newint(0)) {	// int division guaranteed to hit zero
		a = div(a, newint(10))
		count = add(count, newint(1))
	}
	return count
}

func addand(sep bool) string {
	if sep {
		return " and "
	}
	return " "
}

func adddash(sep bool) string {
	if sep {
		return "-"
	}
	return " "
}

func addcomma(sep bool) string {
	if sep {
		return ", "
	}
	return " "
}
