package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pixelbin-dev/pixelbin-go/sdk/utils/url"
)

var urls_to_obj = []map[string]interface{}{
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
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
			"options":   map[string]interface{}{},
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
			"options":   map[string]interface{}{},
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
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "100"},
						{"key": "b", "value": "2.1"},
						{"key": "c", "value": "test"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1()/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1()",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg?dpr=2.5&f_auto=true",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"options": map[string]interface{}{
				"dpr":    2.5,
				"f_auto": true,
			},
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
}

var obj_to_urls = []map[string]interface{}{
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
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
			"options":   map[string]interface{}{},
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
			"options":   map[string]interface{}{},
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
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "100"},
						{"key": "b", "value": "2.1"},
						{"key": "c", "value": "test"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1(a:12)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1()",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "12"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1()",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": make([]map[string]interface{}, 0),
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbinx0.de/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg?dpr=2.5&f_auto=true",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbinx0.de",
			"options": map[string]interface{}{
				"dpr":    2.5,
				"f_auto": true,
			},
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
}

func TestUrlToObjCases(t *testing.T) {
	for i, cases := range urls_to_obj {
		urlstr := cases["url"].(string)
		expObj := cases["obj"].(map[string]interface{})
		obj, err := url.UrlToObj(urlstr)
		if err != nil {
			t.Errorf("Failed case %d ! got err %v", i+1, err)
		} else {
			if fmt.Sprint(obj) == fmt.Sprint(expObj) {
				t.Logf("Success case %d", i+1)
			} else {
				t.Errorf("Failed ! case %d expected %v, got %v", i+1, expObj, obj)
			}
		}
	}
}

func TestObjToUrlCases(t *testing.T) {
	for i, cases := range obj_to_urls {
		expurlstr := cases["url"].(string)
		Obj := cases["obj"].(map[string]interface{})
		url, err := url.ObjToUrl(Obj)
		if err != nil {
			t.Errorf("Failed case %d ! got err %v", i+1, err)
		} else {
			if url == expurlstr {
				t.Logf("Success case %d", i+1)
			} else {
				t.Errorf("Failed ! case %d expected %v, got %v", i+1, expurlstr, url)
			}
		}
	}
}

func TestFailureOptionDprOfQueryParam(t *testing.T) {
	Obj := map[string]interface{}{
		"baseUrl":   "https://cdn.pixelbin.io",
		"filePath":  "__playground/playground-default.jpeg",
		"version":   "v2",
		"zone":      "z-slug",
		"cloudName": "red-scene-95b6ea",
		"pattern":   nil,
		"options": map[string]interface{}{
			"dpr":    5.5,
			"f_auto": true,
		},
		"transformations": []map[string]interface{}{},
	}
	_, err := url.ObjToUrl(Obj)
	if err != nil {
		errIdentifier := strings.Split(err.Error(), " ")[0]
		if errIdentifier == "DPR" {
			t.Logf("Success")
		} else {
			t.Errorf("Failed ! got err %v", err)
		}
	} else {
		t.Errorf("Failed case")
	}
}

func TestFailureOptionFautoOfQueryParam(t *testing.T) {
	Obj := map[string]interface{}{
		"baseUrl":   "https://cdn.pixelbin.io",
		"filePath":  "__playground/playground-default.jpeg",
		"version":   "v2",
		"zone":      "z-slug",
		"cloudName": "red-scene-95b6ea",
		"pattern":   nil,
		"options": map[string]interface{}{
			"dpr":    5.5,
			"f_auto": true,
		},
		"transformations": []map[string]interface{}{},
	}
	_, err := url.ObjToUrl(Obj)
	if err != nil {
		errIdentifier := strings.Split(err.Error(), " ")[0]
		if errIdentifier == "DPR" {
			t.Logf("Success")
		} else {
			t.Errorf("Failed ! got err %v", err)
		}
	} else {
		t.Errorf("Failed case")
	}
}
