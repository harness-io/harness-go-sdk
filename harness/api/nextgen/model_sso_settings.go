/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type SsoSettings struct {
	Uuid string `json:"uuid"`
	AppId string `json:"appId"`
	CreatedBy *EmbeddedUser `json:"createdBy,omitempty"`
	CreatedAt int64 `json:"createdAt,omitempty"`
	LastUpdatedBy *EmbeddedUser `json:"lastUpdatedBy,omitempty"`
	LastUpdatedAt int64 `json:"lastUpdatedAt"`
	Type_ string `json:"type"`
	DisplayName string `json:"displayName,omitempty"`
	Url string `json:"url,omitempty"`
	NextIteration int64 `json:"nextIteration,omitempty"`
	NextIterations []int64 `json:"nextIterations,omitempty"`
	AccountId string `json:"accountId,omitempty"`
}
