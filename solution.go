package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := ""

	iterator := shortestworld(strs)
	fmt.Println(iterator)

	for _, val := range strs {
		matchall := ""

		res = matchall
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
