/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OpenapiCreateSecretRequest struct {
	Data string `json:"data,omitempty"`
	Description string `json:"description,omitempty"`
	SpaceRef string `json:"space_ref,omitempty"`
	Uid string `json:"uid,omitempty"`
}
