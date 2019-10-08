package plus_one

func plusOne(digits []int) []int {
	if len(digits) <= 0 {
		return nil
	}

	end := len(digits) - 1
	inc := 1
	for end >= 0 {
		tmp := digits[end] + inc
		if tmp >= 10 {
			inc = tmp - digits[end]
			digits[end] = tmp - 10
		} else {
			digits[end] = tmp
			inc = 0
			break
		}
		end--
	}
	if inc != 0 {
		result := make([]int, 0, len(digits)+1)
		result = append(result, inc)
		result = append(result, digits...)
		return result
	}
	return digits
}
