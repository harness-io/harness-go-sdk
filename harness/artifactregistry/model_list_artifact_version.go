/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

// A list of Artifact versions
type ListArtifactVersion struct {
	// A list of Artifact versions
	ArtifactVersions []ArtifactVersionMetadata `json:"artifactVersions,omitempty"`
	// The total number of items
	ItemCount int64 `json:"itemCount,omitempty"`
	// The total number of pages
	PageCount int64 `json:"pageCount,omitempty"`
	// The current page
	PageIndex int64 `json:"pageIndex,omitempty"`
	// The number of items per page
	PageSize int32 `json:"pageSize,omitempty"`
}
