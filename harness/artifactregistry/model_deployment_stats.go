/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

// DeploymentStats
type DeploymentStats struct {
	PreProduction int32 `json:"PreProduction"`
	Production    int32 `json:"Production"`
}
