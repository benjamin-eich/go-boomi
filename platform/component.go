package platform

import "encoding/xml"

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

type CertificateComponent struct {
	XMLName xml.Name `xml:"Component"`

	FolderFullPath string `xml:"folderFullPath,attr,omitempty"`
	ComponentId    string `xml:"componentId,attr"`
	Version        int    `xml:"version,attr"`
	Name           string `xml:"name,attr"`
	Type           string `xml:"type,attr"`
	SubType        string `xml:"subType,attr,omitempty"`
	CreatedDate    string `xml:"createdDate,attr"`
	CreatedBy      string `xml:"createdBy,attr"`
	ModifiedDate   string `xml:"modifiedDate,attr"`
	ModifiedBy     string `xml:"modifiedBy,attr"`
	Deleted        bool   `xml:"deleted,attr"`
	CurrentVersion bool   `xml:"currentVersion,attr"`
	FolderName     string `xml:"folderName,attr"`
	FolderId       string `xml:"folderId,attr"`
	BranchName     string `xml:"branchName,attr"`
	BranchId       string `xml:"branchId,attr"`

	Description string `xml:"description"`

	Object CertificateObject `xml:"object"`
}

type CertificateObject struct {
	Certificate Certificate `xml:"CertificateModel"`
}

type Certificate struct {
	MD5Fingerprint     string `xml:"MD5Fingerprint,attr"`
	SHA1Fingerprint    string `xml:"SHA1Fingerprint,attr"`
	SerialNumber       string `xml:"serialNumber,attr"`
	SignatureAlgorithm string `xml:"signatureAlgorithm,attr"`
	Version            string `xml:"version,attr"`

	Type            string              `xml:"Type"`
	IssuedTo        CertificateIdentity `xml:"IssuedTo"`
	Issuer          CertificateIdentity `xml:"Issuer"`
	Validity        Validity            `xml:"Validity"`
	CertificateData string              `xml:"CertificateData"`
}

type CertificateIdentity struct {
	CommonName             string `xml:"commonName,attr"`
	Country                string `xml:"country,attr"`
	FullName               string `xml:"fullName,attr"`
	Organization           string `xml:"organization,attr,omitempty"`
	OrganizationalUnitName string `xml:"organizationalUnit,attr,omitempty"`
}

type Validity struct {
	ExpireDate string `xml:"expireDate,attr"`
	IssueDate  string `xml:"issueDate,attr"`
}
