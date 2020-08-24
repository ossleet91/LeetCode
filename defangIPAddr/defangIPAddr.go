package main

// Defanging an IP address
//
// Given a valid (IPv4) IP address, return a defanged version of that IP address.
//
// A defanged IP address replaces every period "." with "[.]".
//
// Example:
//	Input: address = "1.1.1.1"
//	Output: "1[.]1[.]1[.]1"
//
// https://leetcode.com/problems/defanging-an-ip-address/

import (
	"strings"
)

func defangIPAddr(address string) string {
	s := strings.FieldsFunc(address,
		func(c rune) bool {
			return c == '.'
		})

	return strings.Join(s, "[.]")
}

func main() {
}
