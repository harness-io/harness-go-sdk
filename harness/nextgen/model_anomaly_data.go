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

// This object contains details of a cost anomaly
type AnomalyData struct {
	Id                       string      `json:"id,omitempty"`
	Time                     int64       `json:"time,omitempty"`
	AnomalyRelativeTime      string      `json:"anomalyRelativeTime,omitempty"`
	ActualAmount             float64     `json:"actualAmount,omitempty"`
	ExpectedAmount           float64     `json:"expectedAmount,omitempty"`
	AnomalousSpend           float64     `json:"anomalousSpend,omitempty"`
	AnomalousSpendPercentage float64     `json:"anomalousSpendPercentage,omitempty"`
	ResourceName             string      `json:"resourceName,omitempty"`
	ResourceInfo             string      `json:"resourceInfo,omitempty"`
	Entity                   *EntityInfo `json:"entity,omitempty"`
	Details                  string      `json:"details,omitempty"`
	Status                   string      `json:"status,omitempty"`
	StatusRelativeTime       string      `json:"statusRelativeTime,omitempty"`
	Comment                  string      `json:"comment,omitempty"`
	CloudProvider            string      `json:"cloudProvider,omitempty"`
	AnomalyScore             float64     `json:"anomalyScore,omitempty"`
	UserFeedback             string      `json:"userFeedback,omitempty"`
}
