/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

// Artifact Metadata
type RegistryArtifactMetadata struct {
	DownloadsCount     int64        `json:"downloadsCount,omitempty"`
	Labels             []string     `json:"labels,omitempty"`
	LastModified       string       `json:"lastModified,omitempty"`
	LatestVersion      string       `json:"latestVersion"`
	Name               string       `json:"name"`
	PackageType        *PackageType `json:"packageType,omitempty"`
	RegistryIdentifier string       `json:"registryIdentifier"`
	RegistryPath       string       `json:"registryPath"`
}
