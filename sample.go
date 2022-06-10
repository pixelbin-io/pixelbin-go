package main

import (
	"fmt"
	"os"

	"github.com/pixelbin-dev/pixelbin-go/sdk/platform"
)

func main() {
	// create pixelbin config object
	config := platform.NewPixelbinConfig(
		"API_TOKEN",
		"https://api.pixelbin.io",
	)
	// set oauthclient
	config.SetOAuthClient()

	// create pixelbin client object
	pixelbin := platform.NewPixelbinClient(config)

	file, _ := os.Open("/home/rohit/deidara/1.jpeg")

	// Parameters for FileUpload function
	params := platform.FileUploadXQuery{
		File: file,
	}
	result, err := pixelbin.Assets.FileUpload(params)

	if err != nil {
		fmt.Println(err)
	}
	// use result
	fmt.Println(result)
}
