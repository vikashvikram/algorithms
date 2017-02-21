package algorithms

import (
	"os"
	"fmt"
	"strings"
)

func Reverse(str string) string {
	str_arr := strings.Split(str, " ")
	len := len(str_arr) - 1
	for i := 0; i <= len/2; i++ {
		str_arr[i], str_arr[len-i] = str_arr[len-i], str_arr[i]
	}
	return strings.Join(str_arr, " ")
}
