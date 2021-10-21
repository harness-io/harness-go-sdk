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

type AwsKmsConfigDto struct {
	Name string `json:"name,omitempty"`
	AccountIdentifier string `json:"accountIdentifier,omitempty"`
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	Tags map[string]string `json:"tags,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	Description string `json:"description,omitempty"`
	EncryptionType string `json:"encryptionType,omitempty"`
	HarnessManaged bool `json:"harnessManaged,omitempty"`
	Default_ bool `json:"default,omitempty"`
	BaseAwsKmsConfigDTO *BaseAwsKmsConfigDto `json:"baseAwsKmsConfigDTO,omitempty"`
}
