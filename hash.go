package yangoss

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func Hash(key string) string {
	h := hmac.New(sha256.New, []byte(key))

	randomNumber := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	// 生成一个100000到999999之间的随机数
	message := fmt.Sprintf("%d", randomNumber) // 写入消息数据
	h.Write([]byte(message))

	// 计算HMAC值
	signature := h.Sum(nil)

	// 将签名转换为十六进制字符串
	signatureHex := hex.EncodeToString(signature)
	signatureBase64 := base64.StdEncoding.EncodeToString([]byte(signatureHex))

	return signatureBase64
}
