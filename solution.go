package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	iterator := shortestworld(strs)
	
	res := ""
	for i := 0; i < iterator; i++ {
		actullchar := strs[0][i]

		for _, val := range strs {
			if val[i] != actullchar {
				return res
			}
		}
		res += string(actullchar)
	}

	return res
}

func shortestworld(strs []string) int {
	lenghtofword := len(strs[0])

	for _, val := range strs {
		fmt.Println(val)

		newlenght := len(val)
		if lenghtofword > newlenght {
			lenghtofword = newlenght
		}
	}

	return lenghtofword
}
