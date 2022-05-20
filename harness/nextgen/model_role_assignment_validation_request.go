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

type RoleAssignmentValidationRequest struct {
	RoleAssignment *RoleAssignment `json:"roleAssignment"`
	// Set it to true if the principal needs to be validated
	ValidatePrincipal bool `json:"validatePrincipal,omitempty"`
	// Set it to true if the role needs to be validated
	ValidateRole bool `json:"validateRole,omitempty"`
	// Set it to true if the resource group needs to be validated
	ValidateResourceGroup bool `json:"validateResourceGroup,omitempty"`
}
