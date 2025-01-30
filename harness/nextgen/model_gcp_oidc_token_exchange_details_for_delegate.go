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

type GcpOidcTokenExchangeDetailsForDelegate struct {
	OidcIdToken                             string                          `json:"oidcIdToken,omitempty"`
	OidcAccessTokenStsEndpoint              string                          `json:"oidcAccessTokenStsEndpoint,omitempty"`
	OidcAccessTokenIamSaEndpoint            string                          `json:"oidcAccessTokenIamSaEndpoint,omitempty"`
	GcpServiceAccountEmail                  string                          `json:"gcpServiceAccountEmail,omitempty"`
	OidcWorkloadAccessTokenRequestStructure *OidcWorkloadAccessTokenRequest `json:"oidcWorkloadAccessTokenRequestStructure,omitempty"`
	OidcChartmuseumGcpConfigStructure       *OidcChartmuseumGcpConfig       `json:"oidcChartmuseumGcpConfigStructure,omitempty"`
	IdTokenExpiryTime                       int64                           `json:"idTokenExpiryTime,omitempty"`
}
