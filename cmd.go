package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func GetConfigFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("read home dir error", err.Error())
		os.Exit(-1)
	}
	homeConfigDir := path.Join(homeDir, ".config")
	_, err = os.Stat(homeConfigDir)
	if err != nil && os.IsNotExist(err) {
		os.Mkdir(homeConfigDir, os.ModePerm)
	} else if err != nil {
		fmt.Println("get config dir error", err)
		os.Exit(-1)
	}
	return path.Join(homeConfigDir, "costool.json")
}

func initConfig(config CosConfig) {
	configPath := GetConfigFile()
	buf, err := json.Marshal(config)
	if err != nil {
		fmt.Println("json dumps error", err)
	}
	err = os.WriteFile(configPath, buf, os.ModePerm)
	if err != nil {
		fmt.Println("init config file error")
	} else {
		fmt.Println("init config file ", configPath, " success")
	}
}

func Run() {
	args := os.Args
	if len(args) == 2 && args[1] == "init" {
		fmt.Print("Cos Url: ")
		var url string
		fmt.Scanln(&url)
		var secretId string
		fmt.Print("SecretId: ")
		fmt.Scanln(&secretId)
		var secretKey string
		fmt.Print("SecretKey: ")
		fmt.Scanln(&secretKey)
		var defaultBucket string
		fmt.Print("Default Bucket: ")
		fmt.Scanln(&defaultBucket)
		var keyPrefix string
		fmt.Print("Object Key Prefix: ")
		fmt.Scanln(&keyPrefix)
		fmt.Println(url, secretId, secretKey, defaultBucket, keyPrefix)
		config := CosConfig{
			Url:           url,
			SecretId:      secretId,
			SecretKey:     secretKey,
			DefaultBucket: defaultBucket,
			KeyPrefix:     keyPrefix,
		}
		initConfig(config)
		return
	}

	for i := 1; i < len(args); i++ {
		filePath := args[i]
		_, err := os.Stat(filePath)
		if err != nil {
			panic(err)
		}
		fmt.Println(UploadFile(filePath))
	}

}
