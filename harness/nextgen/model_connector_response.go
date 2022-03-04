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

// This has the Connector details along with its metadata.
type ConnectorResponse struct {
	Connector *ConnectorInfo `json:"connector,omitempty"`
	// This is the time at which the Connector was created.
	CreatedAt int64 `json:"createdAt,omitempty"`
	// This is the time at which the Connector was last modified.
	LastModifiedAt  int64                         `json:"lastModifiedAt,omitempty"`
	Status          *ConnectorConnectivityDetails `json:"status,omitempty"`
	ActivityDetails *ConnectorActivityDetails     `json:"activityDetails,omitempty"`
	// This indicates if this Connector is managed by Harness or not. If True, Harness can manage and modify this Connector.
	HarnessManaged        bool              `json:"harnessManaged,omitempty"`
	GitDetails            *EntityGitDetails `json:"gitDetails,omitempty"`
	EntityValidityDetails *EntityGitDetails `json:"entityValidityDetails,omitempty"`
}
