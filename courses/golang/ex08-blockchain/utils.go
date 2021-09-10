package main

import (
    "bytes"
    "encoding/binary"
    "log"
)

func IntToHex(num int64) []byte {
    buff := new(bytes.Buffer)
    err := binary.Write(buff, binary.BigEndian, num)
    if err != nil {
        log.Panic(err)
    }

    return buff.Bytes()
}

func HexToInt(hex []byte) int64 {
    var num int64
    buff := bytes.NewBuffer(hex)
    binary.Read(buff, binary.LittleEndian, &num)

    return num
}
