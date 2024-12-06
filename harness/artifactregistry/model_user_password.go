/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package artifactregistry

type UserPassword struct {
	SecretIdentifier string `json:"secretIdentifier,omitempty"`
	SecretSpaceId    int32  `json:"secretSpaceId,omitempty"`
	SecretSpacePath  string `json:"secretSpacePath,omitempty"`
	UserName         string `json:"userName"`
}
