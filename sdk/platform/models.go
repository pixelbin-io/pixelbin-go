package platform

import "os"

// folderItem used by Assets
type folderItem struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

// exploreItem used by Assets
type exploreItem struct {
	ID     string  `json:"_id"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
	Path   string  `json:"path"`
	FileId string  `json:"fileId"`
	Format string  `json:"format"`
	Size   float64 `json:"size"`
	Access string  `json:"access"`
}

// page used by Assets
type page struct {
	Type      string  `json:"type"`
	Size      float64 `json:"size"`
	Current   float64 `json:"current"`
	HasNext   bool    `json:"hasNext"`
	ItemTotal float64 `json:"itemTotal"`
}

// exploreResponse used by Assets
type exploreResponse struct {
	Items []exploreItem `json:"items"`
	Page  page          `json:"page"`
}

// ListFilesResponse used by Assets
type ListFilesResponse struct {
	Items []exploreItem `json:"items"`
	Page  page          `json:"page"`
}

// exploreFolderResponse used by Assets
type exploreFolderResponse struct {
	Folder folderItem    `json:"folder"`
	Items  []exploreItem `json:"items"`
	Page   page          `json:"page"`
}

// FileUploadRequest used by Assets
type FileUploadRequest struct {
	File             *os.File               `json:"file"`
	Path             string                 `json:"path"`
	Name             string                 `json:"name"`
	Access           AccessEnum             `json:"access"`
	Tags             []string               `json:"tags"`
	Metadata         map[string]interface{} `json:"metadata"`
	Overwrite        bool                   `json:"overwrite"`
	FilenameOverride bool                   `json:"filenameOverride"`
}

// UrlUploadRequest used by Assets
type UrlUploadRequest struct {
	URL              string                 `json:"url"`
	Path             string                 `json:"path"`
	Name             string                 `json:"name"`
	Access           AccessEnum             `json:"access"`
	Tags             []string               `json:"tags"`
	Metadata         map[string]interface{} `json:"metadata"`
	Overwrite        bool                   `json:"overwrite"`
	FilenameOverride bool                   `json:"filenameOverride"`
}

// UploadResponse used by Assets
type UploadResponse struct {
	ID        string                 `json:"_id"`
	FileId    string                 `json:"fileId"`
	Name      string                 `json:"name"`
	Path      string                 `json:"path"`
	Format    string                 `json:"format"`
	Size      float64                `json:"size"`
	Access    string                 `json:"access"`
	Tags      []string               `json:"tags"`
	Metadata  map[string]interface{} `json:"metadata"`
	URL       string                 `json:"url"`
	Thumbnail string                 `json:"thumbnail"`
}

// SignedUploadRequest used by Assets
type SignedUploadRequest struct {
	Name             string                 `json:"name"`
	Path             string                 `json:"path"`
	Format           string                 `json:"format"`
	Access           AccessEnum             `json:"access"`
	Tags             []string               `json:"tags"`
	Metadata         map[string]interface{} `json:"metadata"`
	Overwrite        bool                   `json:"overwrite"`
	FilenameOverride bool                   `json:"filenameOverride"`
}

// SignedUploadResponse used by Assets
type SignedUploadResponse struct {
	S3PresignedUrl PresignedUrl `json:"s3PresignedUrl"`
}

// PresignedUrl used by Assets
type PresignedUrl struct {
	URL    string                 `json:"url"`
	Fields map[string]interface{} `json:"fields"`
}

// FilesResponse used by Assets
type FilesResponse struct {
	ID        string                 `json:"_id"`
	Name      string                 `json:"name"`
	Path      string                 `json:"path"`
	FileId    string                 `json:"fileId"`
	Format    string                 `json:"format"`
	Size      float64                `json:"size"`
	Access    string                 `json:"access"`
	IsActive  bool                   `json:"isActive"`
	Tags      []string               `json:"tags"`
	Metadata  map[string]interface{} `json:"metadata"`
	URL       string                 `json:"url"`
	Thumbnail string                 `json:"thumbnail"`
}

// UpdateFileRequest used by Assets
type UpdateFileRequest struct {
	Name     string                 `json:"name"`
	Path     string                 `json:"path"`
	Access   string                 `json:"access"`
	IsActive bool                   `json:"isActive"`
	Tags     []string               `json:"tags"`
	Metadata map[string]interface{} `json:"metadata"`
}

// FoldersResponse used by Assets
type FoldersResponse struct {
	ID       string `json:"_id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	IsActive bool   `json:"isActive"`
}

// CreateFolderRequest used by Assets
type CreateFolderRequest struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// UpdateFolderRequest used by Assets
type UpdateFolderRequest struct {
	IsActive bool `json:"isActive"`
}

// TransformationModulesResponse used by Assets
type TransformationModulesResponse struct {
	Delimiters Delimiter                               `json:"delimiters"`
	Plugins    map[string]TransformationModuleResponse `json:"plugins"`
	Presets    []interface{}                           `json:"presets"`
}

// DeleteMultipleFilesRequest used by Assets
type DeleteMultipleFilesRequest struct {
	Ids []string `json:"ids"`
}

// Delimiter used by Assets
type Delimiter struct {
	OperationSeparator string `json:"operationSeparator"`
	ParameterSeparator string `json:"parameterSeparator"`
}

// TransformationModuleResponse used by Assets
type TransformationModuleResponse struct {
	Identifier  string                 `json:"identifier"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Credentials map[string]interface{} `json:"credentials"`
	Operations  []interface{}          `json:"operations"`
	Enabled     bool                   `json:"enabled"`
}

// OrganizationDetailSchema used by Organization
type OrganizationDetailSchema struct {
	ID        float64 `json:"_id"`
	Name      string  `json:"name"`
	CloudName string  `json:"cloudName"`
	OwnerId   string  `json:"ownerId"`
	Active    bool    `json:"active"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

// AppSchema used by Organization
type AppSchema struct {
	ID          float64  `json:"_id"`
	OrgId       float64  `json:"orgId"`
	Name        string   `json:"name"`
	Token       string   `json:"token"`
	Permissions []string `json:"permissions"`
	Active      bool     `json:"active"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}

// AppDetailsByToken used by Organization
type AppDetailsByToken struct {
	App AppSchema                `json:"app"`
	Org OrganizationDetailSchema `json:"org"`
}

// ErrorSchema used by Organization
type ErrorSchema struct {
	Message string `json:"message"`
}
