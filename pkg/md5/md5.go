package md5

import (
	"crypto/md5"
	"encoding/hex"
)

// Sum md5 []byte to string 
func Sum(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
