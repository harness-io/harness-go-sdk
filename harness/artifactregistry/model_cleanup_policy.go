/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

// Cleanup Policy for Harness Artifact Registries
type CleanupPolicy struct {
	ExpireDays    int32    `json:"expireDays,omitempty"`
	Name          string   `json:"name,omitempty"`
	PackagePrefix []string `json:"packagePrefix,omitempty"`
	VersionPrefix []string `json:"versionPrefix,omitempty"`
}
