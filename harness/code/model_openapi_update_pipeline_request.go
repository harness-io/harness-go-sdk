/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OpenapiUpdatePipelineRequest struct {
	ConfigPath string `json:"config_path,omitempty"`
	Description string `json:"description,omitempty"`
	Disabled bool `json:"disabled,omitempty"`
	Uid string `json:"uid,omitempty"`
}
