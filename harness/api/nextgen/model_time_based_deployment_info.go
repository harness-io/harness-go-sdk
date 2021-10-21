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

type TimeBasedDeploymentInfo struct {
	EpochTime int64 `json:"epochTime,omitempty"`
	TotalCount int64 `json:"totalCount,omitempty"`
	SuccessCount int64 `json:"successCount,omitempty"`
	FailedCount int64 `json:"failedCount,omitempty"`
}
