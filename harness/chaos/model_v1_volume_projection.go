/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1VolumeProjection struct {
	// information about the configMap data to project +optional
	ConfigMap *AllOfv1VolumeProjectionConfigMap `json:"configMap,omitempty"`
	// information about the downwardAPI data to project +optional
	DownwardAPI *AllOfv1VolumeProjectionDownwardApi `json:"downwardAPI,omitempty"`
	// information about the secret data to project +optional
	Secret *AllOfv1VolumeProjectionSecret `json:"secret,omitempty"`
	// information about the serviceAccountToken data to project +optional
	ServiceAccountToken *AllOfv1VolumeProjectionServiceAccountToken `json:"serviceAccountToken,omitempty"`
}
