/*
 * Infrastructure as Code Management
 *
 * Services for Harness IaCM Module.
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type IacmUpsertDefaultPipelineRequestBody struct {
	// The default pipeline is associated with this operation
	Operation string `json:"operation"`
	// The default pipeline associated with the provisioner operation
	Pipeline string `json:"pipeline"`
	// The default pipeline is associated with this provisioner
	Provisioner string `json:"provisioner"`
	// Time the default pipeline was last updated
	Updated int64 `json:"updated"`
	// The default pipeline is associated with this workspace if specified
	Workspace string `json:"workspace,omitempty"`
}
