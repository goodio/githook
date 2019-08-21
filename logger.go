package main

import "hook/logger"

var log *logger.Log

func init() {

	// 初始化
	log = logger.NewLog(1000)

	// 设置log级别
	log.SetLevel("Debug")

	// 设置输出引擎
	log.SetEngine("file", `{"level":4, "spilt":"size", "filename":"`+*storeDir+`/hook.log", "maxsize":10}`)
}
