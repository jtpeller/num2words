// ============================================================================
// = num2words.go															  =
// = 	Description: Converts a number to its English counterpart			  =
// = 			This is unique because it allows arbitrary precision numbers  =
// =	Based on: https://github.com/divan/num2words						  =
// = 	Date: December 12, 2021												  =
// ============================================================================

package num2words

import (
	"fmt"
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

// centillion is 10^303, need to manage the gap b/w vigintillion & this
var final = "centillion"				// nothing defined after this

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
	if GreaterEqual(num, Zero()) && LessEqual(num, New(19)) {
		return unique[num.Int64()]
	} else if GreaterEqual(num, New(-19)) && LessEqual(num, Zero()) {
		return "negative " + unique[Abs(num).Int64()]
	}
	pos := Abs(num)

	// count up how many 3-digit groups there are
	ndigits := countDigits(pos)
	ngroups := Div(ndigits, New(3)).Int64() + 1
	fmt.Println(ndigits, ngroups)

	// separate into 3-digit groups
	groups := make([]int64, 0)		// 3-digit groups won't need big.Int
	for i := int64(0); i < ngroups; i++ {
		groups = append(groups, Mod(pos, New(1000)).Int64())
		pos.Div(pos, New(1000))
	}
	fmt.Println(groups)

	// convert each of the groups
	strgroups := make([]string, len(groups))
	for i := 0; i < len(groups); i++ {
		strgroups[i] = convertGroup(groups[i], sep)
	}
	fmt.Println(strgroups)

	// assemble (and handle larger denominations than hundreds)
	words := strgroups[0]
	for i := int64(1); i < ngroups; i++ {
		if groups[i] != 0 {
			prefix := strgroups[i] + " " + more[i]
			if len(words) != 0 {
				prefix += addand(sep)
			}
			words = prefix + words
		}
	}
	
	// set sign
	if num.Sign() == -1 {
		words = "negative " + words
	}

	return words
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
	a := Copy(num)		// create deep copy of number so we don't modify the original
	count := New(0)
	for !Equal(a, New(0)) {	// int division guaranteed to hit zero
		a = Div(a, New(10))
		count = Add(count, New(1))
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