/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

// Artifact Metadata
type ArtifactMetadata struct {
	DeploymentMetadata *DeploymentMetadata `json:"deploymentMetadata,omitempty"`
	DownloadsCount     int64               `json:"downloadsCount,omitempty"`
	Labels             []string            `json:"labels,omitempty"`
	LastModified       string              `json:"lastModified,omitempty"`
	LatestVersion      string              `json:"latestVersion"`
	Name               string              `json:"name"`
	PackageType        *PackageType        `json:"packageType,omitempty"`
	PullCommand        string              `json:"pullCommand,omitempty"`
	RegistryIdentifier string              `json:"registryIdentifier"`
	RegistryPath       string              `json:"registryPath"`
	ScannedDigest      []StoDigestMetadata `json:"scannedDigest,omitempty"`
	ScannedDigestCount int64               `json:"scannedDigestCount,omitempty"`
	StoMetadata        *StoMetadata        `json:"stoMetadata,omitempty"`
	TotalDigestCount   int64               `json:"totalDigestCount,omitempty"`
	Version            string              `json:"version"`
}
