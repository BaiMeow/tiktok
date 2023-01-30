package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"os/exec"
	"tiktok/config"
	"tiktok/log"
	"time"
)

func Init() {
	// check ffmpeg if existed
	ffmpegPath = config.Conf.GetString("video.ffmpeg")
	if ffmpegPath == "" {
		ffmpegPath = "ffmpeg"
	}
	_, err := exec.Command(ffmpegPath, "-version").CombinedOutput()
	if err != nil {
		log.Fatalln("ffmpeg is not existed, please install ffmpeg first")
	}
	// load jwt secret
	jwtSecret = []byte(config.Conf.GetString("auth.jwt_key"))
}

// sha256加密
func SHA256(password, salt string) string {
	code := []byte(salt)
	saltedPassword := append([]byte(password), ':')
	saltedPassword = append(saltedPassword, code...)

	hash := sha256.Sum256(saltedPassword)
	return hex.EncodeToString(hash[:])
}

// 生成随机salt
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var src = rand.NewSource(time.Now().UnixNano())

func GenerateSalt() string {
	b := make([]byte, 8)
	// A rand.Int63() generates 63 random bits
	for i, cache := 7, src.Int63(); i >= 0; i-- {
		b[i] = letters[int(cache)%len(letters)]
		cache >>= 6
	}
	return string(b[:])
}
