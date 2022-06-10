package common

import (
	// "bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"

	// "io/ioutil"
	"log"
	// "net/http"
	"net/url"
	// "regexp"
	"sort"
	"strings"
	"time"
)

//EncodeToBase64 gives base64 encoded string
func EncodeToBase64(val string) string {
	return base64.StdEncoding.EncodeToString([]byte(val))
}

//MapToURLEncodedString converts map to url encoded string
func MapToURLEncodedString(valMap map[string]string) (encodedData string) {
	data := url.Values{}
	for k, v := range valMap {
		data.Set(k, v)
	}
	encodedData = data.Encode()
	return
}

type SignatureModel struct {
	domain         string
	method         string
	url            string
	queryString    string
	headers        map[string]string
	body           interface{}
	excludeHeaders []string
}

func NewSignatureModel(domain string, method string, url string, queryString string, headers map[string]string, body interface{}, excludeHeaders []string) *SignatureModel {
	return &SignatureModel{domain, method, url, queryString, headers, body, excludeHeaders}
}

//GetRequestedDateTime reuse and returns the date time created
func GetRequestedDateTime() string {
	return strings.Replace(strings.Replace(time.Now().UTC().Format(time.RFC3339), "-", "", -1), ":", "", -1)
}

func hash(b []byte) string {
	h := sha256.New()
	_, err := h.Write(b)
	if err != nil {
		log.Println("Failed to write request body to hash. err = ", err)
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}

func sortHeaders(header, cutset string) string {
	splittedHeaders := strings.Split(header, cutset)
	sort.SliceStable(splittedHeaders, func(i, j int) bool {
		return splittedHeaders[i] < splittedHeaders[j]
	})
	return strings.Join(splittedHeaders, cutset)
}

func (model *SignatureModel) AddSignatureToHeaders(signQuery bool) (map[string]string, error) {
	queryString := model.queryString
	ebgDate := GetRequestedDateTime()
	headersStr := ""
	host := strings.Replace(strings.Replace(model.domain, "http://", "", 1), "https://", "", 1)
	model.headers["host"] = host
	if !signQuery {
		model.headers["x-ebg-param"] = ebgDate
	} else {
		if queryString != "" {
			queryString = fmt.Sprintf("%s&x-ebg-param=%s", queryString, ebgDate)
		} else {
			queryString = fmt.Sprintf("%s?x-ebg-param=%s", queryString, ebgDate)

		}
	}
	excluded_headers := make(map[string]string)
	for x := 0; x < len(model.excludeHeaders); x++ {
		if value, ok := model.headers[model.excludeHeaders[x]]; ok {
			excluded_headers[model.excludeHeaders[x]] = value
			delete(model.headers, model.excludeHeaders[x])
		} else {
			excluded_headers[model.excludeHeaders[x]] = ""
		}
	}

	keys := make([]string, 0, len(model.headers))
	for k := range model.headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		headersStr += fmt.Sprintf("%s:%s\n", k, model.headers[k])
	}

	bodyHex := hash([]byte(""))
	if model.body != nil {
		b, err := json.Marshal(model.body)
		if err != nil {
			return nil, err
		}
		bodyHex = hash([]byte(string(b)))
	}

	var HEADERS_TO_INCLUDE = []string{"host", "x-ebg-.*"}
	//generate signedHeaders
	var signedHeaders string
	for key := range model.headers {
		if contains(HEADERS_TO_INCLUDE, strings.ToLower(key), true) {
			signedHeaders += fmt.Sprintf("%s;", strings.Trim(strings.ToLower(key), " "))
		}

	}
	signedHeaders = strings.TrimSuffix(signedHeaders, ";")
	signedHeaders = sortHeaders(signedHeaders, ";")

	//generate canonical request
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		strings.ToUpper(model.method),
		model.url,
		queryString,
		headersStr,
		signedHeaders,
		bodyHex,
	)
	canonicalRequestHashed := hash([]byte(canonicalRequest))
	if canonicalRequestHashed == "" {
		log.Println("Cannot sign the request. Empty canonical request hash")
		return nil, errors.New("cannot sign the request. Empty canonical request hash")
	}

	//generate string to sign
	strToSign := fmt.Sprintf("%s\n%s", ebgDate, canonicalRequestHashed)

	//generate final signature
	h := hmac.New(sha256.New, []byte("1234567"))
	_, err := h.Write([]byte(strToSign))
	if err != nil {
		log.Println("Cannot sign the request. Failed to write strToSign to hash")
		return nil, errors.New("cannot sign the request. Failed to write strToSign to hash")
	}
	sha := hex.EncodeToString(h.Sum(nil))
	signature := fmt.Sprintf("v1:%s", sha)
	if signQuery {
		queryString += fmt.Sprintf("&x-ebg-signature=%s", signature)
	} else {
		model.headers["x-ebg-signature"] = signature
	}
	for key, element := range excluded_headers {
		if element != "" {
			model.headers[key] = element
		}
	}
	return model.headers, nil

}

// contains checks if a string is present in a slice
func contains(s []string, str string, isRegex bool) bool {
	for _, v := range s {
		if isRegex {
			r, _ := regexp.Compile(v)
			if !r.MatchString(str) {
				continue
			} else {
				return true
			}
		}
		if v == str {
			return true
		}
	}
	return false
}
