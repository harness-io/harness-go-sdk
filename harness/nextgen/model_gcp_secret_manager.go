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

// This contains details of GCP Secret Manager
type GcpSecretManager struct {
	// Reference to the secret containing credentials of IAM service account for Google Secret Manager
	CredentialsRef string `json:"credentialsRef,omitempty"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors []string `json:"delegateSelectors,omitempty"`
	// Should the secret manager execute operations on the delegate, or via Harness platform
	ExecuteOnDelegate bool `json:"executeOnDelegate"`
	// Boolean value to indicate that Credentials are taken from the Delegate.
	AssumeCredentialsOnDelegate            bool                                    `json:"assumeCredentialsOnDelegate"`
	Credential                             *GcpConnectorCredential                 `json:"credential,omitempty"`
	GcpOidcTokenExchangeDetailsForDelegate *GcpOidcTokenExchangeDetailsForDelegate `json:"gcpOidcTokenExchangeDetailsForDelegate,omitempty"`
	IgnoreTestConnection                   bool                                    `json:"ignoreTestConnection,omitempty"`
	Default_                               bool                                    `json:"default"`
	ConnectorType                          string                                  `json:"connectorType"`
}
