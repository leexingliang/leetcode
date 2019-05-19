package coins

import (
	"fmt"
	"math"
	"testing"
)

func Test_CalcCoins(t *testing.T) {
	calcCoins(coins, 0, 0)
	fmt.Println(num)
}

func Test_CalcCoinsMem(t *testing.T) {
	calcCoinsMem(coins, 0, 0)
	fmt.Println(num)
}

func Test_CalcCoinsMem_2(t *testing.T) {
	num = math.MaxInt32
	calcCoinsMem([]int{1}, 0, 0)
	fmt.Println(num)
}

func Test_calcCoinsF(t *testing.T) {
	result := calcCoinsF([]int{1})
	fmt.Println(result)

	result = calcCoinsF(coins)
	fmt.Println(result)
}
