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
	prefixFlag    = flag.String("pre", "static", "prefix name for file")
	portFlag      = flag.String("port", "2333", "Port on which the server will listen")
	uploadDir     string
)

func main() {
	//解析命令行参数
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
	prefix := fmt.Sprintf("/%s/", *prefixFlag)

	// 在程序启动时创建自定义映射目录
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating upload directory: %v", err)
	}

	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(uploadDir))))

	fmt.Printf("Server started on 0.0.0.0:%s%s\n", port, prefix)
	erro := http.ListenAndServe("0.0.0.0:"+port, nil)
	if erro != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
