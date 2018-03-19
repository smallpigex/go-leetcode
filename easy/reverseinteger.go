package main

func main() {
	reverse(1234)
}

func reverse(x int) int {

	var total int = 0
	count := false

	for x != 0 {
		t := x % 10
		if count {
			total = t
			count = true
		} else {
			total *= 10
			total += t
		}

		x /= 10
	}

	result := int(total)

	if result > -2147483648 && result < 2147483647 {
		return result
	}
	return 0
}
