/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running +optional
type AllOfv1VolumeFlocker struct {
	// Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated +optional
	DatasetName string `json:"datasetName,omitempty"`
	// UUID of the dataset. This is unique identifier of a Flocker dataset +optional
	DatasetUUID string `json:"datasetUUID,omitempty"`
}
