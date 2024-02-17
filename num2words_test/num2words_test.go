package num2words_test

import (
	"fmt"
	"math/big"
	n2w "num2words"
	"testing"

	gb "github.com/jtpeller/gobig"
)

// runs test without a separator
func TestNum2Words(t *testing.T) {
	// test the uniques
	t.Run("0", testn2w(gb.New(0), false, "zero"))
	t.Run("1", testn2w(gb.New(1), false, "one"))
	t.Run("7", testn2w(gb.New(7), false, "seven"))
	t.Run("-10", testn2w(gb.New(-10), false, "negative ten"))
	t.Run("13", testn2w(gb.New(13), false, "thirteen"))
	t.Run("19", testn2w(gb.New(19), false, "nineteen"))

	// test hundreds
	t.Run("123", testn2w(gb.New(123), false, "one hundred twenty three"))
	t.Run("420", testn2w(gb.New(420), false, "four hundred twenty"))
	t.Run("810", testn2w(gb.New(810), false, "eight hundred ten"))

	// test up through int64
	t.Run("1910", testn2w(gb.New(1910), false, "one thousand nine hundred ten"))
	t.Run("123456789", testn2w(gb.New(123456789), false, "one hundred twenty three million four hundred fifty six thousand seven hundred eighty nine"))
	t.Run("1000000000000000000", testn2w(gb.New(1000000000000000000), false, "one quintillion"))

	// test massive numbers
	t.Run("10^25", testn2w(gb.Pow(gb.New(10), gb.New(25)), false, "ten septillion"))
	t.Run("10^64", testn2w(gb.Pow(gb.New(10), gb.New(64)), false, "ten vigintillion"))
}

// runs tests with a separator
func TestNum2WordsWithSep(t *testing.T) {
	// test the uniques
	t.Run("0", testn2w(gb.New(0), true, "zero"))
	t.Run("1", testn2w(gb.New(1), true, "one"))
	t.Run("7", testn2w(gb.New(7), true, "seven"))
	t.Run("-10", testn2w(gb.New(-10), true, "negative ten"))
	t.Run("13", testn2w(gb.New(13), true, "thirteen"))
	t.Run("19", testn2w(gb.New(19), true, "nineteen"))

	// test hundreds
	t.Run("123", testn2w(gb.New(123), true, "one hundred and twenty-three"))
	t.Run("420", testn2w(gb.New(420), true, "four hundred and twenty"))
	t.Run("810", testn2w(gb.New(810), true, "eight hundred and ten"))

	// test more
	t.Run("1910", testn2w(gb.New(1910), true, "one thousand, nine hundred and ten"))
	t.Run("123456789", testn2w(gb.New(123456789), true, "one hundred and twenty-three million, four hundred and fifty-six thousand, seven hundred and eighty-nine"))
	t.Run("1000000000000000000", testn2w(gb.New(1000000000000000000), true, "one quintillion"))

	// test massive numbers
	t.Run("10^25", testn2w(gb.Pow(gb.New(10), gb.New(25)), true, "ten septillion"))
	t.Run("10^64", testn2w(gb.Pow(gb.New(10), gb.New(64)), true, "ten vigintillion"))
}

// runs the test and handles the call
func testn2w(num *big.Int, f bool, expected string) func(*testing.T) {
	n2w := n2w.Num2Words
	return func(t *testing.T) {
		actual := n2w(num, f)
		if actual != expected {
			t.Errorf(fmt.Sprintf("Expected: %s, got: %s", expected, actual))
		}
	}
}

