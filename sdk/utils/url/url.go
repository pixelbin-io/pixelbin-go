package url

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type UrlRegex struct {
	WithZone          string
	WithoutZone       string
	WithWorkerAndZone string
	WithWorker        string
}

var pixelBinDomainRegex = UrlRegex{
	WithZone:          "^/([a-zA-Z0-9_-]*)/([a-zA-Z0-9_-]{6})/(.+)/(.*)$",
	WithoutZone:       "^/([a-zA-Z0-9_-]*)/(.+)/(.*)",
	WithWorkerAndZone: "^/([a-zA-Z0-9_-]*)/([a-zA-Z0-9_-]{6})/wrkr/(.*)$",
	WithWorker:        "^/([a-zA-Z0-9_-]*)/wrkr/(.*)$",
}

var customDomainRegex = UrlRegex{
	WithoutZone:       "^/(.+)/(.*)",
	WithZone:          "^/([a-zA-Z0-9_-]{6})/(.+)/(.*)$",
	WithWorkerAndZone: "^/([a-zA-Z0-9_-]{6})/wrkr/(.*)$",
	WithWorker:        "^/wrkr/(.*)$",
}

var OPERTATION_SEPARATOR string = "~"
var PARAMETER_SEPARATOR string = ","
var VERSION2_REGEX string = "^v[1-2]$"
var URL_WITH_ZONE string = "^/([a-zA-Z0-9_-]*)/([a-zA-Z0-9_-]{6})/(.+)/(.*)$"
var URL_WITHOUT_ZONE string = "/([a-zA-Z0-9_-]*)/(.+)/(.*)"
var ZONE_SLUG string = "([a-zA-Z0-9_-]{6})"
var BASE_URL string = "https://cdn.pixelbin.io"

// UrlToObjOption is a functional option for configuring the UrlToObj function.
type UrlToObjOption func(*urlToObjConfig)

// WithCustomDomain sets the IsCustomDomain field in the configuration.
func WithCustomDomain(IsCustomDomain bool) UrlToObjOption {
	return func(config *urlToObjConfig) {
		config.IsCustomDomain = IsCustomDomain
	}
}

type urlToObjConfig struct {
	IsCustomDomain bool
}

func UrlToObj(url string, opts ...UrlToObjOption) (map[string]interface{}, error) {
	config := urlToObjConfig{
		IsCustomDomain: false,
	}

	for _, opt := range opts {
		opt(&config)
	}

	return getObjFromUrl(url, config)
}

func ObjToUrl(obj map[string]interface{}) (string, error) {
	return getUrlFromObj(obj)
}

