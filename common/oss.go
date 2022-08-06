package common

import (
	"net/http"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var GlobalOSSClient cos.Client

func InitOSSClient() {
	client := cos.NewClient(nil, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv("COS_SECRET_ID"),
			SecretKey: os.Getenv("COS_SECRET_KEY"),
		},
	})
	GlobalOSSClient = *client
}
