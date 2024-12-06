/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

// Harness Artifact Summary
type ArtifactSummary struct {
	CreatedAt      string       `json:"createdAt,omitempty"`
	DownloadsCount int64        `json:"downloadsCount,omitempty"`
	ImageName      string       `json:"imageName"`
	Labels         []string     `json:"labels,omitempty"`
	ModifiedAt     string       `json:"modifiedAt,omitempty"`
	PackageType    *PackageType `json:"packageType"`
}
