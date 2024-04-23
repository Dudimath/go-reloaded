package go_reloaded

import (
	"fmt"
	"strconv"
)

func HextoInt(hexadecimal string) string {
	num, _ := strconv.ParseInt(hexadecimal,16,64)
	return fmt.Sprint(num)
}