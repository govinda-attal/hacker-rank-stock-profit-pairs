package main

import "fmt"

func profitPairs(sp []int32, target int32) (total int) {
	var pos int
	var pairs [][]int32
	for {
		if pos == len(sp) {
			// looped over every element in the array
			// exit
			break
		}
		var i, j int
		var paired bool
	Loop:
		for i = pos; i < len(sp); i++ {
			for j = i + 1; j < len(sp); j++ {
				if sp[i]+sp[j] == target {
					paired = true
					break Loop
				}
			}
		}
		if paired {
			pairs = append(pairs, []int32{sp[i], sp[j]})
			// remove the paired profits from the array
			sp = append(sp[:j], sp[j+1:]...)
			sp = append(sp[:i], sp[i+1:]...)
			total++
		} else {
			pos++
		}
	}
	// fmt.Println(pairs)
	// fmt.Println(total)
	return total
}

func main() {
	fmt.Println(profitPairs([]int32{1, 2, 2, 2, 3, 4, 4, 4}, 5) == 2)
	fmt.Println(profitPairs([]int32{1, 2, 3, 6, 7, 8, 9, 1}, 10) == 3)
	fmt.Println(profitPairs([]int32{3, 2, 1, 45, 27, 6, 78, 9, 0}, 9) == 2)
	fmt.Println(profitPairs([]int32{3, 3, 2, 1, 45, 27, 6, 78, 9, 0}, 9) == 2)
	fmt.Println(profitPairs([]int32{1, 5, 66, 2, 3, 4, 7, 0, 2, 5}, 7) == 4)
	fmt.Println(profitPairs([]int32{}, 5) == 0)
	fmt.Println(profitPairs([]int32{5}, 5) == 0)
	fmt.Println(profitPairs([]int32{3, 4}, 7) == 1)
}
