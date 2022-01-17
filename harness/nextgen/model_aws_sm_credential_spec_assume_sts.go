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

// Returns credentials like RoleArn and externalId and STS role duration specific to IAM role for the AWS Secret Manager.
type AwsSmCredentialSpecAssumeSts struct {
	RoleArn               string `json:"roleArn"`
	ExternalId            string `json:"externalId,omitempty"`
	AssumeStsRoleDuration int32  `json:"assumeStsRoleDuration,omitempty"`
}
