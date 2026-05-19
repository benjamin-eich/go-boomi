package platform

type Component struct {
	ComponentId                string         `json:"componentId"`
	Version                    int            `json:"version"`
	Name                       string         `json:"name"`
	Type                       string         `json:"type"`
	SubType                    string         `json:"subType"`
	CreatedDate                string         `json:"createdDate"` // format: "2026-04-20T14:39:41Z"
	CreatedBy                  string         `json:"createdBy"`
	ModifiedDate               string         `json:"modifiedDate"` // format: "2026-04-20T14:39:41Z"
	ModifiedBy                 string         `json:"modifiedBy"`
	Deleted                    bool           `json:"deleted"`
	CurrentVersion             bool           `json:"currentVersion"`
	FolderName                 string         `json:"folderName"`
	FolderFullPath             string         `json:"folderFullPath"`
	FolderId                   string         `json:"folderId"`
	BranchName                 string         `json:"branchName"`
	BranchId                   string         `json:"branchId"`
	CopiedFromComponentId      string         `json:"copiedFromComponentId"`
	CopiedFromComponentVersion int            `json:"copiedFromComponentVersion"`
	EncryptedValues            any            `json:"encryptedValues"`
	Description                string         `json:"description"`
	Object                     map[string]any `json:"object"`
}
