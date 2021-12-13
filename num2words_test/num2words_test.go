package num2words_test

import (
	"fmt"
	"goconvert/num2words"
	"math/big"
	"testing"
)

func TestNum2Words(t *testing.T) {
	newnum := num2words.New
	
	// test the uniques
	t.Run("0", tester(newnum(0), false, "zero"))
	t.Run("1", tester(newnum(1), false, "one"))
	t.Run("7", tester(newnum(7), false, "seven"))
	t.Run("-10", tester(newnum(-10), false, "negative ten"))
	t.Run("13", tester(newnum(13), false, "thirteen"))
	t.Run("19", tester(newnum(19), false, "nineteen"))

	// test hundreds
	t.Run("123", tester(newnum(123), false, "one hundred twenty three"))
	t.Run("420", tester(newnum(420), false, "four hundred twenty"))
	t.Run("810", tester(newnum(810), false, "eight hundred ten"))

	// test more
	t.Run("1910", tester(newnum(1910), false, "one thousand nine hundred ten"))
	t.Run("123456789", tester(newnum(123456789), false, "one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine"))
	t.Run("1000000000000000000", tester(newnum(1000000000000000000), false, "one quintillion"))
}

func TestNum2WordsWithSep(t *testing.T) {
	newnum := num2words.New
	
	// test the uniques
	t.Run("0", tester(newnum(0), true, "zero"))
	t.Run("1", tester(newnum(1), true, "one"))
	t.Run("7", tester(newnum(7), true, "seven"))
	t.Run("-10", tester(newnum(-10), true, "negative ten"))
	t.Run("13", tester(newnum(13), true, "thirteen"))
	t.Run("19", tester(newnum(19), true, "nineteen"))

	// test hundreds
	t.Run("123", tester(newnum(123), true, "one hundred and twenty-three"))
	t.Run("420", tester(newnum(420), true, "four hundred and twenty"))
	t.Run("810", tester(newnum(810), true, "eight hundred and ten"))

	// test more
	t.Run("1910", tester(newnum(1910), true, "one thousand and nine hundred and ten"))
	t.Run("123456789", tester(newnum(123456789), true, "one hundred and twenty-three million and four hundred and fifty-six thousand and seven hundred and eighty-nine"))
	t.Run("1000000000000000000", tester(newnum(1000000000000000000), true, "one quintillion"))
}

func tester(num *big.Int, f bool, expected string) func(*testing.T) {
	n2w := num2words.Num2Words
	return func(t *testing.T) {
		actual := n2w(num, f)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected: %s, got: %s", expected, actual))
		}
	}
}