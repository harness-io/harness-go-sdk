/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TypesMembershipSpace struct {
	AddedBy *TypesPrincipalInfo `json:"added_by,omitempty"`
	Created int32 `json:"created,omitempty"`
	Role *EnumMembershipRole `json:"role,omitempty"`
	Space *TypesSpace `json:"space,omitempty"`
	Updated int32 `json:"updated,omitempty"`
}
