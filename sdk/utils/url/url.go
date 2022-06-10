package url

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

var OPERTATION_SEPARATOR string = "~"
var PARAMETER_SEPARATOR string = ","
var VERSION2_REGEX string = "^v[1-2]$"
var URL_WITH_ZONE string = "^/([a-zA-Z0-9_-]*)/([a-zA-Z0-9_-]{6})/(.+)/(.*)$"
var URL_WITHOUT_ZONE string = "/([a-zA-Z0-9_-]*)/(.+)/(.*)"
var ZONE_SLUG string = "([a-zA-Z0-9_-]{6})"
var BASE_URL string = "https://cdn.pixelbin.io"

func UrlToObj(url string) (map[string]interface{}, error) {
	return getObjFromUrl(url)
}

func ObjToUrl(obj map[string]interface{}) (string, error) {
	return getUrlFromObj(obj)
}

func getUrlParts(pixelbinUrl string) (map[string]interface{}, error) {
	parseUrl, err := url.Parse(pixelbinUrl)
	if err != nil {
		return nil, err
	}
	urlDetails := map[string]interface{}{
		"protocol": parseUrl.Scheme,
		"host":     parseUrl.Host,
		"version":  "v1",
	}
	parts := strings.Split(parseUrl.Path, "/")
	if ok, _ := regexp.MatchString(VERSION2_REGEX, parts[1]); ok {
		urlDetails["version"] = parts[1]
		parts = append(parts[:1], parts[2:]...)
	}
	if len(parts[1]) < 3 {
		return nil, errors.New("invalid pixelbin url. Please make sure the url is correct")
	}

	if ok, _ := regexp.MatchString(URL_WITH_ZONE, strings.Join(parts, "/")); ok {
		urlDetails["cloudName"] = parts[1]
		parts = append(parts[:1], parts[2:]...)
		urlDetails["zoneSlug"] = parts[1]
		parts = append(parts[:1], parts[2:]...)
		urlDetails["pattern"] = parts[1]
		parts = append(parts[:1], parts[2:]...)
		urlDetails["filePath"] = strings.Join(parts[1:], "/")
	} else if ok, _ := regexp.MatchString(URL_WITHOUT_ZONE, strings.Join(parts, "/")); ok {
		urlDetails["cloudName"] = parts[1]
		parts = append(parts[:1], parts[2:]...)
		urlDetails["pattern"] = parts[1]
		parts = append(parts[:1], parts[2:]...)
		urlDetails["filePath"] = strings.Join(parts[1:], "/")
	} else {
		return nil, errors.New("invalid pixelbin url. Please make sure the url is correct")
	}
	return urlDetails, nil
}

func getPartsFromUrl(url string) (map[string]interface{}, error) {
	parts, err := getUrlParts(url)
	if err != nil {
		return nil, err
	}
	parts["zone"] = nil
	if val, ok := parts["zoneSlug"]; ok {
		parts["zone"] = val
		delete(parts, "zoneSlug")
	}
	parts["baseUrl"] = fmt.Sprintf("%s://%s", parts["protocol"], parts["host"])
	delete(parts, "protocol")
	delete(parts, "host")
	return parts, nil
}

func removeLeadingDash(str string) string {
	if len(str) > 0 && str[0] == '-' {
		return str[1:]
	}
	return str
}

func getParamsList(dSplit string, prefix string) []string {
	s := strings.Split(dSplit, "(")
	s1 := strings.Replace(s[1], ")", "", -1)
	s2 := strings.Replace(s1, prefix, "", -1)
	return strings.Split(removeLeadingDash(s2), ",")
}

func getParamsObject(paramsList []string) []map[string]string {
	params := []map[string]string{}
	for _, v := range paramsList {
		if strings.Contains(v, ":") {
			it := strings.Split(v, ":")
			if it[0] != "" {
				params = append(params, map[string]string{"key": it[0], "value": it[1]})
			}
		}
	}
	if len(params) > 0 {
		return params
	}
	return nil
}

func txtToOptions(dSplit string) map[string]interface{} {
	// Figure Out Module
	fullFnName := strings.Split(dSplit, "(")[0]

	arr := strings.Split(fullFnName, ".")
	pluginId := arr[0]
	operationName := arr[1]

	if pluginId == "p" {
		params := getParamsObject(getParamsList(dSplit, ""))
		for _, v := range params {
			if v["key"] == "n" {
				return map[string]interface{}{
					"plugin": pluginId,
					"name":   v["key"],
				}
			}
		}
		return nil
	}

	values := getParamsObject(getParamsList(dSplit, "."))
	transformation := map[string]interface{}{
		"plugin": pluginId,
		"name":   operationName,
	}
	if len(values) > 0 {
		transformation["values"] = values
	}
	return transformation
}

