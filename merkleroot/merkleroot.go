package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Block struct {
	Txs []Hash `json:"tx"`
}

type Hash struct {
	Hash string `json:"hash"`
}

func main() {
	url := "https://blockchain.info/rawblock/0000000000000000000cde477b9421fc6e54a3eba20eacc682a2fb23bf29d21b"

	resp, err := http.Get(url)
	if err != nil {
	}
	defer resp.Body.Close()
	bf, err := ioutil.ReadAll(resp.Body)
	blockData := Block{}
	err = json.Unmarshal(bf, &blockData)
	txs := make([]string, 0)
	for i := 0; i < len(blockData.Txs); i++ {
		txs = append(txs, blockData.Txs[i].Hash)
	}
	fmt.Printf("%s", calMerkleRoot(txs))

}
func calMerkleRoot(txs []string) string {
	size := len(txs)
	for size > 1 {
		for j := 0; j < size; j = j + 2 {
			if j == size-1 {
				txs[j/2] = hash(txs[j], txs[j])
			} else {
				txs[j/2] = hash(txs[j], txs[j+1])
			}
		}
		size = (size + 1) / 2
	}
	return txs[0]
}

//计算tx1 以及tx2 合成的hash值
func hash(tx1, tx2 string) string {
	tb1, _ := hex.DecodeString(tx1)
	tb2, _ := hex.DecodeString(tx2)
	tb := make([]byte, 0)
	for i := 0; i < len(tb1); i++ {
		tb = append(tb, tb1[len(tb1)-1-i])
	}
	for i := 0; i < len(tb2); i++ {
		tb = append(tb, tb2[len(tb2)-1-i])
	}
	fsh256 := sha256.Sum256(tb[0:])
	ssh256 := sha256.Sum256(fsh256[0:])
	for i := 0; i < 16; i++ {
		ssh256[i], ssh256[32-1-i] = ssh256[32-1-i], ssh256[i]
	}
	return hex.EncodeToString(ssh256[0:])
}
