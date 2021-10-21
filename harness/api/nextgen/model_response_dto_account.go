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

type ResponseDtoAccount struct {
	Status string `json:"status,omitempty"`
	Data *Account `json:"data,omitempty"`
	MetaData *interface{} `json:"metaData,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}
