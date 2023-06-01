/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type GithubStore struct {
	ConnectorRef string `json:"connectorRef"`
	GitFetchType string `json:"gitFetchType"`
	Branch string `json:"branch,omitempty"`
	CommitId string `json:"commitId,omitempty"`
	Paths []string `json:"paths,omitempty"`
	FolderPath string `json:"folderPath,omitempty"`
	RepoName string `json:"repoName,omitempty"`
}
