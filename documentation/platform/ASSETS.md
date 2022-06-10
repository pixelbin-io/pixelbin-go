##### [Back to Pixelbin API docs](./README.md)

## Assets Methods

Asset Uploader Service

-   [FileUpload](#fileupload)
-   [UrlUpload](#urlupload)
-   [CreateSignedUrl](#createsignedurl)
-   [ListFiles](#listfiles)
-   [ListFilesPaginator](#listfilespaginator)
-   [GetFileById](#getfilebyid)
-   [GetFileByFileId](#getfilebyfileid)
-   [UpdateFile](#updatefile)
-   [DeleteFile](#deletefile)
-   [DeleteFiles](#deletefiles)
-   [CreateFolder](#createfolder)
-   [UpdateFolder](#updatefolder)
-   [DeleteFolder](#deletefolder)
-   [GetModules](#getmodules)
-   [GetModule](#getmodule)

## Methods with example and description

### FileUpload

**Summary**: Upload File

```golang
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

    // Parameters for FileUpload function
    params := platform.FileUploadXQuery{
        File: os.Open("your-file-path"),
        Path: "path/to/containing/folder",
        Name: "filename",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
    }
    result, err := pixelbin.Assets.FileUpload(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| File             | \*os.File              | yes      | Asset file                                                                                                                                                                                                                       |
| Path             | string                 | no       | Path where you want to store the asset. Path of containing folder                                                                                                                                                                |
| Name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| Metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

Upload File to Pixelbin

_Returned Response:_

[UploadResponse](#uploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### UrlUpload

**Summary**: Upload Asset with url

```golang
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

    // Parameters for FileUpload function
    params := platform.UrlUploadXQuery{
        URL: "www.dummy.com/image.png",
        Path: "path/to/containing/folder",
        Name: "filename",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
    }
    result, err := pixelbin.Assets.UrlUpload(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| URL              | string                 | yes      | Asset URL                                                                                                                                                                                                                        |
| Path             | string                 | no       | Path where you want to store the asset. Path of containing folder.                                                                                                                                                               |
| Name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| Metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

Upload Asset with url

_Returned Response:_

[UploadResponse](#uploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### CreateSignedUrl

**Summary**: S3 Signed URL upload

```golang
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

    // Parameters for FileUpload function
    params := platform.CreateSignedUrlXQuery{
        Name: "filename",
        Path: "path/to/containing/folder",
        Format: "jpeg",
        Access: "public-read",
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{},
        Overwrite: false,
        FilenameOverride: true,
    }
    result, err := pixelbin.Assets.CreateSignedUrl(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument         | Type                   | Required | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Name             | string                 | no       | name of the file                                                                                                                                                                                                                 |
| Path             | string                 | no       | Path of containing folder.                                                                                                                                                                                                       |
| Format           | string                 | no       | Format of the file                                                                                                                                                                                                               |
| Access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| Tags             | []string               | no       | Tags associated with the file.                                                                                                                                                                                                   |
| Metadata         | map[string]interface{} | no       | Metadata associated with the file.                                                                                                                                                                                               |
| Overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| FilenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

For the given asset details, a S3 signed URL will be generated,
which can be then used to upload your asset.

_Returned Response:_

[SignedUploadResponse](#signeduploadresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "s3PresignedUrl": {
        "url": "https://domain.com/xyz",
        "fields": {
            "field1": "value",
            "field2": "value"
        }
    }
}
```

</details>

### ListFiles

**Summary**: List and search files and folders.

```golang
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

    // Parameters for FileUpload function
    params := platform.ListFilesXQuery{
        Name: "cat",
        Path: "cat-photos",
        Format: "jpeg",
        Tags: ["cats","animals"],
        OnlyFiles: "false",
        OnlyFolders: "false",
        PageNo: 1,
        PageSize: 10,
        Sort: "name",
    }
    result, err := pixelbin.Assets.ListFiles(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument    | Type     | Required | Description                                                                  |
| ----------- | -------- | -------- | ---------------------------------------------------------------------------- |
| Name        | string   | no       | Find items with matching name                                                |
| Path        | string   | no       | Find items with matching path                                                |
| Format      | string   | no       | Find items with matching format                                              |
| Tags        | []string | no       | Find items containing these tags                                             |
| OnlyFiles   | bool     | no       | If true will fetch only files                                                |
| OnlyFolders | bool     | no       | If true will fetch only folders                                              |
| PageNo      | float64  | no       | Page No.                                                                     |
| PageSize    | float64  | no       | Page Size                                                                    |
| Sort        | string   | no       | Key to sort results by. A "-" suffix will sort results in descending orders. |

List all files and folders in root folder. Search for files if name is provided. If path is provided, search in the specified path.

_Returned Response:_

[ListFilesResponse](#listfilesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "items": [
        {
            "_id": "dummy-uuid",
            "name": "dir",
            "type": "folder"
        },
        {
            "_id": "dummy-uuid",
            "name": "asset2",
            "type": "file",
            "path": "dir",
            "fileId": "dir/asset2",
            "format": "jpeg",
            "size": 1000,
            "access": "private"
        },
        {
            "_id": "dummy-uuid",
            "name": "asset1",
            "type": "file",
            "path": "dir",
            "fileId": "dir/asset1",
            "format": "jpeg",
            "size": 1000,
            "access": "private"
        }
    ],
    "page": {
        "type": "number",
        "size": 4,
        "current": 1,
        "hasNext": false
    }
}
```

</details>

### GetFileById

**Summary**: Get file details with \_id

```golang
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

    // Parameters for FileUpload function
    params := platform.GetFileByIdXQuery{
        ID: "c9138153-94ea-4dbe-bea9-65d43dba85ae",
    }
    result, err := pixelbin.Assets.GetFileById(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description  |
| -------- | ------ | -------- | ------------ |
| ID       | string | yes      | \_id of File |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### GetFileByFileId

**Summary**: Get file details with fileId

```golang
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

    // Parameters for FileUpload function
    params := platform.GetFileByFileIdXQuery{
        FileId: "path/to/file/name",
    }
    result, err := pixelbin.Assets.GetFileByFileId(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                              |
| -------- | ------ | -------- | ---------------------------------------- |
| FileId   | string | yes      | Combination of `path` and `name` of file |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### UpdateFile

**Summary**: Update file details

```golang
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

    // Parameters for FileUpload function
    params := platform.UpdateFileXQuery{
        FileId: "path/to/file/name",,
        Name: "asset",
        Path: "dir",
        Access: "private",
        IsActive: false,
        Tags: []string{"tag1","tag2"},
        Metadata: map[string]interface{}{"key":"value"},
    }
    result, err := pixelbin.Assets.UpdateFile(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type                   | Required | Description                                                     |
| -------- | ---------------------- | -------- | --------------------------------------------------------------- |
| FileId   | string                 | yes      | Combination of `path` and `name`                                |
| Name     | string                 | no       | Name of the file                                                |
| Path     | string                 | no       | path of containing folder.                                      |
| Access   | string                 | no       | Access level of asset, can be either `public-read` or `private` |
| IsActive | bool                   | no       | Whether the file is active                                      |
| Tags     | []string               | no       | Tags associated with the file                                   |
| Metadata | map[string]interface{} | no       | Metadata associated with the file                               |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### DeleteFile

**Summary**: Delete file

```golang
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

    // Parameters for FileUpload function
    params := platform.DeleteFileXQuery{
        FileId: "path/to/file/name",
    }
    result, err := pixelbin.Assets.DeleteFile(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                      |
| -------- | ------ | -------- | -------------------------------- |
| FileId   | string | yes      | Combination of `path` and `name` |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### DeleteFiles

**Summary**: Delete multiple files

```golang
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

    // Parameters for FileUpload function
    params := platform.DeleteFilesXQuery{
        Ids: []string{"_id_1","_id_2","_id_3"},
    }
    result, err := pixelbin.Assets.DeleteFiles(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type     | Required | Description                   |
| -------- | -------- | -------- | ----------------------------- |
| Ids      | []string | yes      | Array of file \_ids to delete |

_Returned Response:_

[FilesResponse](#filesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "asset",
    "path": "dir",
    "fileId": "dir/asset",
    "format": "jpeg",
    "size": 1000,
    "access": "private",
    "isActive": true,
    "tags": ["tag1", "tag2"],
    "metadata": {
        "key": "value"
    },
    "url": "https://domain.com/filename.jpeg"
}
```

</details>

### CreateFolder

**Summary**: Create folder

```golang
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

    // Parameters for FileUpload function
    params := platform.CreateFolderXQuery{
        Name: "subDir",
        Path: "dir",
    }
    result, err := pixelbin.Assets.CreateFolder(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                |
| -------- | ------ | -------- | -------------------------- |
| Name     | string | yes      | Name of the folder         |
| Path     | string | no       | path of containing folder. |

Create a new folder at the specified path. Also creates the ancestors if they do not exist.

_Returned Response:_

[[]FoldersResponse](#[foldersresponse)

Success - List of all created folders

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
[
    {
        "_id": "dummy-uuid",
        "name": "subDir",
        "path": "dir",
        "isActive": true
    },
    {
        "_id": "dummy-uuid-2",
        "name": "subDir",
        "path": "dir",
        "isActive": true
    }
]
```

</details>

### UpdateFolder

**Summary**: Update folder details

```golang
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

    // Parameters for FileUpload function
    params := platform.UpdateFolderXQuery{
        FolderId: "path/to/folder/name",,
        IsActive: false,
    }
    result, err := pixelbin.Assets.UpdateFolder(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                      |
| -------- | ------ | -------- | -------------------------------- |
| FolderId | string | yes      | combination of `path` and `name` |
| IsActive | bool   | no       | whether the folder is active     |

Update folder details. Eg: Soft delete it
by making `isActive` as `false`.
We currently do not support updating folder name or path.

_Returned Response:_

[FoldersResponse](#foldersresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "_id": "dummy-uuid",
    "name": "subDir",
    "path": "dir",
    "isActive": true
}
```

</details>

### DeleteFolder

**Summary**: Delete folder

```golang
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

    // Parameters for FileUpload function
    params := platform.DeleteFolderXQuery{
        ID: "c9138153-94ea-4dbe-bea9-65d43dba85ae",
    }
    result, err := pixelbin.Assets.DeleteFolder(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument | Type   | Required | Description                  |
| -------- | ------ | -------- | ---------------------------- |
| ID       | string | yes      | \_id of folder to be deleted |

Delete folder and all its children permanently.

_Returned Response:_

[[]FoldersResponse](#[foldersresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
[
    {
        "_id": "dummy-uuid",
        "name": "subDir",
        "path": "dir",
        "isActive": true
    },
    {
        "_id": "dummy-uuid-2",
        "name": "subDir",
        "path": "dir",
        "isActive": true
    }
]
```

</details>

### GetModules

**Summary**: Get all transformation modules

```golang
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

    // Parameters for FileUpload function
    params := platform.GetModulesXQuery{
    }
    result, err := pixelbin.Assets.GetModules(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

Get all transformation modules.

_Returned Response:_

[TransformationModulesResponse](#transformationmodulesresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "delimiters": {
        "operationSeparator": "~",
        "parameterSeparator": ":"
    },
    "plugins": {
        "erase": {
            "identifier": "erase",
            "name": "EraseBG",
            "description": "EraseBG Background Removal Module",
            "credentials": {
                "required": false
            },
            "operations": [
                {
                    "params": {
                        "name": "Industry Type",
                        "type": "enum",
                        "enum": ["general", "ecommerce"],
                        "default": "general",
                        "identifier": "i",
                        "title": "Industry type"
                    },
                    "displayName": "Remove background of an image",
                    "method": "bg",
                    "description": "Remove the background of any image"
                }
            ],
            "enabled": true
        }
    },
    "presets": [
        {
            "_id": "dummy-id",
            "createdAt": "2022-02-14T10:06:17.803Z",
            "updatedAt": "2022-02-14T10:06:17.803Z",
            "isActive": true,
            "orgId": "265",
            "presetName": "compressor",
            "transformation": "t.compress(q:95)",
            "archived": false
        }
    ]
}
```

</details>

### GetModule

**Summary**: Get Transformation Module by module identifier

```golang
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

    // Parameters for FileUpload function
    params := platform.GetModuleXQuery{
        Identifier: "t",
    }
    result, err := pixelbin.Assets.GetModule(params)

    if err != nil {
        fmt.Println(err)
    }
    // use result
    fmt.Println(result)
}

```

| Argument   | Type   | Required | Description                         |
| ---------- | ------ | -------- | ----------------------------------- |
| Identifier | string | yes      | identifier of Transformation Module |

Get Transformation Module by module identifier

_Returned Response:_

[TransformationModuleResponse](#transformationmoduleresponse)

Success

<details>
<summary><i>&nbsp; Example:</i></summary>

```json
{
    "identifier": "erase",
    "name": "EraseBG",
    "description": "EraseBG Background Removal Module",
    "credentials": {
        "required": false
    },
    "operations": [
        {
            "params": {
                "name": "Industry Type",
                "type": "enum",
                "enum": ["general", "ecommerce"],
                "default": "general",
                "identifier": "i",
                "title": "Industry type"
            },
            "displayName": "Remove background of an image",
            "method": "bg",
            "description": "Remove the background of any image"
        }
    ],
    "enabled": true
}
```

</details>

### Schemas

#### folderItem

| Properties | Type   | Nullable | Description                          |
| ---------- | ------ | -------- | ------------------------------------ |
| \_id       | string | yes      | Id of the folder item                |
| name       | string | yes      | Name of the folder item              |
| path       | string | yes      | Path of containing folder            |
| type       | string | yes      | Type of the item. `file` or `folder` |

#### exploreItem

| Properties | Type    | Nullable | Description                                                     |
| ---------- | ------- | -------- | --------------------------------------------------------------- |
| \_id       | string  | yes      | id of the exploreItem                                           |
| name       | string  | yes      | name of the item                                                |
| type       | string  | yes      | Type of item whether `file` or `folder`                         |
| path       | string  | yes      | Path of containing folder                                       |
| fileId     | string  | no       | Combination of `path` and `name` of file                        |
| format     | string  | no       | Format of the file                                              |
| size       | float64 | no       | Size of the file in bytes                                       |
| access     | string  | no       | Access level of asset, can be either `public-read` or `private` |

#### page

| Properties | Type    | Nullable | Description                   |
| ---------- | ------- | -------- | ----------------------------- |
| type       | string  | yes      | Type of page                  |
| size       | float64 | yes      | Number of items on the page   |
| current    | float64 | yes      | Current page number.          |
| hasNext    | bool    | yes      | Whether the next page exists. |
| itemTotal  | float64 | yes      | Total number of items.        |

#### exploreResponse

| Properties | Type          | Nullable | Description                  |
| ---------- | ------------- | -------- | ---------------------------- |
| items      | []exploreItem | yes      | exploreItems in current page |
| page       | page          | yes      | page details                 |

#### ListFilesResponse

| Properties | Type          | Nullable | Description                  |
| ---------- | ------------- | -------- | ---------------------------- |
| items      | []exploreItem | yes      | exploreItems in current page |
| page       | page          | yes      | page details                 |

#### exploreFolderResponse

| Properties | Type          | Nullable | Description                  |
| ---------- | ------------- | -------- | ---------------------------- |
| folder     | folderItem    | yes      | requested folder item        |
| items      | []exploreItem | yes      | exploreItems in current page |
| page       | page          | yes      | page details                 |

#### FileUploadRequest

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| file             | \*os.File              | yes      | Asset file                                                                                                                                                                                                                       |
| path             | string                 | no       | Path where you want to store the asset. Path of containing folder                                                                                                                                                                |
| name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

#### UrlUploadRequest

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| url              | string                 | yes      | Asset URL                                                                                                                                                                                                                        |
| path             | string                 | no       | Path where you want to store the asset. Path of containing folder.                                                                                                                                                               |
| name             | string                 | no       | Name of the asset, if not provided name of the file will be used. Note - The provided name will be slugified to make it URL safe                                                                                                 |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Asset tags                                                                                                                                                                                                                       |
| metadata         | map[string]interface{} | no       | Asset related metadata                                                                                                                                                                                                           |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

#### UploadResponse

| Properties | Type                   | Nullable | Description                                                 |
| ---------- | ---------------------- | -------- | ----------------------------------------------------------- |
| \_id       | string                 | yes      | \_id of the item                                            |
| fileId     | string                 | yes      | Combination of `path` and `name` of file                    |
| name       | string                 | yes      | name of the item                                            |
| path       | string                 | yes      | path to the parent folder                                   |
| format     | string                 | yes      | format of the file                                          |
| size       | float64                | yes      | size of file in bytes                                       |
| access     | string                 | yes      | Access level of asset, can be either public-read or private |
| tags       | []string               | no       | tags associated with the item                               |
| metadata   | map[string]interface{} | no       | metadata associated with the item                           |
| url        | string                 | no       | url of the item                                             |
| thumbnail  | string                 | no       | url of item thumbnail                                       |

#### SignedUploadRequest

| Properties       | Type                   | Nullable | Description                                                                                                                                                                                                                      |
| ---------------- | ---------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name             | string                 | no       | name of the file                                                                                                                                                                                                                 |
| path             | string                 | no       | Path of containing folder.                                                                                                                                                                                                       |
| format           | string                 | no       | Format of the file                                                                                                                                                                                                               |
| access           | AccessEnum             | no       | Access level of asset, can be either `public-read` or `private`                                                                                                                                                                  |
| tags             | []string               | no       | Tags associated with the file.                                                                                                                                                                                                   |
| metadata         | map[string]interface{} | no       | Metadata associated with the file.                                                                                                                                                                                               |
| overwrite        | bool                   | no       | Overwrite flag. If set to `true` will overwrite any file that exists with same path, name and type. Defaults to `false`.                                                                                                         |
| filenameOverride | bool                   | no       | If set to `true` will add unique characters to name if asset with given name already exists. If overwrite flag is set to `true`, preference will be given to overwrite flag. If both are set to `false` an error will be raised. |

#### SignedUploadResponse

| Properties     | Type         | Nullable | Description                                  |
| -------------- | ------------ | -------- | -------------------------------------------- |
| s3PresignedUrl | PresignedUrl | yes      | `signedDetails` for upload with frontend sdk |
|  |

#### PresignedUrl

| Properties | Type                   | Nullable | Description                                 |
| ---------- | ---------------------- | -------- | ------------------------------------------- |
| url        | string                 | no       | `presigned url` for upload                  |
|  |
| fields     | map[string]interface{} | no       | signed fields to be sent along with request |

#### FilesResponse

| Properties | Type                   | Nullable | Description                                                    |
| ---------- | ---------------------- | -------- | -------------------------------------------------------------- |
| \_id       | string                 | yes      | \_id of the file                                               |
| name       | string                 | yes      | name of the file                                               |
| path       | string                 | yes      | path of containing folder.                                     |
| fileId     | string                 | yes      | Combination of `path` and `name` of file                       |
| format     | string                 | yes      | format of the file                                             |
| size       | float64                | yes      | size of the file in bytes                                      |
| access     | string                 | yes      | Access level of file, can be either `public-read` or `private` |
| isActive   | bool                   | yes      | Whether the file is active                                     |
| tags       | []string               | no       | Tags associated with the file                                  |
| metadata   | map[string]interface{} | no       | Metadata associated with the file                              |
| url        | string                 | no       | url of the file                                                |
| thumbnail  | string                 | no       | url of the thumbnail of the file                               |

#### UpdateFileRequest

| Properties | Type                   | Nullable | Description                                                     |
| ---------- | ---------------------- | -------- | --------------------------------------------------------------- |
| name       | string                 | no       | Name of the file                                                |
| path       | string                 | no       | path of containing folder.                                      |
| access     | string                 | no       | Access level of asset, can be either `public-read` or `private` |
| isActive   | bool                   | no       | Whether the file is active                                      |
| tags       | []string               | no       | Tags associated with the file                                   |
| metadata   | map[string]interface{} | no       | Metadata associated with the file                               |

#### FoldersResponse

| Properties | Type   | Nullable | Description                  |
| ---------- | ------ | -------- | ---------------------------- |
| \_id       | string | yes      | \_id of the folder           |
| name       | string | yes      | name of the folder           |
| path       | string | yes      | path of containing folder.   |
| isActive   | bool   | yes      | whether the folder is active |

#### CreateFolderRequest

| Properties | Type   | Nullable | Description                |
| ---------- | ------ | -------- | -------------------------- |
| name       | string | yes      | Name of the folder         |
| path       | string | no       | path of containing folder. |

#### UpdateFolderRequest

| Properties | Type | Nullable | Description                  |
| ---------- | ---- | -------- | ---------------------------- |
| isActive   | bool | no       | whether the folder is active |

#### TransformationModulesResponse

| Properties | Type                                    | Nullable | Description                                         |
| ---------- | --------------------------------------- | -------- | --------------------------------------------------- |
| delimiters | Delimiter                               | no       | Delimiter for parsing plugin schema                 |
| plugins    | map[string]TransformationModuleResponse | no       | Transformations currently supported by the pixelbin |
| presets    | []interface{}                           | no       | List of saved presets                               |

#### DeleteMultipleFilesRequest

| Properties | Type     | Nullable | Description                   |
| ---------- | -------- | -------- | ----------------------------- |
| ids        | []string | yes      | Array of file \_ids to delete |

#### Delimiter

| Properties         | Type   | Nullable | Description                                                              |
| ------------------ | ------ | -------- | ------------------------------------------------------------------------ |
| operationSeparator | string | no       | separator to separate operations in the url pattern                      |
| parameterSeparator | string | no       | separator to separate parameters used with operations in the url pattern |

#### TransformationModuleResponse

| Properties  | Type                   | Nullable | Description                                     |
| ----------- | ---------------------- | -------- | ----------------------------------------------- |
| identifier  | string                 | no       | identifier for the plugin type                  |
| name        | string                 | no       | name of the plugin                              |
| description | string                 | no       | description of the plugin                       |
| credentials | map[string]interface{} | no       | credentials, if any, associated with the plugin |
| operations  | []interface{}          | no       | supported operations in the plugin              |
| enabled     | bool                   | no       | whether the plugin is enabled                   |

### Enums

#### [AccessEnum](#AccessEnum)

Type : string

| Name        | Value       | Description                                |
| ----------- | ----------- | ------------------------------------------ |
| public-read | public-read | Object is available for public read access |
| private     | private     | Object is private                          |

---