// FlattenSlice flattens nested slices of string
func FlattenSlice(slice []interface{}) []string {
	var flat []string

	for _, element := range slice {
		switch element.(type) {
		case []interface{}:
			flat = append(flat, FlattenSlice(element.([]interface{}))...)
		case []string:
			flat = append(flat, element.([]string)...)
		case string:
			flat = append(flat, element.(string))
		}
	}

	return flat
}

func getTransformationsFromPattern(pattern string, url string) []map[string]interface{} {
	if pattern == "original" {
		return []map[string]interface{}{}
	}
	dSplit := strings.Split(pattern, OPERTATION_SEPARATOR)
	arr := make([]map[string]interface{}, len(dSplit))
	for i, v := range dSplit {
		if strings.HasPrefix(v, "p:") {
			a := strings.Split(v, ":")
			v = fmt.Sprintf("p.apply(n:%s)", a[1])
		}
		arr[i] = txtToOptions(v)
	}

	// opts := FlattenSlice(arr)

	return arr
}

func getObjFromUrl(url string) (map[string]interface{}, error) {
	parts, err := getPartsFromUrl(url)
	if err != nil {
		return nil, errors.New("error Processing url. Please check the url is correct" + err.Error())
	}
	transformations := getTransformationsFromPattern(
		parts["pattern"].(string),
		url,
	)
	parts["transformations"] = transformations
	return parts, nil
}

func getPatternFromTransformations(tflist interface{}) (string, error) {
	transformationList := tflist.([]map[string]interface{})
	if len(transformationList) == 0 {
		return "", nil
	}
	newtransformationList := []string{}
	for _, o := range transformationList {
		if _, ok := o["name"]; ok {
			if o["plugin"] == "p" {
				newtransformationList = append(newtransformationList, fmt.Sprintf("p:%s", o["plugin"]))
			} else {
				if _, ok := o["values"]; !ok {
					o["values"] = []map[string]interface{}{}
				}
				paramlist := []string{}
				for _, items := range o["values"].([]map[string]interface{}) {
					if _, ok := items["key"]; !ok {
						return "", errors.New("key not specified")
					}
					if _, ok := items["value"]; !ok {
						return "", errors.New("value not specified for" + items["key"].(string))
					}
					paramlist = append(paramlist, fmt.Sprintf("%s:%s", items["key"], items["value"]))
				}
				paramstr := strings.Join(paramlist, PARAMETER_SEPARATOR)
				newtransformationList = append(newtransformationList, fmt.Sprintf("%s.%s(%s)", o["plugin"], o["name"], paramstr))
			}
		}
	}
	return strings.Join(newtransformationList, OPERTATION_SEPARATOR), nil
}

func getUrlFromObj(obj map[string]interface{}) (string, error) {
	if _, ok := obj["baseUrl"]; !ok {
		obj["baseUrl"] = BASE_URL
	}
	if _, ok := obj["cloudName"]; !ok {
		return "", errors.New("key cloudName should be defined")
	}
	if _, ok := obj["filePath"]; !ok {
		return "", errors.New("key filePath should be defined")
	}

	patternValue, err := getPatternFromTransformations(obj["transformations"])
	if err != nil {
		return "", err
	}
	if patternValue != "" {
		obj["pattern"] = patternValue
	} else {
		obj["pattern"] = "original"
	}
	_, versionOk := obj["version"]
	if !versionOk {
		obj["version"] = "v2"
	} else {
		match, _ := regexp.MatchString(VERSION2_REGEX, obj["version"].(string))
		if !match {
			obj["version"] = "v2"
		}
	}
	v, zoneOk := obj["zone"]
	if !zoneOk {
		obj["zone"] = ""
	} else {
		if v == nil {
			obj["zone"] = ""
		} else {
			match, _ := regexp.MatchString(ZONE_SLUG, obj["zone"].(string))
			if !match {
				obj["version"] = "v2"
			}
		}
	}

	urlKeySorted := []string{"baseUrl", "version", "cloudName", "zoneSlug", "pattern", "filePath"}
	urlArr := []string{}
	for _, v2 := range urlKeySorted {
		if val, ok := obj[v2]; ok {
			urlArr = append(urlArr, val.(string))
		}
	}
	return strings.Join(urlArr, "/"), nil
}
