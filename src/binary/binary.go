package binary

import (
	"strings"
)

type ByteOrder uint8

// 字节顺序
const (
	_             ByteOrder = iota
	BIG_ENDIAN              // 大端
	LITTLE_ENDIAN           // 小端
)

// 16进制
var hexUpperCase = []rune("0123456789ABCDEF")
var hexLowerCase = []rune("0123456789abcdef")

// 二进制字符串
var binaryStrings = [...]string{
	"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111",
	"1000", "1001", "1010", "1011", "1100", "1101", "1110", "1111",
}

// 16进制字符
var hexChars = [...]byte{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F',
}

// mask
var masks = [][]byte{
	{0b00000001, 0b00000010, 0b00000100, 0b00001000, 0b00010000, 0b00100000, 0b01000000, 0b10000000},
	{0b00000011, 0b00000110, 0b00001100, 0b00011000, 0b00110000, 0b01100000, 0b11000000},
	{0b00000111, 0b00001110, 0b00011100, 0b00111000, 0b01110000, 0b11100000},
	{0b00001111, 0b00011110, 0b00111100, 0b01111000, 0b11110000},
	{0b00011111, 0b00111110, 0b01111100, 0b11111000},
	{0b00111111, 0b01111110, 0b11111100},
	{0b01111111, 0b11111110},
	{0b11111111},
}

// Mask 取值
// bits: 标志位
// size: bit数量(1~8)
// position: bit位置(0~7)
// Returns 返回转换后的数值
func Mask(bits byte, size byte, position byte) byte {
	if size <= 0 || size > 8 {
		panic("Required size between 1 and 8")
	}
	var mask = masks[(size-1)%8]
	var i = max(int((position%8)+1-size), 0)
	return (bits & mask[i] & 0xFF) >> i
}

// Int16ToBytes 将整数转换成字节数组，num为16位整数，size为字节长度，order为字节顺序
// Returns 返回转换后的字节数组
func Int16ToBytes(num int16, size int, order ByteOrder) []byte {
	var bit = size * 8
	var bytes = make([]byte, size)
	for i := 0; i < size; i++ {
		// 大端存储：高位在前，低位在后
		// 小端存储：低位在前，高位在后
		if order == BIG_ENDIAN {
			bytes[i] = byte(num >> ((bit - 8) - i*8))
		} else {
			bytes[i] = byte(num >> (i * 8))
		}
	}
	return bytes
}

// Int16ToBytesBE 将整数转换成大端的字节数组，num为16位整数，size为字节长度
// Returns 返回转换后的字节数组
func Int16ToBytesBE(num int16, size int) []byte {
	return Int16ToBytes(num, size, BIG_ENDIAN)
}

// Int16ToBytesLE 将整数转换成小端的字节数组，num为16位整数，size为字节长度
// Returns 返回转换后的字节数组
func Int16ToBytesLE(num int16, size int) []byte {
	return Int16ToBytes(num, size, LITTLE_ENDIAN)
}

// Int32ToBytes 将整数转换成字节数组，num为32位整数，size为字节长度，order为字节顺序
// Returns 返回转换后的字节数组
func Int32ToBytes(num int32, size int, order ByteOrder) []byte {
	var bit = size * 8
	var bytes = make([]byte, size)
	for i := 0; i < size; i++ {
		// 大端存储：高位在前，低位在后
		// 小端存储：低位在前，高位在后
		if order == BIG_ENDIAN {
			bytes[i] = byte(num >> ((bit - 8) - i*8))
		} else {
			bytes[i] = byte(num >> (i * 8))
		}
	}
	return bytes
}

// Int32ToBytesBE 将整数转换成大端的字节数组，num为32位整数，size为字节长度
// Returns 返回转换后的字节数组
func Int32ToBytesBE(num int32, size int) []byte {
	return Int32ToBytes(num, size, BIG_ENDIAN)
}

// Int32ToBytesLE 将整数转换成小端的字节数组，num为32位整数，size为字节长度
// Returns 返回转换后的字节数组
func Int32ToBytesLE(num int32, size int) []byte {
	return Int32ToBytes(num, size, LITTLE_ENDIAN)
}

