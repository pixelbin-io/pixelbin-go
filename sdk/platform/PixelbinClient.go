package platform

import (
	"encoding/json"
	"fmt"
	"github.com/pixelbin-dev/pixelbin-go/sdk/common"
	"os"
)

// PixelbinClient holds all the PixelbinConfig object properties
type PixelbinClient struct {
	Config       *PixelbinConfig
	Assets       *Assets
	Organization *Organization
}

// NewPixelbinClient returns pixelbin service instances
func NewPixelbinClient(config *PixelbinConfig) *PixelbinClient {
	return &PixelbinClient{
		Config:       config,
		Assets:       NewAssets(config),
		Organization: NewOrganization(config),
	}
}

// Assets holds Assets object properties
type Assets struct {
	config *PixelbinConfig
}

// NewAssets returns new Assets instance
func NewAssets(config *PixelbinConfig) *Assets {
	return &Assets{config: config}
}

type FileUploadXQuery struct {
	File             *os.File               `json:"file,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
}

/*
summary: Upload File

description: Upload File to Pixelbin

params: FileUploadXQuery
*/
func (c *Assets) FileUpload(
	p FileUploadXQuery,
) (map[string]interface{}, error) {

	type body struct {
		File *os.File `json:"file,omitempty"`

		Path string `json:"path,omitempty"`

		Name string `json:"name,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`
	}
	bodydata := &body{

		File: p.File,

		Path: p.Path,

		Name: p.Name,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/upload/direct",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "multipart/form-data",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UrlUploadXQuery struct {
	URL              string                 `json:"url,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
}

/*
summary: Upload Asset with url

description: Upload Asset with url

params: UrlUploadXQuery
*/
func (c *Assets) UrlUpload(
	p UrlUploadXQuery,
) (map[string]interface{}, error) {

	type body struct {
		URL string `json:"url,omitempty"`

		Path string `json:"path,omitempty"`

		Name string `json:"name,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`
	}
	bodydata := &body{

		URL: p.URL,

		Path: p.Path,

		Name: p.Name,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/upload/url",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type CreateSignedUrlXQuery struct {
	Name             string                 `json:"name,omitempty"`
	Path             string                 `json:"path,omitempty"`
	Format           string                 `json:"format,omitempty"`
	Access           AccessEnum             `json:"access,omitempty"`
	Tags             []string               `json:"tags,omitempty"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
	Overwrite        bool                   `json:"overwrite,omitempty"`
	FilenameOverride bool                   `json:"filenameOverride,omitempty"`
}

/*
summary: S3 Signed URL upload

description: For the given asset details, a S3 signed URL will be generated,
which can be then used to upload your asset.


params: CreateSignedUrlXQuery
*/
func (c *Assets) CreateSignedUrl(
	p CreateSignedUrlXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		Format string `json:"format,omitempty"`

		Access AccessEnum `json:"access,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`

		Overwrite bool `json:"overwrite,omitempty"`

		FilenameOverride bool `json:"filenameOverride,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,

		Format: p.Format,

		Access: p.Access,

		Tags: p.Tags,

		Metadata: p.Metadata,

		Overwrite: p.Overwrite,

		FilenameOverride: p.FilenameOverride,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/upload/signed-url",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type ListFilesXQuery struct {
	Name        string        `json:"name,omitempty"`
	Path        string        `json:"path,omitempty"`
	Format      string        `json:"format,omitempty"`
	Tags        []interface{} `json:"tags,omitempty"`
	OnlyFiles   bool          `json:"onlyFiles,omitempty"`
	OnlyFolders bool          `json:"onlyFolders,omitempty"`
	PageNo      float64       `json:"pageNo,omitempty"`
	PageSize    float64       `json:"pageSize,omitempty"`
	Sort        string        `json:"sort,omitempty"`
}

/*
summary: List and search files and folders.

description: List all files and folders in root folder. Search for files if name is provided. If path is provided, search in the specified path.


params: ListFilesXQuery
*/
func (c *Assets) ListFiles(
	p ListFilesXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	if p.Name != "" {
		queryParams["name"] = fmt.Sprintf("%v", p.Name)
	}

	if p.Path != "" {
		queryParams["path"] = fmt.Sprintf("%v", p.Path)
	}

	if p.Format != "" {
		queryParams["format"] = fmt.Sprintf("%v", p.Format)
	}

	if p.Tags != nil {
		queryParams["tags"] = fmt.Sprintf("%v", p.Tags)
	}

	if p.OnlyFiles != false {
		queryParams["onlyFiles"] = fmt.Sprintf("%v", p.OnlyFiles)
	}

	if p.OnlyFolders != false {
		queryParams["onlyFolders"] = fmt.Sprintf("%v", p.OnlyFolders)
	}

	if p.PageNo != 0 {
		queryParams["pageNo"] = fmt.Sprintf("%v", p.PageNo)
	}

	if p.PageSize != 0 {
		queryParams["pageSize"] = fmt.Sprintf("%v", p.PageSize)
	}

	if p.Sort != "" {
		queryParams["sort"] = fmt.Sprintf("%v", p.Sort)
	}

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/listFiles",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetFileByIdXQuery struct {
	ID string `json:"_id,omitempty"`
}

/*
summary: Get file details with _id

description:

params: GetFileByIdXQuery
*/
func (c *Assets) GetFileById(
	p GetFileByIdXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/id/%s", p.ID),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetFileByFileIdXQuery struct {
	FileId string `json:"fileId,omitempty"`
}

/*
summary: Get file details with fileId

description:

params: GetFileByFileIdXQuery
*/
func (c *Assets) GetFileByFileId(
	p GetFileByFileIdXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/%s", p.FileId),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UpdateFileXQuery struct {
	FileId   string                 `json:"fileId,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Path     string                 `json:"path,omitempty"`
	Access   string                 `json:"access,omitempty"`
	IsActive bool                   `json:"isActive,omitempty"`
	Tags     []string               `json:"tags,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

/*
summary: Update file details

description:

params: UpdateFileXQuery
*/
func (c *Assets) UpdateFile(
	p UpdateFileXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`

		Access string `json:"access,omitempty"`

		IsActive bool `json:"isActive,omitempty"`

		Tags []string `json:"tags,omitempty"`

		Metadata map[string]interface{} `json:"metadata,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,

		Access: p.Access,

		IsActive: p.IsActive,

		Tags: p.Tags,

		Metadata: p.Metadata,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "patch",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/%s", p.FileId),
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteFileXQuery struct {
	FileId string `json:"fileId,omitempty"`
}

/*
summary: Delete file

description:

params: DeleteFileXQuery
*/
func (c *Assets) DeleteFile(
	p DeleteFileXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "delete",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/files/%s", p.FileId),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteFilesXQuery struct {
	Ids []string `json:"ids,omitempty"`
}

/*
summary: Delete multiple files

description:

params: DeleteFilesXQuery
*/
func (c *Assets) DeleteFiles(
	p DeleteFilesXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Ids []string `json:"ids,omitempty"`
	}
	bodydata := &body{

		Ids: p.Ids,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/files/delete",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type CreateFolderXQuery struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}

/*
summary: Create folder

description: Create a new folder at the specified path. Also creates the ancestors if they do not exist.


params: CreateFolderXQuery
*/
func (c *Assets) CreateFolder(
	p CreateFolderXQuery,
) (map[string]interface{}, error) {

	type body struct {
		Name string `json:"name,omitempty"`

		Path string `json:"path,omitempty"`
	}
	bodydata := &body{

		Name: p.Name,

		Path: p.Path,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "post",
		Url:         "/service/platform/assets/v1.0/folders",
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type UpdateFolderXQuery struct {
	FolderId string `json:"folderId,omitempty"`
	IsActive bool   `json:"isActive,omitempty"`
}

/*
summary: Update folder details

description: Update folder details. Eg: Soft delete it
by making `isActive` as `false`.
We currently do not support updating folder name or path.


params: UpdateFolderXQuery
*/
func (c *Assets) UpdateFolder(
	p UpdateFolderXQuery,
) (map[string]interface{}, error) {

	type body struct {
		IsActive bool `json:"isActive,omitempty"`
	}
	bodydata := &body{

		IsActive: p.IsActive,
	}

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "patch",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/folders/%s", p.FolderId),
		Query:       queryParams,
		Body:        bodydata,
		ContentType: "application/json",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type DeleteFolderXQuery struct {
	ID string `json:"_id,omitempty"`
}

/*
summary: Delete folder

description: Delete folder and all its children permanently.


params: DeleteFolderXQuery
*/
func (c *Assets) DeleteFolder(
	p DeleteFolderXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "delete",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/folders/%s", p.ID),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetModulesXQuery struct {
}

/*
summary: Get all transformation modules

description: Get all transformation modules.


params: GetModulesXQuery
*/
func (c *Assets) GetModules(
	p GetModulesXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         "/service/platform/assets/v1.0/playground/plugins",
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

type GetModuleXQuery struct {
	Identifier string `json:"identifier,omitempty"`
}

/*
summary: Get Transformation Module by module identifier

description: Get Transformation Module by module identifier


params: GetModuleXQuery
*/
func (c *Assets) GetModule(
	p GetModuleXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/assets/v1.0/playground/plugins/%s", p.Identifier),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}

// Organization holds Organization object properties
type Organization struct {
	config *PixelbinConfig
}

// NewOrganization returns new Organization instance
func NewOrganization(config *PixelbinConfig) *Organization {
	return &Organization{config: config}
}

type GetAppByTokenXQuery struct {
	Token string `json:"token,omitempty"`
}

/*
summary: Get App Details

description: Get App and org details with the API_TOKEN

params: GetAppByTokenXQuery
*/
func (c *Organization) GetAppByToken(
	p GetAppByTokenXQuery,
) (map[string]interface{}, error) {

	queryParams := make(map[string]string)

	apiClient := &APIClient{
		Conf:        c.config,
		Method:      "get",
		Url:         fmt.Sprintf("/service/platform/organization/v1.0/apps/%s", p.Token),
		Query:       queryParams,
		Body:        nil,
		ContentType: "",
	}

	response, err := apiClient.Execute()
	if err != nil {
		return nil, err
	}
	resp := map[string]interface{}{}
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, common.NewFDKError(err.Error())
	}
	return resp, nil

}
