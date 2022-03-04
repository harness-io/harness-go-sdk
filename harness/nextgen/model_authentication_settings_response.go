/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This contains information on the Authentication Settings defined in Harness.
type AuthenticationSettingsResponse struct {
	// List of Auth Settings configured for an Account.
	NgAuthSettings []NgAuthSettings `json:"ngAuthSettings,omitempty"`
	// List of the whitelisted domains.
	WhitelistedDomains []string `json:"whitelistedDomains,omitempty"`
	// Indicates if the Authentication Mechanism is SSO or NON-SSO.
	AuthenticationMechanism string `json:"authenticationMechanism,omitempty"`
	// If Two Factor Authentication is enabled, this value is true. Otherwise, it is false.
	TwoFactorEnabled bool `json:"twoFactorEnabled,omitempty"`
}
