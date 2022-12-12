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

type AppDMetricDefinitions struct {
	Identifier string `json:"identifier"`
	MetricName string `json:"metricName"`
	RiskProfile *RiskProfile `json:"riskProfile,omitempty"`
	Analysis *AnalysisDto `json:"analysis,omitempty"`
	Sli *Slidto `json:"sli,omitempty"`
	GroupName string `json:"groupName,omitempty"`
	BaseFolder string `json:"baseFolder,omitempty"`
	MetricPath string `json:"metricPath,omitempty"`
	CompleteMetricPath string `json:"completeMetricPath,omitempty"`
	CompleteServiceInstanceMetricPath string `json:"completeServiceInstanceMetricPath,omitempty"`
}
