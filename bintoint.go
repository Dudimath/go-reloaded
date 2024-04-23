package go_reloaded

import (
	"fmt"
	"strconv"
)

func BintoInt(binary string) string  {
	num, _ := strconv.ParseInt(binary,2,64)
	return fmt.Sprint(num)	
}