package tests

import (
	"fmt"
	"strings"
)

func alreadyInArray(array []string, content string) bool {
	for _, item := range array {
		return content == item
	}

	return false
}

func IsStringInArray() {
	var searched []string
	var a string = "aa asd"
	var b []string = []string{"aab aasdwd", "bbaaa raat", "ccc yyy"}

	var listA []string = strings.Split(a, " ")
	for _, word := range listA {
		for _, val := range b {
			if strings.Contains(val, word) && !alreadyInArray(searched, val) {
				searched = append(searched, val)
			}
		}
	}

	for _, v := range searched {
		fmt.Println(v)
	}
}
