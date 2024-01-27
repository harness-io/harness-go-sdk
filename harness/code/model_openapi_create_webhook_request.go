/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OpenapiCreateWebhookRequest struct {
	Description string `json:"description,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Insecure bool `json:"insecure,omitempty"`
	Secret string `json:"secret,omitempty"`
	Triggers []EnumWebhookTrigger `json:"triggers,omitempty"`
	Url string `json:"url,omitempty"`
}
