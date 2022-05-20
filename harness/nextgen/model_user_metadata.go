/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This is the view of the UserMetadata entity defined in Harness
type UserMetadata struct {
	Name              string `json:"name,omitempty"`
	Email             string `json:"email,omitempty"`
	Uuid              string `json:"uuid,omitempty"`
	Locked            bool   `json:"locked,omitempty"`
	Disabled          bool   `json:"disabled,omitempty"`
	ExternallyManaged bool   `json:"externallyManaged,omitempty"`
}
