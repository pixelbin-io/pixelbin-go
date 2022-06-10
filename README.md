# Pixelbin Backend SDK for Golang

Pixelbin Backend SDK for Golang helps you integrate the core Pixelbin features with your application.

## Getting Started

Getting started with Pixelbin Backend SDK for Golang

### Installation

```
go get -u "github.com/pixelbin-dev/pixelbin-go"
```

---

### Usage

#### Quick Example

```go
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
    fmt.Println(result)
}
```

## Utilities

Pixelbin provides url utilities to construct and deconstruct Pixelbin urls.

### UrlToObj

Deconstruct a pixelbin url

| parameter            | description          | example                                                                                               |
| -------------------- | -------------------- | ----------------------------------------------------------------------------------------------------- |
| pixelbinUrl (string) | A valid pixelbin url | `https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg` |

**Returns**:

| property                | description                            | example                    |
| ----------------------- | -------------------------------------- | -------------------------- |
| cloudName (string)      | The cloudname extracted from the url   | `your-cloud-name`          |
| zone (string)           | 6 character zone slug                  | `z-slug`                   |
| version (string)        | cdn api version                        | `v2`                       |
| transformations (array) | Extracted transformations from the url |                            |
| filePath                | Path to the file on Pixelbin storage   | `/path/to/image.jpeg`      |
| baseUrl (string)        | Base url                               | `https://cdn.pixelbin.io/` |

Example:

```golang
import (
	"fmt"
	"os"
	"github.com/pixelbin-dev/pixelbin-go/sdk/utils/url"
)

func main() {
    pixelbinUrl := "https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg"
    obj := url.UrlToObj(pixelbinUrl)
}
// obj
// {
//     "cloudName": "your-cloud-name",
//     "zone": "z-slug",
//     "version": "v2",
//     "transformations": [
//         {
//             "plugin": "t",
//             "name": "resize",
//             "values": [
//                 {
//                     "key": "h",
//                     "value": "100"
//                 },
//                 {
//                     "key": "w",
//                     "value": "200"
//                 }
//             ]
//         },
//         {
//             "plugin": "t",
//             "name": "flip",
//         }
//     ],
//     "filePath": "path/to/image.jpeg",
//     "baseUrl": "https://cdn.pixelbin.io"
// }
```

### ObjToUrl

Converts the extracted url obj to a Pixelbin url.

| property                | description                            | example                    |
| ----------------------- | -------------------------------------- | -------------------------- |
| cloudName (string)      | The cloudname extracted from the url   | `your-cloud-name`          |
| zone (string)           | 6 character zone slug                  | `z-slug`                   |
| version (string)        | cdn api version                        | `v2`                       |
| transformations (array) | Extracted transformations from the url |                            |
| filePath                | Path to the file on Pixelbin storage   | `/path/to/image.jpeg`      |
| baseUrl (string)        | Base url                               | `https://cdn.pixelbin.io/` |

```golang
import (
	"fmt"
	"os"
	"github.com/pixelbin-dev/pixelbin-go/sdk/utils/url"
)
func main() {
    obj := map[string]interface{}{
        cloudName: "your-cloud-name",
        zone: "z-slug",
        version: "v2",
        transformations: [
            {
                plugin: "t",
                name: "resize",
                values: [
                    {
                        key: "h",
                        value: "100",
                    },
                    {
                        key: "w",
                        value: "200",
                    },
                ],
            },
            {
                plugin: "t",
                name: "flip",
            },
        ],
        filePath: "path/to/image.jpeg",
        baseUrl: "https://cdn.pixelbin.io",
    }
    urlstring := url.ObjToUrl(obj) // obj is as shown above
}
// urlstring
// https://cdn.pixelbin.io/v2/your-cloud-name/z-slug/t.resize(h:100,w:200)~t.flip()/path/to/image.jpeg
```

## Documentation

-   [API docs](documentation/platform/README.md)
