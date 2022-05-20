/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This contains details of Git Sync Config
type GitSyncConfig struct {
	// Git Sync Config Id.
	Identifier string `json:"identifier,omitempty"`
	// Name of the repository. Any leading/trailing spaces will be removed.
	Name string `json:"name,omitempty"`
	// Project Identifier for the Entity.
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Organization Identifier for the Entity.
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Id of the Connector referenced in Git
	GitConnectorRef string `json:"gitConnectorRef,omitempty"`
	// URL of the repository. Any leading/trailing spaces will be removed.
	Repo string `json:"repo,omitempty"`
	// Name of the branch. Any leading/trailing spaces will be removed.
	Branch string `json:"branch,omitempty"`
	// Connector Type
	GitConnectorType string `json:"gitConnectorType"`
	// List of all Root Folder Details
	GitSyncFolderConfigDTOs []GitSyncFolderConfig `json:"gitSyncFolderConfigDTOs,omitempty"`
}
