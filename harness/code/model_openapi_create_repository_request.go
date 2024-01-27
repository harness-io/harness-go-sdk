/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OpenapiCreateRepositoryRequest struct {
	DefaultBranch string `json:"default_branch,omitempty"`
	Description string `json:"description,omitempty"`
	ForkId int32 `json:"fork_id,omitempty"`
	GitIgnore string `json:"git_ignore,omitempty"`
	IsPublic bool `json:"is_public,omitempty"`
	License string `json:"license,omitempty"`
	ParentRef string `json:"parent_ref,omitempty"`
	Readme bool `json:"readme,omitempty"`
	Uid string `json:"uid,omitempty"`
}
