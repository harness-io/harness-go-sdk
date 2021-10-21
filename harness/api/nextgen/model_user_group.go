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

type UserGroup struct {
	Id string `json:"id,omitempty"`
	AccountIdentifier string `json:"accountIdentifier"`
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	Identifier string `json:"identifier,omitempty"`
	IsSsoLinked bool `json:"isSsoLinked,omitempty"`
	LinkedSsoType string `json:"linkedSsoType,omitempty"`
	LinkedSsoId string `json:"linkedSsoId,omitempty"`
	LinkedSsoDisplayName string `json:"linkedSsoDisplayName,omitempty"`
	SsoGroupId string `json:"ssoGroupId,omitempty"`
	SsoGroupName string `json:"ssoGroupName,omitempty"`
	Name string `json:"name,omitempty"`
	Users []string `json:"users"`
	NotificationConfigs []NotificationSettingConfig `json:"notificationConfigs"`
	HarnessManaged bool `json:"harnessManaged,omitempty"`
	Description string `json:"description"`
	Tags []NgTag `json:"tags"`
	CreatedAt int64 `json:"createdAt,omitempty"`
	LastModifiedAt int64 `json:"lastModifiedAt,omitempty"`
	Version int64 `json:"version,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
}
