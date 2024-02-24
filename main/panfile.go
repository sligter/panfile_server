package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	uploadDirFlag = flag.String("upload-dir", "", "Directory to store uploaded files")
	portFlag      = flag.String("port", "2334", "Port on which the server will listen")
	uploadDir     string
)

func main() {
	// 解析命令行参数
	flag.Parse()

	// 如果没有提供-upload-dir参数，使用当前工作目录
	if *uploadDirFlag == "" {
		var err error
		uploadDir, err = os.Getwd() // 获取当前工作目录
		if err != nil {
			log.Fatalf("Error getting current directory: %v", err)
		}
	} else {
		uploadDir = *uploadDirFlag // 使用用户提供的目录
	}
	port := *portFlag

	// 在程序启动时创建自定义映射目录
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating upload directory: %v", err)
	}

	// 设置文件服务器，使用根路径来服务文件
	http.Handle("/", http.FileServer(http.Dir(uploadDir)))

	fmt.Printf("Server started on 0.0.0.0:%s\n", port)
	err = http.ListenAndServe("0.0.0.0:"+port, nil) // 注意这里`erro`变为`err`
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
