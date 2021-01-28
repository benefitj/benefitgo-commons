package main

import (
	"github.com/benefitj/benefitgo-commons/binary"
	"time"
)

func main() {
	//bytes := binary.Int32ToBytesBE(365536, 4)
	now := time.Now()
	bytes := binary.Int64ToBytesBE(now.Unix(), 4)
	println("bytes: ", binary.BytesToHex(bytes, false))
	println("bytes2: ", binary.BytesToSplitHex(bytes, false, " ", 1))
	println("value: ", binary.BytesToInt32BE(bytes, false))
	println("binary: ", binary.BytesToBinary(bytes, " ", 1))
	println("now: ", now.Format("2006-01-02 15:04:05"))
	println("now: ", now.Unix())
}
