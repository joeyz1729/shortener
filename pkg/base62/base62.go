package base62

import (
	"strings"
)

// 62进制转换的模块
// 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ

// const base62Str = `0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`
// const base62Str = `J0rs12O5TUV8IW7D9aBdXeCfghiMQj3klmop6qtuvbcwx4zAEFGHKLNnPRYSZy`
// 为了避免被人恶意请求，我们可以将上面的字符串打乱

var (
	baseStr string
	// baseStrLen uint64
)
// MustInit 要使用base62这包必须要调用该函数完成初始化
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string!")
	}
	baseStr = bs
	// baseStrLen = uint64(len(bs))
}

// Int2String convert decimal number to 62 base number
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(baseStr[0])
	}
	bl := []byte{} // 23 40 1
	for seq > 0 {
		mod := seq % 62
		div := seq / 62
		bl = append(bl, baseStr[mod])
		seq = div
	}
	// reverse
	return string(reverse(bl))
}

// String2Int convert 62 base number to decimal number
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	// traverse from right to left
	base := 1
	for _, b := range bl {
		// base := math.Pow(62, float64(idx))
		seq += uint64(strings.Index(baseStr, string(b))) * uint64(base)
		base *= 62
	}
	return seq
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
