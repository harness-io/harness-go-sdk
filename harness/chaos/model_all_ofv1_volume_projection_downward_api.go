/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// information about the downwardAPI data to project +optional
type AllOfv1VolumeProjectionDownwardApi struct {
	// Items is a list of DownwardAPIVolume file +optional
	Items []V1DownwardApiVolumeFile `json:"items,omitempty"`
}
