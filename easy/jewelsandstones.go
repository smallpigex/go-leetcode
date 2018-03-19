package main

func numJewelsInStones(J string, S string) int {
	numOfjewel := 0

	for _, jewel := range J {
		for _, stone := range S {
			if stone == jewel {
				numOfjewel++
			}
		}
	}

	return numOfjewel
}
