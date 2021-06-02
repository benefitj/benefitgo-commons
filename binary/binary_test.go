package binary

import (
	"testing"
	"time"
)

// 测试二进制工具
func TestBinary(t *testing.T) {
	//bytes := Int32ToBytesBE(365536, 4)
	now := time.Now()
	bytes := Int64ToBytesBE(now.Unix(), 4)
	println("bytes: ", BytesToHex(bytes, false))
	println("bytes2: ", BytesToSplitHex(bytes, false, " ", 1))
	println("value: ", BytesToInt32BE(bytes, false))
	println("binary: ", BytesToBinary(bytes, " ", 1))
	println("now: ", now.Format("2006-01-02 15:04:05"))
	println("now: ", now.Unix())
}
