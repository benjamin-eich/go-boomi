package platform

type ComponentMetadata struct {
	ComponentId    string `json:"componentId"`
	Version        int    `json:"version"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	SubType        string `json:"subType"`
	CreatedDate    string `json:"createdDate"` // format: "2026-04-20T14:39:41Z"
	CreatedBy      string `json:"createdBy"`
	ModifiedDate   string `json:"modifiedDate"` // format: "2026-04-20T14:39:41Z"
	ModifiedBy     string `json:"modifiedBy"`
	Deleted        bool   `json:"deleted"`
	CurrentVersion bool   `json:"currentVersion"`
	FolderName     string `json:"folderName"`
	FolderId       string `json:"folderId"`
	BranchName     string `json:"branchName"`
	BranchId       string `json:"branchId"`
}
