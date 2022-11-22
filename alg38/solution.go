package alg38

import (
	"strconv"
)

func countAndSay(n int) string {
    
	if n == 1 {
		return "1"
	}

	if (n <= 30) && (n > 1) {
		
		result := ""

		say := countAndSay(n - 1)

		length := len(say)

        unique := say[0:1]
		count := 1

		for i := 1; i < length; i++ {
			if say[i:i+1] != unique {
				result += strconv.Itoa(count) + unique
                
				unique = say[i:i+1]
				count = 1
			} else {
				count += 1
			}
		}
		
		result += strconv.Itoa(count) + unique

		return result
	}

	panic(0)
}

func Solution(n int) string {
	return countAndSay(n)
}
