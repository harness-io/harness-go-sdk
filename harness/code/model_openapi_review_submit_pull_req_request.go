/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type OpenapiReviewSubmitPullReqRequest struct {
	CommitSha string `json:"commit_sha,omitempty"`
	Decision *EnumPullReqReviewDecision `json:"decision,omitempty"`
}
