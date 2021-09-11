package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	bc *Blockchain
}

func (cli *CLI) printUsage() {
	fmt.Println("Usage: ")
	fmt.Println("  add <data> - add new block on top of blockchain")
	fmt.Println("  list - list blocks in blockchain")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) addBlock(data string) {
	cli.bc.AddBlock(data)
	fmt.Println("Success!")
}

func (cli *CLI) listBlocks() {
	rows, err := cli.bc.db.Query("SELECT * FROM blocks")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id int
		b := Block{}
		err = rows.Scan(&id, &b.Timestamp, &b.Data, &b.Hash, &b.PrevBlockHash, &b.Nonce)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Prev. hash: %x\n", b.PrevBlockHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		pow := NewProofOfWork(&b)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	switch os.Args[1] {
	case "add":
		if len(os.Args[2:]) > 0 {
			cli.addBlock(strings.Join(os.Args[2:], " "))
		} else {
			cli.printUsage()
			os.Exit(1)
		}
	case "list":
		cli.listBlocks()
	default:
		cli.printUsage()
		os.Exit(1)
	}
}
