/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

// Harness Manifests
type DockerManifests struct {
	ImageName       string                  `json:"imageName"`
	IsLatestVersion bool                    `json:"isLatestVersion,omitempty"`
	Manifests       []DockerManifestDetails `json:"manifests,omitempty"`
	Version         string                  `json:"version"`
}
