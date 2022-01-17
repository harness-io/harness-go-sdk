/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Contains information of Git Sync Config
type GitSyncConfig struct {
	// Git Sync Config Id
	Identifier string `json:"identifier,omitempty"`
	// Repo Name
	Name string `json:"name,omitempty"`
	// Project Identifier for the Entity
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Organization Identifier for the Entity
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Referenced Connector Identifier
	GitConnectorRef string `json:"gitConnectorRef,omitempty"`
	// Repo Url
	Repo string `json:"repo,omitempty"`
	// Branch Name
	Branch string `json:"branch,omitempty"`
	// Connector Type
	GitConnectorType string `json:"gitConnectorType"`
	// List of all Root Folder Details
	GitSyncFolderConfigDTOs []GitSyncFolderConfig `json:"gitSyncFolderConfigDTOs,omitempty"`
}
