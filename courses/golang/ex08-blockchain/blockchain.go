package main

import (
    "database/sql"
    "fmt"

    _ "github.com/mattn/go-sqlite3"
)

const dbName = "blockchain.db"

type Blockchain struct {
    db *sql.DB
    lastHash []byte
}

func NewBlockChain() *Blockchain {
    db, err := sql.Open("sqlite3", dbName)
    if err != nil {
        panic(err)
    }

    statement, _ := db.Prepare(`
        CREATE TABLE IF NOT EXISTS blocks (
            id INTEGER PRIMARY KEY,
            timestamp BIG INT,
            data TEXT,
            hash BLOB,
            prev_hash BLOB,
            nonce INT 
        )
    `)
    statement.Exec()

    row := db.QueryRow("SELECT hash FROM blocks WHERE ID = (SELECT MAX(ID) FROM blocks)")
    var prevHash []byte
    err = row.Scan(&prevHash)

    if err != nil {
        fmt.Println("No existing blockchain found. Creating a new one...")
        gns := NewGenesisBlock()

        _, err = db.Exec(`
            INSERT INTO blocks (timestamp, data, hash, prev_hash, nonce)
            values ($1, $2, $3, $4, $5)
        `, gns.Timestamp, string(gns.Data), gns.Hash, gns.PrevBlockHash, gns.Nonce)
        if err != nil {
            panic(err)
        }

        prevHash = gns.Hash
    }


    return &Blockchain{ db: db, lastHash: prevHash }
}

func (bc *Blockchain) AddBlock(data string) {
    nb := NewBlock(data, bc.lastHash)
    bc.lastHash = nb.Hash
    
    _, err := bc.db.Exec(`
        INSERT INTO blocks (timestamp, data, hash, prev_hash, nonce)
        values($1, $2, $3, $4, $5)
    `, nb.Timestamp, string(nb.Data), nb.Hash, nb.PrevBlockHash, nb.Nonce)
    if err != nil {
        panic(err)
    }
}
