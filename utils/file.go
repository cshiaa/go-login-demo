package utils

import (
	"os"
	"go.uber.org/zap"

	"github.com/cshiaa/go-login-demo/global"
)

// 判断文件是否存在
func IsFileExist(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        global.RY_LOG.Error("文件不存在", zap.String("filename", path))
		return false
    } else {
        global.RY_LOG.Error("找到文件", zap.String("filename", path))
		return true
    }
}