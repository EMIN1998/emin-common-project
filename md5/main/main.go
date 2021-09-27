package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"bitbucket.org/klook/klook-libs/src/klook.libs/logger"
)

func GetFileMd5(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("os Open error")
		return "", err
	}
	md5 := md5.New()
	_, err = io.Copy(md5, file)
	if err != nil {
		fmt.Println("io copy error")
		return "", err
	}
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str, nil
}

func GetStringMd5(s string) string {
	md5 := md5.New()
	md5.Write([]byte(s))
	md5Str := hex.EncodeToString(md5.Sum(nil))
	return md5Str
}

func main() {
	str := "OtherInfoUsageStatistics_"
	res := GetStringMd5(str)

	logger.Infof("================res:%s", res)
}
