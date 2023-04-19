/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This contains details of the Github API access credentials Specs such as references of private key
type GithubAppSpec struct {
	InstallationId    string `json:"installationId,omitempty"`
	ApplicationId     string `json:"applicationId,omitempty"`
	InstallationIdRef string `json:"installationIdRef,omitempty"`
	ApplicationIdRef  string `json:"applicationIdRef,omitempty"`
	PrivateKeyRef     string `json:"privateKeyRef"`
}
