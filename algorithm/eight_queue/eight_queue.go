package eight_queue

import "fmt"

var result = [8]int{} // 下标表示行, 值表示列

func cal8queues(row int) {
	if row == 8 {
		fmt.Println("agagin")
		return
	}

	for column := 0; column < 8; column++ {
		if isOK(row, column) {
			result[row] = column
			cal8queues(row + 1)
		}
	}
}

func isOK(row, column int) bool {
	leftup, rightup := column-1, column+1

	for row = row - 1; row >= 0; row-- {
		if result[row] == column {
			return false
		}

		if leftup >= 0 && result[row] == leftup {
			return false
		}
		if rightup < 8 && result[row] == rightup {
			return false
		}
		leftup--
		rightup++
	}
	return true
}

func printQueues(arr [8]int) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if j == arr[i] {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

}
