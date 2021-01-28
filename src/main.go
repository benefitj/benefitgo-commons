package main

import (
	"github.com/benefitj/simulator-go/src/binary"
	"time"
)

func main() {
	//bytes := binary.Int32ToBytesBE(365536, 4)
	now := time.Now()
	bytes := binary.Int64ToBytesBE(now.Unix(), 8)
	println("bytes: ", binary.BytesToHex(bytes, false))
	println("value: ", binary.BytesToInt32BE(bytes, false))
	println("now: ", now.Format("2006-01-02 15:04:05"))
	println("now: ", now.Unix())
}
