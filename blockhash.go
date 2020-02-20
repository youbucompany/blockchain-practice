package main

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

func main() {
	nVersion := 536870912
	hashPreBlock := "00000000000000000000d796ea8055d750678eaa287425015fdd150032d47f05"
	hashMerkleRoot := "f48d98288f525fea827083224e999fabba15ebeecf2171dfdebba30670fa0945"
	nTime := 1582183500
	nBits := 387062484
	nNonce := 1672903945
	hb := make([]byte, 80)
	binary.LittleEndian.PutUint32(hb[0:4], uint32(nVersion))
	hashPreBlockByte, _ := hex.DecodeString(hashPreBlock)
	for i := 0; i < 32; i++ {
		hb[4+i] = hashPreBlockByte[32-1-i]
	}
	hashMerkleRootByte, _ := hex.DecodeString(hashMerkleRoot)
	for i := 0; i < 32; i++ {
		hb[36+i] = hashMerkleRootByte[32-1-i]
	}
	binary.LittleEndian.PutUint32(hb[68:72], uint32(nTime))
	binary.LittleEndian.PutUint32(hb[72:76], uint32(nBits))
	binary.LittleEndian.PutUint32(hb[76:80], uint32(nNonce))
	fsh256 := sha256.Sum256(hb[0:])
	ssh256 := sha256.Sum256(fsh256[0:])
	for i := 0; i < 16; i++ {
		ssh256[i], ssh256[32-1-i] = ssh256[32-1-i], ssh256[i]
	}
	hashCurrBlock := hex.EncodeToString(ssh256[0:])
	fmt.Printf("%s\n", hashCurrBlock)
}
