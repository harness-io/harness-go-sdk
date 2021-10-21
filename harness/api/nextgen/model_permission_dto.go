/*
 * Access Control API Reference
 *
 * This is the Open Api Spec 3 for the Access Control Service. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type PermissionDto struct {
	Identifier string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	IncludeInAllRoles bool `json:"includeInAllRoles,omitempty"`
	AllowedScopeLevels []string `json:"allowedScopeLevels,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Action string `json:"action,omitempty"`
}
