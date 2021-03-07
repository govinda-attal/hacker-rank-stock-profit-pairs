package main

import "fmt"

func profitPairs(sp []int32, target int32) (total int) {
	var pos int
	var pairs [][]int32
	for {
		if pos == len(sp) {
			break
		}
		var i, j int
		var paired bool
		for i = pos; i < len(sp); i++ {

			for j = i + 1; j < len(sp); j++ {
				if sp[i]+sp[j] == target {
					total++
					paired = true
					pairs = append(pairs, []int32{sp[i], sp[j]})
					break
				}
			}
			if paired {
				break
			}
		}
		if paired {
			sp = append(sp[:i], sp[i+1:]...)
			sp = append(sp[:j-1], sp[j:]...)
		} else {
			pos++
		}
	}
	fmt.Println(pairs)
	fmt.Println(total)
	return total
}

func main() {
	fmt.Println(profitPairs([]int32{1, 2, 2, 2, 3, 4, 4, 4}, 5) == 2)
	fmt.Println(profitPairs([]int32{1, 2, 3, 6, 7, 8, 9, 1}, 10) == 3)
	fmt.Println(profitPairs([]int32{3, 2, 1, 45, 27, 6, 78, 9, 0}, 9) == 2)
}
