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

// This contains a list of Entities corresponding to a specific Git Sync Config Id
type GitSyncRepoFiles struct {
	// Git Sync Config Id
	GitSyncConfigIdentifier string `json:"gitSyncConfigIdentifier,omitempty"`
	// List of all Git Sync Entities based on their Type specific to the given Repo
	GitSyncEntityLists []GitSyncEntityList `json:"gitSyncEntityLists,omitempty"`
}
