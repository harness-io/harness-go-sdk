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

// Details of an independent resource that can be controlled by a schedule
type StaticScheduleResource struct {
	// ID of the resource to be controlled by schedule. For AutoStopping rule, this would be the ID of the rule
	Id string `json:"id,omitempty"`
	// Type of the resource to be controlled
	Type_ string `json:"type,omitempty"`
}
