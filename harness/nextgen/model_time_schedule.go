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

// Specifies the occurrence schedule. Occurrence schedule can either be specified as period or as days
type TimeSchedule struct {
	Period *TimeSchedulePeriod `json:"period,omitempty"`
	Days   *TimeScheduleDays   `json:"days,omitempty"`
}
