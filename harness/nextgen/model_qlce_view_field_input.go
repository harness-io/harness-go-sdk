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

type QlceViewFieldInput struct {
	FieldId   string `json:"fieldId,omitempty"`
	FieldName string `json:"fieldName,omitempty"`
	// Perspective filter Category, CLUSTER means Kubernetes
	Identifier     string `json:"identifier,omitempty"`
	IdentifierName string `json:"identifierName,omitempty"`
}
