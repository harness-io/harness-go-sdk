/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type SystemConfigOutput struct {
	PublicResourceCreationEnabled bool `json:"public_resource_creation_enabled,omitempty"`
	UserSignupAllowed bool `json:"user_signup_allowed,omitempty"`
}
