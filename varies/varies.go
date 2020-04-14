package main

import (
	"encoding/hex"
	"fmt"
)

var twoBytesMax int = int(^(uint16(0)))
var fourBytesMax int = int(^(uint32(0)))

func main() {
	fmt.Printf("%v\n", hex.EncodeToString(serInt2Byte(103)))
}

func serInt2Byte(va int) (rs []byte) {
	if va < 253 {
		rs = append(rs, byte(va))
	} else if va <= twoBytesMax {
		rs = append(rs, byte(253))
		rs = append(rs, byte(va))    //取低位
		rs = append(rs, byte(va>>8)) //取高位
	} else if va <= fourBytesMax {
		rs = append(rs, byte(254))
		rs = append(rs, byte(va))
		rs = append(rs, byte(va>>8))
		rs = append(rs, byte(va>>16))
		rs = append(rs, byte(va>>24))
	} else {
		rs = append(rs, byte(255))
		rs = append(rs, byte(va>>8))
		rs = append(rs, byte(va>>16))
		rs = append(rs, byte(va>>24))
		rs = append(rs, byte(va>>32))
		rs = append(rs, byte(va>>40))
		rs = append(rs, byte(va>>48))
		rs = append(rs, byte(va>>56))
	}
	return
}
