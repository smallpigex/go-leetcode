package main

func numJewelsInStones(J string, S string) int {
	numOfJewels := 0

	for _, jewel := range J {
		for _, stone := range S {
			if stone == jewel {
				numOfJewels++
			}
		}
	}

	return numOfJewels
}
