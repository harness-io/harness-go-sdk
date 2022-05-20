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

// This is the SSH key authentication details defined in Harness.
type SshKeySpec struct {
	ErrorMessageForInvalidYaml string `json:"errorMessageForInvalidYaml,omitempty"`
	Type_                      string `json:"type"`
	// SSH port
	Port int32    `json:"port,omitempty"`
	Auth *SshAuth `json:"auth"`
}
