package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateFormattedString() string {
	// 定义随机数生成函数
	rand.Seed(time.Now().UnixNano())
	random := func() string {
		return fmt.Sprintf("%d", rand.Intn(10)) // 随机生成0-999的三位数字
	}

	// 定义获取时间戳后8位的函数
	timestampLast8Digits := func() string {
		timestamp := time.Now().Unix() // 获取当前时间戳（秒级）
		timestampStr := fmt.Sprintf("%d", timestamp)
		length := len(timestampStr)
		if length < 8 {
			return fmt.Sprintf("%08d", timestamp) // 如果不足8位，补零
		}
		return timestampStr[length-8:] // 返回最后8位
	}

	// 获取时间戳的后8位
	timestampPart := timestampLast8Digits()

	// 按照指定格式拼接字符串
	return fmt.Sprintf("%s%s%s%s%s%s%s", random(), timestampPart[:1], random(), timestampPart[1:5], random(), timestampPart[5:8], random())
}

func main() {
	// 初始化随机数种子

	// 调用函数生成格式化字符串
	for i := 0; i < 100; i++ {
		result := generateFormattedString()
		fmt.Println(result)
	}

}
