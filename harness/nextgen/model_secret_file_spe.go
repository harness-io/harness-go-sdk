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

// This has details of Secret File defined in harness
type SecretFileSpe struct {
	ErrorMessageForInvalidYaml string `json:"errorMessageForInvalidYaml,omitempty"`
	Type_                      string `json:"type"`
	// Identifier of the Secret Manager used to manage the secret.
	SecretManagerIdentifier string `json:"secretManagerIdentifier"`
}
