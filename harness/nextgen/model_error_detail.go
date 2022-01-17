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

// List of errors and their details.
type ErrorDetail struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int32  `json:"code,omitempty"`
}
