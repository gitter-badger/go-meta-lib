package main

import (
	"encoding/json"
	"fmt"
	"github.com/ipfs/go-ipfs-api"
)

type MetaWrapper struct {
	Metamodel struct {
		Operations []struct {
			Stack int `json:"stack"`
			Op    []struct {
				Mode   string  `json:"mode"`
				Weight float64 `json:"weight"`
			} `json:"op"`
		} `json:"operations"`
	} `json:"metamodel"`
	Metadata struct {
		Created       int    `json:"created"`
		Author        string `json:"author"`
		AuthorAddress string `json:"author_address"`
	} `json:"metadata"`
}

func main() {
	testJSON := []byte(`{"metamodel":{"operations":[{"stack":1,"op":[{"mode":"TxAGHDJE764Y2HT875Y98Y32GH","weight":1}]},{"stack":2,"op":[{"mode":"TxAGHDJE764Y2HT875Y98Y32GH","weight":1}]}]},"metadata":{"created":1513648355,"author":"Sam Sendelbach","author_address":"0x1832yfhsdfgy83288h5"}}`)

	var testUnmarshal MetaWrapper
	err := json.Unmarshal(testJSON, &testUnmarshal)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", testUnmarshal)

	sh := shell.NewShell("localhost:5001")
	err = sh.Get("QmW2WQi7j6c7UgJTarActp7tDNikE4B2qXtFCfLPdsgaTQ/cat.jpg", "cat.jpg")
	if err != nil {
		fmt.Println("fucked up")
	}
}
