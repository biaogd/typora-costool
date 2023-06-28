package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 腾讯cos配置
type CosConfig struct {
	Url           string `json:"url"`
	SecretId      string `json:"secretId"`
	SecretKey     string `json:"secretKey"`
	DefaultBucket string `json:"defaultBucket"`
	KeyPrefix     string `json:"keyPrefix"`
}

var config CosConfig

var client *cos.Client

func loadConfig() {
	configFilePath := GetConfigFile()
	configFileBuf, err := os.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configFileBuf, &config)
	if err != nil {
		panic(err)
	}

	u, _ := url.Parse(config.Url)
	b := &cos.BaseURL{BucketURL: u}
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.SecretId,
			SecretKey: config.SecretKey,
		},
	})
}

func UploadFile(filePath string) string {
	loadConfig()
	fileExt := path.Ext(filePath)
	fileName := GeneralUUID() + fileExt
	fileKey := config.KeyPrefix + "/" + fileName
	_, err := client.Object.PutFromFile(context.Background(), fileKey, filePath, &cos.ObjectPutOptions{})
	if err != nil {
		panic(err)
	}
	url := client.Object.GetObjectURL(fileKey)
	return url.String()
}
