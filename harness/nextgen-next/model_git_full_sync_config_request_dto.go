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

type GitFullSyncConfigRequestDto struct {
	Branch            string `json:"branch"`
	Message           string `json:"message,omitempty"`
	BaseBranch        string `json:"baseBranch,omitempty"`
	CreatePullRequest bool   `json:"createPullRequest,omitempty"`
	RepoIdentifier    string `json:"repoIdentifier"`
}
