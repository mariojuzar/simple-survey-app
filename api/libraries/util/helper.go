package util

import "strconv"

func StrToUint(word string) uint32 {

	num, _ := strconv.Atoi(word)
	return uint32(num)
}