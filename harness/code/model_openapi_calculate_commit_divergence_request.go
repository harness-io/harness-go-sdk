/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type OpenapiCalculateCommitDivergenceRequest struct {
	MaxCount int32                         `json:"max_count,omitempty"`
	Requests []RepoCommitDivergenceRequest `json:"requests,omitempty"`
}
