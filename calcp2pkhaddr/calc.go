package main

import (
	"base58"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	str:="c76d4e3417d712557686371fc24b8329d3863d2d"
	bs,_:=hex.DecodeString(str)
	buf:=make([]byte,0)
	preByte:=[]byte{0}
	buf=append(buf,preByte...)
	buf=append(buf,bs...)
	ts:=sha256.Sum256(buf)
	ts=sha256.Sum256(ts[0:])
	buf=append(buf,ts[0:4]...)
	fmt.Println(base58.Encode(buf))

}
