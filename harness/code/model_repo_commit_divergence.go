/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RepoCommitDivergence struct {
	Ahead int32 `json:"ahead,omitempty"`
	Behind int32 `json:"behind,omitempty"`
}
