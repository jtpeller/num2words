// ============================================================================
// = bignum.go																  =
// = 	Description: VERY helpful wrapper functions for big.Int				  =
// = 	Date: December 09, 2021												  =
// ============================================================================

package num2words

import (
	"math/big"

	gb "github.com/jtpeller/gobig"
)

// INITS
func zero() *big.Int {
	return gb.Zero()
}

func newint(i int64) *big.Int {
	return gb.New(i)
}

func copy(a *big.Int) *big.Int {
	return add(zero(), a)
}

// COMPARISONS
func equals(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}

func less(a, b *big.Int) bool {
	return a.Cmp(b) == -1
}

func greater(a, b *big.Int) bool {
	return a.Cmp(b) == 1
}

func lessEqual(a, b *big.Int) bool {
	return less(a, b) || equals(a, b)
}

func greaterEqual(a, b *big.Int) bool {
	return greater(a, b) || equals(a, b)
}

// MATHEMATICAL OPERATIONS
func abs(a *big.Int) *big.Int {
	return zero().Abs(a)
}

func add(a, b *big.Int) *big.Int {
	return zero().Add(a, b)
}

func div(a, b *big.Int) *big.Int {
	return zero().Div(a, b)
}

func mod(a, b *big.Int) *big.Int {
	return zero().Mod(a, b)
}

func pow(a *big.Int, e *big.Int) *big.Int {
	return zero().Exp(a, e, big.NewInt(0))
}
