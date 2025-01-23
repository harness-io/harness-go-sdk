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

// This has configuration details for the AWS KMS Secret Manager.
type AwsKmsConnector struct {
	Credential *AwsKmsConnectorCredential `json:"credential"`
	// ARN for AWS KMS.
	KmsArn string `json:"kmsArn"`
	// Region for AWS KMS.
	Region    string `json:"region"`
	IsDefault bool   `json:"isDefault,omitempty"`
	// List of Delegate Selectors that belong to the same Delegate and are used to connect to the Secret Manager.
	DelegateSelectors                      []string                                `json:"delegateSelectors,omitempty"`
	AwsOidcTokenExchangeDetailsForDelegate *AwsOidcTokenExchangeDetailsForDelegate `json:"awsOidcTokenExchangeDetailsForDelegate,omitempty"`
	IgnoreTestConnection                   bool                                    `json:"ignoreTestConnection,omitempty"`
	// Should the secret manager execute operations on the delegate, or via Harness platform
	ExecuteOnDelegate bool   `json:"executeOnDelegate,omitempty"`
	Default_          bool   `json:"default,omitempty"`
	ConnectorType     string `json:"connectorType"`
}
