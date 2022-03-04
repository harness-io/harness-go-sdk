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

type ServiceUsageRecord struct {
	ServiceId    float64 `json:"service_id,omitempty"`
	Path         string  `json:"path,omitempty"`
	SessionId    string  `json:"session_id,omitempty"`
	IdleTimeMins float64 `json:"idle_time_mins,omitempty"`
	CreatedAt    string  `json:"created_at,omitempty"`
}
