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

// This contains Validity Details of the Entity
type EntityGitDetails struct {
	// Indicates if the Entity is valid
	Valid bool `json:"valid,omitempty"`
	// This has the Git File content if the entity is invalid
	InvalidYaml string `json:"invalidYaml,omitempty"`
}
