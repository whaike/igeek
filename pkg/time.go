package pkg

import "time"

// 通用格式化时间戳
func CommonTime() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}