// Int64ToBytes 将整数转换成字节数组，num为64位整数，size为字节长度，order为字节顺序
// Returns 返回转换后的字节数组
func Int64ToBytes(num int64, size int, order ByteOrder) []byte {
	var bit = size * 8
	var bytes = make([]byte, size)
	for i := 0; i < size; i++ {
		// 大端存储：高位在前，低位在后
		// 小端存储：低位在前，高位在后
		if order == BIG_ENDIAN {
			bytes[i] = byte(num >> ((bit - 8) - i*8))
		} else {
			bytes[i] = byte(num >> (i * 8))
		}
	}
	return bytes
}

// Int64ToBytesBE 将整数转换成大端的字节数组，num为64位整数，size为字节长度
// Returns 返回转换后的字节数组
func Int64ToBytesBE(num int64, size int) []byte {
	return Int64ToBytes(num, size, BIG_ENDIAN)
}

// Int64ToBytesLE 将整数转换成小端的字节数组，num为64位整数，size为字节长度
// Returns 返回转换后的字节数组
func Int64ToBytesLE(num int64, size int) []byte {
	return Int64ToBytes(num, size, LITTLE_ENDIAN)
}

// BytesToInt16 将字节数组转换成整数，bytes为字节数组，order为字节顺序，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt16(bytes []byte, order ByteOrder, signed bool) int16 {
	// 大端存储：高位在前，低位在后
	// 小端存储：低位在前，高位在后
	// 正数的原码，高位为0，反码/补码均与原码相同；
	// 负数的原码：高位为1, 其他为正数的原码；反码是除符号位，其它按位取反；补码在反码的基础上 + 1
	var v int16 = 0
	if order == BIG_ENDIAN {
		if signed && ((bytes[0]&0b10000000)>>7) == 1 {
			for i := 0; i < len(bytes); i++ {
				v <<= 8
				v |= int16(^bytes[i] & 0xFF)
			}
			v = (-v) - 1
		} else {
			for i := 0; i < len(bytes); i++ {
				v <<= 8
				v |= int16(bytes[i] & 0xFF)
			}
		}
	} else {
		if signed && ((bytes[len(bytes)-1]&0b10000000)>>7) == 1 {
			for i := len(bytes) - 1; i >= 0; i-- {
				v <<= 8
				v |= int16(^bytes[i] & 0xFF)
			}
			v = (-v) - 1
		} else {
			for i := len(bytes) - 1; i >= 0; i-- {
				v <<= 8
				v |= int16(bytes[i] & 0xFF)
			}
		}
	}
	return v
}

// BytesToInt16 将大端的字节数组转换成整数，bytes为字节数组，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt16BE(bytes []byte, signed bool) int16 {
	return BytesToInt16(bytes, BIG_ENDIAN, signed)
}

// BytesToInt16 将小端的字节数组转换成整数，bytes为字节数组，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt16LE(bytes []byte, signed bool) int16 {
	return BytesToInt16(bytes, LITTLE_ENDIAN, signed)
}

// BytesToInt32 将字节数组转换成整数，bytes为字节数组，order为字节顺序，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt32(bytes []byte, order ByteOrder, signed bool) int32 {
	// 大端存储：高位在前，低位在后
	// 小端存储：低位在前，高位在后
	// 正数的原码，高位为0，反码/补码均与原码相同；
	// 负数的原码：高位为1, 其他为正数的原码；反码是除符号位，其它按位取反；补码在反码的基础上 + 1
	var v int32 = 0
	if order == BIG_ENDIAN {
		if signed && ((bytes[0]&0b10000000)>>7) == 1 {
			for i := 0; i < len(bytes); i++ {
				v <<= 8
				v |= int32(^bytes[i] & 0xFF)
			}
			v = (-v) - 1
		} else {
			for i := 0; i < len(bytes); i++ {
				v <<= 8
				v |= int32(bytes[i] & 0xFF)
			}
		}
	} else {
		if signed && ((bytes[len(bytes)-1]&0b10000000)>>7) == 1 {
			for i := len(bytes) - 1; i >= 0; i-- {
				v <<= 8
				v |= int32(^bytes[i] & 0xFF)
			}
			v = (-v) - 1
		} else {
			for i := len(bytes) - 1; i >= 0; i-- {
				v <<= 8
				v |= int32(bytes[i] & 0xFF)
			}
		}
	}
	return v
}

