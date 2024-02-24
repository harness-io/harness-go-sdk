/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type EnumPullReqActivityType string

// List of EnumPullReqActivityType
const (
	BRANCH_DELETE_EnumPullReqActivityType EnumPullReqActivityType = "branch-delete"
	BRANCH_UPDATE_EnumPullReqActivityType EnumPullReqActivityType = "branch-update"
	CODE_COMMENT_EnumPullReqActivityType  EnumPullReqActivityType = "code-comment"
	COMMENT_EnumPullReqActivityType       EnumPullReqActivityType = "comment"
	MERGE_EnumPullReqActivityType         EnumPullReqActivityType = "merge"
	REVIEW_SUBMIT_EnumPullReqActivityType EnumPullReqActivityType = "review-submit"
	STATE_CHANGE_EnumPullReqActivityType  EnumPullReqActivityType = "state-change"
	TITLE_CHANGE_EnumPullReqActivityType  EnumPullReqActivityType = "title-change"
)
