/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type ProtectionDefApprovals struct {
	RequireCodeOwners   bool  `json:"require_code_owners,omitempty"`
	RequireLatestCommit bool  `json:"require_latest_commit,omitempty"`
	RequireMinimumCount int32 `json:"require_minimum_count,omitempty"`
}
