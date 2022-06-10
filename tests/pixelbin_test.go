package tests

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/pixelbin-dev/pixelbin-go/sdk/platform"
)

var settings = map[string]string{"domain": "https://api.testdomain.com", "apiSecret": "test-api-secret"}
var config *platform.PixelbinConfig
var pixelbin *platform.PixelbinClient
var domain string
var apiSecret string

func Init() {
	settings["apiSecret"] = apiSecret
	settings["domain"] = domain
	settings["folderName"] = "testdir"
	settings["folderPath"] = "/"
	config = platform.NewPixelbinConfig(settings["apiSecret"], settings["domain"])
	config.SetOAuthClient()
	pixelbin = platform.NewPixelbinClient(config)
}

func TestCreateFolder(t *testing.T) {
	params := platform.CreateFolderXQuery{Name: settings["folderName"], Path: settings["folderPath"]}
	_, err := pixelbin.Assets.CreateFolder(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestFileUploadCase1(t *testing.T) {
	file, _ := os.Open("1.jpeg")
	params := platform.FileUploadXQuery{File: file}
	_, err := pixelbin.Assets.FileUpload(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestFileUploadCase2(t *testing.T) {
	file, _ := os.Open("1.jpeg")
	params := platform.FileUploadXQuery{
		File:             file,
		Path:             settings["folderName"],
		Name:             "1",
		Access:           "public-read",
		Tags:             []string{"tag1", "tag2"},
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	_, err := pixelbin.Assets.FileUpload(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestListFilesCase1(t *testing.T) {
	params := platform.ListFilesXQuery{}
	_, err := pixelbin.Assets.ListFiles(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestListFilesCase2(t *testing.T) {
	params := platform.ListFilesXQuery{}
	_, err := pixelbin.Assets.ListFiles(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUrlUpload(t *testing.T) {
	params := platform.UrlUploadXQuery{
		URL:              "https://www.fetchfind.com/blog/wp-content/uploads/2017/08/cat-2734999_1920-5-common-cat-sounds.jpg",
		Path:             settings["folderName"],
		Name:             "2",
		Access:           "public-read",
		Tags:             []string{"cat", "animal"},
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	_, err := pixelbin.Assets.UrlUpload(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestCreateSignedUrlCase1(t *testing.T) {
	params := platform.CreateSignedUrlXQuery{}
	_, err := pixelbin.Assets.CreateSignedUrl(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestCreateSignedUrlCase2(t *testing.T) {
	params := platform.CreateSignedUrlXQuery{
		Name:             "1",
		Path:             settings["folderName"],
		Format:           "jpeg",
		Access:           "public-read",
		Tags:             []string{"tag1", "tag2"},
		Metadata:         map[string]interface{}{},
		Overwrite:        false,
		FilenameOverride: true,
	}
	_, err := pixelbin.Assets.CreateSignedUrl(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateFileCase1(t *testing.T) {
	params := platform.UpdateFileXQuery{
		FileId: "1",
		Name:   "1_",
	}
	_, err := pixelbin.Assets.UpdateFile(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateFileCase2(t *testing.T) {
	params := platform.UpdateFileXQuery{
		FileId:   fmt.Sprintf("%s/1", settings["folderName"]),
		Name:     fmt.Sprintf("%s_", settings["folderName"]),
		Path:     settings["folderName"],
		Access:   "private",
		IsActive: false,
		Tags:     []string{"tag1", "tag2"},
		Metadata: map[string]interface{}{"key": "value"},
	}
	_, err := pixelbin.Assets.UpdateFile(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetFileByFileId(t *testing.T) {
	params := platform.GetFileByFileIdXQuery{
		FileId: fmt.Sprintf("%s/2", settings["folderName"]),
	}
	_, err := pixelbin.Assets.GetFileByFileId(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestDeleteFile(t *testing.T) {
	params := platform.DeleteFileXQuery{
		FileId: "1_",
	}
	_, err := pixelbin.Assets.DeleteFile(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestUpdateFolder(t *testing.T) {
	params := platform.UpdateFolderXQuery{
		FolderId: settings["folderName"],
		IsActive: false,
	}
	_, err := pixelbin.Assets.UpdateFolder(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestDeleteFolder(t *testing.T) {
	params := platform.ListFilesXQuery{}
	res, err := pixelbin.Assets.ListFiles(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}
	jsonbody, err := json.Marshal(res)
	if err != nil {
		t.Errorf("Failed ! error in marshalling %v", err)
	}

	var id string = ""
	var x platform.ListFilesResponse
	err = json.Unmarshal(jsonbody, &x)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	}
	for i := range x.Items {
		item := x.Items[i]
		if item.Type == "folder" && item.Name == settings["folderName"] {
			id = item.ID
			break
		}
	}
	if id != "" {
		param := platform.DeleteFolderXQuery{
			ID: id,
		}
		_, err = pixelbin.Assets.DeleteFolder(param)
		if err != nil {
			t.Errorf("Failed ! got err %v", err)
		} else {
			t.Logf("Success")
		}
	}
}

func TestGetModules(t *testing.T) {
	params := platform.GetModulesXQuery{}
	_, err := pixelbin.Assets.GetModules(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestGetModule(t *testing.T) {
	params := platform.GetModuleXQuery{
		Identifier: "t",
	}
	_, err := pixelbin.Assets.GetModule(params)
	if err != nil {
		t.Errorf("Failed ! got err %v", err)
	} else {
		t.Logf("Success")
	}
}

func TestMain(m *testing.M) {

	flag.StringVar(&domain, "domain", settings["domain"], "Domain Address")
	flag.StringVar(&apiSecret, "apiSecret", settings["apiSecret"], "Api secret token")

	flag.Parse()
	Init()
	code := m.Run()
	os.Exit(code)
}
