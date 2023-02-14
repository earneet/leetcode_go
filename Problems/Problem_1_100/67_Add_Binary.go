package Problem_1_100

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	res := make([]string, 0)
	carry := 0
	maxLen := 0
	if len(a) > len(b) {
		maxLen = len(a)
	} else {
		maxLen = len(b)
	}

	for i := 0; i < maxLen; i++ {
		adderA := 0
		adderB := 0
		if i < len(a) {
			adderA, _ = strconv.Atoi(string(a[len(a)-(i+1)]))
		}
		if i < len(b) {
			adderB, _ = strconv.Atoi(string(b[len(b)-(i+1)]))
		}
		digit := adderA + adderB + carry
		carry = digit / 2
		digit = digit % 2
		res = append(res, strconv.Itoa(digit))
	}
	if carry > 0 {
		res = append(res, "1")
	}
	return strings.Join(reverse(res), "")
}

func reverse(arr []string) []string {
	mid := len(arr) / 2
	for i := 0; i < mid; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	return arr
}