// BytesToInt32 将大端的字节数组转换成整数，bytes为字节数组，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt32BE(bytes []byte, signed bool) int32 {
	return BytesToInt32(bytes, BIG_ENDIAN, signed)
}

// BytesToInt32 将小端的字节数组转换成整数，bytes为字节数组，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt32LE(bytes []byte, signed bool) int32 {
	return BytesToInt32(bytes, LITTLE_ENDIAN, signed)
}

// BytesToInt64 将字节数组转换成整数，bytes为字节数组，order为字节顺序，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt64(bytes []byte, order ByteOrder, signed bool) int64 {
	// 大端存储：高位在前，低位在后
	// 小端存储：低位在前，高位在后
	// 正数的原码，高位为0，反码/补码均与原码相同；
	// 负数的原码：高位为1, 其他为正数的原码；反码是除符号位，其它按位取反；补码在反码的基础上 + 1
	var v int64 = 0
	if order == BIG_ENDIAN {
		if signed && ((bytes[0]&0b10000000)>>7) == 1 {
			for i := 0; i < len(bytes); i++ {
				v <<= 8
				v |= int64(^bytes[i] & 0xFF)
			}
			v = (-v) - 1
		} else {
			for i := 0; i < len(bytes); i++ {
				v <<= 8
				v |= int64(bytes[i] & 0xFF)
			}
		}
	} else {
		if signed && ((bytes[len(bytes)-1]&0b10000000)>>7) == 1 {
			for i := len(bytes) - 1; i >= 0; i-- {
				v <<= 8
				v |= int64(^bytes[i] & 0xFF)
			}
			v = (-v) - 1
		} else {
			for i := len(bytes) - 1; i >= 0; i-- {
				v <<= 8
				v |= int64(bytes[i] & 0xFF)
			}
		}
	}
	return v
}

// BytesToInt64 将大端的字节数组转换成整数，bytes为字节数组，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt64BE(bytes []byte, signed bool) int64 {
	return BytesToInt64(bytes, BIG_ENDIAN, signed)
}

// BytesToInt64 将小端的字节数组转换成整数，bytes为字节数组，signed为是否为有符号位
// Returns 返回转换后的整数
func BytesToInt64LE(bytes []byte, signed bool) int64 {
	return BytesToInt64(bytes, LITTLE_ENDIAN, signed)
}

// BytesToBinary 字节数组转换成二进制字符串，bytes为字节数组，split为分隔符，size为分割的长度
func BytesToBinary(bytes []byte, split string, size int) string {
	var sb strings.Builder
	for i := 0; i < len(bytes); i++ {
		// 高四位
		sb.WriteString(binaryStrings[(bytes[i]&0xF0)>>4])
		// 低四位
		sb.WriteString(binaryStrings[bytes[i]&0x0F])
		// 分割
		if i%size == 0 {
			if i > 0 && i < len(bytes)-1 {
				sb.WriteString(split)
			}
		}
	}
	return sb.String()
}

// BytesToHex 字节数组转换成16进制字符串
func BytesToHex(bytes []byte, lowerCase bool) string {
	return BytesToCustomHex(bytes, lowerCase, "", "", 1)
}

// BytesToSplitHex 字节数组转换成16进制字符串，并且按照长度填充
func BytesToSplitHex(bytes []byte, lowerCase bool, fill string, length int) string {
	return BytesToCustomHex(bytes, lowerCase, "", fill, length)
}

// BytesToCustomHex 字节数组转换成16进制字符串
func BytesToCustomHex(bytes []byte, lowerCase bool, prefix string, suffix string, length int) string {
	var hex []rune
	if lowerCase {
		hex = hexLowerCase
	} else {
		hex = hexUpperCase
	}
	var split = max(length, 1)
	var sb strings.Builder
	for i, j := 0, 1; i < len(bytes); i, j = i+1, j+1 {
		var b = bytes[i]
		// 填充前缀
		if i%split == 0 {
			sb.WriteString(prefix)
		}
		// 原16进制数据
		sb.WriteRune(hex[((b & 0xF0) >> 4)])
		sb.WriteRune(hex[(b & 0x0F)])
		// 填充后缀
		if i < (len(bytes)-1) && (j%split == 0) {
			sb.WriteString(suffix)
		}
	}
	return sb.String()
}

// 取最大值
func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// 取最小值
func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
