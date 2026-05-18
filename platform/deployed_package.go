package platform

type DeployedPackage struct {
	DeploymentId     string `json:"deploymentId"`
	Version          int    `json:"version"`
	PackageId        string `json:"packageId"`
	PackageVersion   int    `json:"packageVersion"`
	EnvironmentId    string `json:"environmentId"`
	ComponentId      string `json:"componentId"`
	ComponentVersion int    `json:"componentVersion"`
	ComponentType    string `json:"componentType"`
	DeployedDate     string `json:"deployedDate"`
	DeployedBy       string `json:"deployedBy"`
	Notes            string `json:"notes"`
	Active           bool   `json:"active"`
	ListenerStatus   string `json:"listenerStatus"`
	BranchName       string `json:"branchName"`
	Message          string `json:"message"`
}
