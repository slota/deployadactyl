// Package structs contains structs that are reused in multiple locations.
package structs

// DeploymentInfo is a collection of properties necessary for a deployment.
type DeploymentInfo struct {
	ArtifactURL string `json:"artifact_url"`
	Manifest    string `json:"manifest"`
	Username    string
	Password    string
	Environment string
	Org         string
	Space       string
	AppName     string
	UUID        string
	SkipSSL     bool

	// Generic map used for users to provide their own deployment properties in JSON format.
	Data map[string]interface{} `json:"data"`
}
