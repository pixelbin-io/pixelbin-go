package common

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

func HttpRequest(method string, apiUrl string, queryParams map[string]string, data interface{}, headers map[string]string) ([]byte, error) {
	params := url.Values{}
	var (
		req         *http.Request
		err         error
		reqBodyMap  map[string]interface{}
		reqBodyJSON []byte
	)
	payload := &bytes.Reader{}
	if method == "GET" && queryParams != nil {
		for k, v := range queryParams {
			params.Set(k, v)
		}
	}
	if method != "GET" && data != nil {
		reqBodyJSON, reqBodyMap, err = ConvertInterfaceToByteAndMap(data)
		if err != nil {
			return nil, err
		}
		if _, ok := reqBodyMap["file"]; ok {
			metaValue := reflect.ValueOf(data).Elem()
			field := metaValue.FieldByName("File")
			if field.Type().String() == "*os.File" {
				var file *os.File
				var multipartContentType string
				file = (*os.File)(field.UnsafePointer())
				payload, multipartContentType, err = CreateMultiPartFormPayload(file, reqBodyMap)
				if err != nil {
					return nil, err
				}
				headers["content-Type"] = multipartContentType
			}
		} else {
			payload = bytes.NewReader(reqBodyJSON)
		}
	}
	req, err = http.NewRequest(method, apiUrl, payload)
	if err != nil {
		return nil, err
	}
	//Setting headers
	for k, v := range headers {
		req.Header[k] = []string{v}
	}
	//Setting query params
	req.URL.RawQuery = params.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	// read all response body
	resData, e := processHTTPResponse(res)
	if e != nil {
		return []byte{}, e
	}
	return resData, nil
}

func processHTTPResponse(res *http.Response) ([]byte, error) {
	var errResp *FDKError
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	// log.Println("data", string(data))
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = json.Unmarshal(data, &errResp)
		if err != nil {
			return []byte{}, err
		}
		return []byte{}, errResp
	}
	return data, nil
}

func ConvertInterfaceToByteAndMap(Body interface{}) ([]byte, map[string]interface{}, error) {
	reqBodyJSON, err := json.Marshal(Body)
	if err != nil {
		return nil, nil, err
	}
	reqBodyMap := make(map[string]interface{})
	err = json.Unmarshal(reqBodyJSON, &reqBodyMap)
	if err != nil {
		return nil, nil, err
	}
	return reqBodyJSON, reqBodyMap, nil
}

func CreateMultiPartFormPayload(file *os.File, reqBodyMap map[string]interface{}) (*bytes.Reader, string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		return nil, "", err
	}
	io.Copy(part, file)
	contentType := writer.FormDataContentType()
	for key, val := range reqBodyMap {
		if key != "file" {
			if b, ok := val.(string); ok {
				_ = writer.WriteField(key, b)
			}
		}
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	payload := bytes.NewReader(body.Bytes())
	return payload, contentType, nil
}
