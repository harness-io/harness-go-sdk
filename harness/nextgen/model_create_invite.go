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

// This is the view of the CreateInvite entity defined in Harness
type CreateInvite struct {
	Users        []string      `json:"users,omitempty"`
	RoleBindings []RoleBinding `json:"roleBindings,omitempty"`
	InviteType   string        `json:"inviteType,omitempty"`
	UserGroups   []string      `json:"userGroups,omitempty"`
}
