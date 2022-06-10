package tests

import (
	"fmt"
	"testing"

	"github.com/pixelbin-dev/pixelbin-go/sdk/utils/url"
)

var urls_and_obj = []map[string]interface{}{
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)~t.rotate(a:-249)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{{"key": "a", "value": "-249"}},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{{"key": "a", "value": "-249"}},
				},
				{"plugin": "t", "name": "flip"},
				{
					"plugin": "t",
					"name":   "trim",
					"values": []map[string]interface{}{{"key": "t", "value": "217"}},
				},
			},
		},
	},
}

func TestUrlToObjCases(t *testing.T) {
	for i, cases := range urls_and_obj {
		urlstr := cases["url"].(string)
		expObj := cases["obj"].(map[string]interface{})
		obj, err := url.UrlToObj(urlstr)
		if err != nil {
			t.Errorf("Failed case %d ! got err %v", i+1, err)
		} else {
			if fmt.Sprint(obj) == fmt.Sprint(expObj) {
				t.Logf("Success case %d", i+1)
			} else {
				t.Errorf("Failed ! caes %d expected %v, got %v", i+1, expObj, obj)
			}
		}
	}
}

func TestObjToUrlCases(t *testing.T) {
	for i, cases := range urls_and_obj {
		expurlstr := cases["url"].(string)
		Obj := cases["obj"].(map[string]interface{})
		url, err := url.ObjToUrl(Obj)
		if err != nil {
			t.Errorf("Failed case %d ! got err %v", i+1, err)
		} else {
			if url == expurlstr {
				t.Logf("Success case %d", i+1)
			} else {
				t.Errorf("Failed ! caes %d expected %v, got %v", i+1, expurlstr, url)
			}
		}
	}
}