func getUrlParts(pixelbinUrl string, config urlToObjConfig) (map[string]interface{}, error) {
	parseUrl, err := url.Parse(pixelbinUrl)
	if err != nil {
		return nil, err
	}
	urlDetails := map[string]interface{}{
		"protocol": parseUrl.Scheme,
		"host":     parseUrl.Host,
		"search": map[string]string{
			"dpr":    parseUrl.Query().Get("dpr"),
			"f_auto": parseUrl.Query().Get("f_auto"),
		},
		"version":    "v1",
		"cloudName":  nil,
		"worker":     false,
		"workerPath": "",
	}
	parts := strings.Split(parseUrl.Path, "/")

	if config.IsCustomDomain {
		if ok, _ := regexp.MatchString(VERSION2_REGEX, parts[1]); ok {
			urlDetails["version"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
		} else {
			return nil, errors.New("Invalid pixelbin url. Please make sure the url is correct.")
		}

		if ok, _ := regexp.MatchString(customDomainRegex.WithWorkerAndZone, strings.Join(parts, "/")); ok {
			urlDetails["zoneSlug"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["worker"] = true
			urlDetails["workerPath"] = strings.Join(parts[2:], "/")
			urlDetails["pattern"] = ""
			urlDetails["filePath"] = ""
		} else if ok, _ := regexp.MatchString(customDomainRegex.WithWorker, strings.Join(parts, "/")); ok {
			urlDetails["worker"] = true
			urlDetails["workerPath"] = strings.Join(parts[2:], "/")
			urlDetails["pattern"] = ""
			urlDetails["filePath"] = ""
		} else if ok, _ := regexp.MatchString(customDomainRegex.WithZone, strings.Join(parts, "/")); ok {
			urlDetails["zoneSlug"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["pattern"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["filePath"] = strings.Join(parts[1:], "/")
		} else if ok, _ := regexp.MatchString(customDomainRegex.WithoutZone, strings.Join(parts, "/")); ok {
			urlDetails["pattern"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["filePath"] = strings.Join(parts[1:], "/")
		} else {
			return nil, errors.New("invalid pixelbin url. Please make sure the url is correct")
		}
	} else {
		if ok, _ := regexp.MatchString(VERSION2_REGEX, parts[1]); ok {
			urlDetails["version"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
		}
		if len(parts[1]) < 3 {
			return nil, errors.New("invalid pixelbin url. Please make sure the url is correct")
		}

		if ok, _ := regexp.MatchString(pixelBinDomainRegex.WithWorkerAndZone, strings.Join(parts, "/")); ok {
			urlDetails["cloudName"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["zoneSlug"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["worker"] = true
			urlDetails["workerPath"] = strings.Join(parts[2:], "/")
			urlDetails["pattern"] = ""
			urlDetails["filePath"] = ""
		} else if ok, _ := regexp.MatchString(pixelBinDomainRegex.WithWorker, strings.Join(parts, "/")); ok {
			urlDetails["cloudName"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["worker"] = true
			urlDetails["workerPath"] = strings.Join(parts[2:], "/")
			urlDetails["pattern"] = ""
			urlDetails["filePath"] = ""
		} else if ok, _ := regexp.MatchString(pixelBinDomainRegex.WithZone, strings.Join(parts, "/")); ok {
			urlDetails["cloudName"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["zoneSlug"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["pattern"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["filePath"] = strings.Join(parts[1:], "/")
		} else if ok, _ := regexp.MatchString(pixelBinDomainRegex.WithoutZone, strings.Join(parts, "/")); ok {
			urlDetails["cloudName"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["pattern"] = parts[1]
			parts = append(parts[:1], parts[2:]...)
			urlDetails["filePath"] = strings.Join(parts[1:], "/")
		} else {
			return nil, errors.New("invalid pixelbin url. Please make sure the url is correct")
		}
	}
	return urlDetails, nil
}

func getPartsFromUrl(url string, config urlToObjConfig) (map[string]interface{}, error) {
	parts, err := getUrlParts(url, config)
	if err != nil {
		return nil, err
	}
	queryObj, err := processQueryParams(parts)
	parts["zone"] = nil
	if val, ok := parts["zoneSlug"]; ok {
		parts["zone"] = val
		delete(parts, "zoneSlug")
	}
	parts["baseUrl"] = fmt.Sprintf("%s://%s", parts["protocol"], parts["host"])
	parts["options"] = queryObj
	delete(parts, "protocol")
	delete(parts, "host")
	delete(parts, "search")
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

func getParamsObject(paramsList []string) map[string]string {
	params := make(map[string]string)
	for _, v := range paramsList {
		if strings.Contains(v, ":") {
			it := strings.Split(v, ":")
			if it[0] != "" {
				params[it[0]] = it[1]
			}
		}
	}
	if len(params) > 0 {
		return params
	}
	return nil
}

// previously txtToOptions
func getOperationDetailsFromOperation(dSplit string) map[string]interface{} {
	// Figure Out Module
	fullFnName := strings.Split(dSplit, "(")[0]

	var pluginId string
	var operationName string
	if strings.HasPrefix(dSplit, "p:") {
		arr := strings.Split(fullFnName, ":")
		pluginId = arr[0]
		operationName = arr[1]
	} else {
		arr := strings.Split(fullFnName, ".")
		pluginId = arr[0]
		operationName = arr[1]
	}

	values := make(map[string]string)
	if pluginId == "p" {
		if strings.Contains(dSplit, "(") {
			values = getParamsObject(getParamsList(dSplit, ""))
		}
	} else {
		values = getParamsObject(getParamsList(dSplit, ""))
	}

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

func getTransformationDetailsFromPattern(pattern string, url string) []map[string]interface{} {
	if pattern == "original" || pattern == "" {
		return []map[string]interface{}{}
	}
	dSplit := strings.Split(pattern, OPERTATION_SEPARATOR)
	arr := make([]map[string]interface{}, len(dSplit))
	for i, v := range dSplit {
		// if strings.HasPrefix(v, "p:") {
		// 	a := strings.Split(v, ":")
		// 	v = fmt.Sprintf("p.apply(n:%s)", a[1])
		// }
		result := getOperationDetailsFromOperation(v)
		name := result["name"]
		plugin := result["plugin"]
		var values map[string]string
		if result["values"] != nil && len(result["values"].(map[string]string)) > 0 {
			values = result["values"].(map[string]string)
		}

		if len(values) > 0 {
			keys := make([]string, 0)
			for k := range values {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			valuesList := []map[string]string{}
			for _, k := range keys {
				valuesList = append(valuesList, map[string]string{"key": k, "value": values[k]})
			}

			arr[i] = map[string]interface{}{
				"name":   name,
				"plugin": plugin,
				"values": valuesList,
			}
		} else {
			arr[i] = map[string]interface{}{
				"name":   name,
				"plugin": plugin,
			}
		}
	}

	// opts := FlattenSlice(arr)

	return arr
}

func getObjFromUrl(url string, config urlToObjConfig) (map[string]interface{}, error) {
	parts, err := getPartsFromUrl(url, config)
	if err != nil {
		return nil, errors.New("error Processing url. Please check the url is correct" + err.Error())
	}
	transformations := getTransformationDetailsFromPattern(
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
			if o["plugin"] == "p" {
				if len(paramstr) > 0 {
					newtransformationList = append(newtransformationList, fmt.Sprintf("%s:%s(%s)", o["plugin"], o["name"], paramstr))
				} else {
					newtransformationList = append(newtransformationList, fmt.Sprintf("%s:%s", o["plugin"], o["name"]))
				}
			} else {
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
	if obj["isCustomDomain"] == nil && obj["cloudName"] == nil {
		return "", errors.New("key cloudName should be defined")
	}
	if obj["isCustomDomain"] != nil && obj["isCustomDomain"].(bool) && obj["cloudName"] != nil {
		return "", errors.New("key cloudName is not valid for custom domains")
	}

	var isWorker bool
	
	if obj["worker"] != nil && obj["worker"].(bool) {
		isWorker = true
	} else {
		isWorker = false
	}

	if !isWorker && obj["filePath"] == nil {
		return "", errors.New("key filePath should be defined")
	}
	if isWorker && obj["workerPath"] == nil {
		return "", errors.New("key workerPath should be defined")
	}

	if isWorker {
		obj["pattern"] = "wrkr"
	} else {
		patternValue, err := getPatternFromTransformations(obj["transformations"])
		if err != nil {
			return "", err
		}
		if patternValue != "" {
			obj["pattern"] = patternValue
		} else {
			obj["pattern"] = "original"
		}
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
	if _, ok := obj["zone"]; !ok || obj["zone"] == nil || !regexp.MustCompile(ZONE_SLUG).MatchString(obj["zone"].(string)) {
		obj["zone"] = nil
	}

	urlKeySorted := []string{"baseUrl", "version", "cloudName", "zone", "pattern"}
	if isWorker {
		urlKeySorted = append(urlKeySorted, "workerPath")
	} else {
		urlKeySorted = append(urlKeySorted, "filePath")
	}

	urlArr := []string{}
	for _, key := range urlKeySorted {
		if val, ok := obj[key]; ok && val != nil {
			urlArr = append(urlArr, val.(string))
		}
	}
	queryArr := []string{}
	if _, ok := obj["options"]; ok {
		queryParams := obj["options"].(map[string]interface{})
		if len(queryParams) > 0 {
			if dpr, ok := queryParams["dpr"];  ok && dpr != "" {
				dpr, err := parseDPR(dpr)
				if err != nil {
					return "", err
				}
				queryArr = append(queryArr, "dpr="+fmt.Sprint(dpr))
			}
			if f_auto, ok := queryParams["f_auto"]; ok && f_auto != "" {
				_, err := validateFAuto(f_auto.(bool))
				if err != nil {
					return "", err
				}
				queryArr = append(queryArr, "f_auto="+strconv.FormatBool(f_auto.(bool)))
			}
		}
	}
	urlStr := strings.Join(urlArr, "/")
	if len(queryArr) > 0 {
		urlStr += "?" + strings.Join(queryArr, "&")
	}
	return urlStr, nil
}

func parseDPR(dpr interface{}) (string, error) {
	switch dprValue := dpr.(type) {
    case string:
		if dprValue == "auto" {
			return dpr.(string), nil
		}
	case int:
		if dprValue < 1 || dprValue > 5 {
			return "", errors.New("DPR value should be between 0.1 to 5.0")
		}

		return strconv.Itoa(int(dprValue)), nil
    case float64:
		if dprValue < 0.1 || dprValue > 5.0 {
			return "", errors.New("DPR value should be between 0.1 to 5.0")
		}
	
		return strconv.FormatFloat(dprValue, 'f', 1, 64), nil
	
    default:
        return "", errors.New("Invalid DPR value")
    }

	return "", errors.New("Invalid DPR value")
}

func validateDPR(dpr float64) (map[string]interface{}, error) {
	if reflect.TypeOf(dpr).Kind() != reflect.Float64 || dpr < 0.1 || dpr > 5.0 {
		return nil, errors.New("DPR value should be numeric and should be between 0.1 to 5.0")
	}
	return nil, nil
}

func validateFAuto(f_auto bool) (map[string]interface{}, error) {
	if reflect.TypeOf(f_auto).Kind() != reflect.Bool {
		return nil, errors.New("F_auto value should be boolean")
	}
	return nil, nil
}

func processQueryParams(urlParts map[string]interface{}) (map[string]string, error) {
	queryParams := urlParts["search"].(map[string]string)
	queryObj := map[string]string{}
	if queryParams["dpr"] != "" {
		queryObj["dpr"] = queryParams["dpr"]
	}
	if queryParams["f_auto"] != "" {
		queryObj["f_auto"] = queryParams["f_auto"]
	}
	return queryObj, nil
}
