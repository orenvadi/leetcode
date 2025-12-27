package main

import "strings"

func addBinary(a, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	var builder strings.Builder
	builder.Grow(len(a) + 1)

	i, j := len(a)-1, len(b)-1
	carry := 0

	for i >= 0 {
		sum := carry
		sum += int(a[i] - '0')
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		carry = sum / 2
		builder.WriteByte('0' + byte(sum%2))
		i--
	}
	if carry > 0 {
		builder.WriteByte('1')
	}

	s := builder.String()
	// Reverse the string
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
