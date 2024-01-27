/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TypesPullReqReviewer struct {
	AddedBy *TypesPrincipalInfo `json:"added_by,omitempty"`
	Created int32 `json:"created,omitempty"`
	LatestReviewId int32 `json:"latest_review_id,omitempty"`
	ReviewDecision *EnumPullReqReviewDecision `json:"review_decision,omitempty"`
	Reviewer *TypesPrincipalInfo `json:"reviewer,omitempty"`
	Sha string `json:"sha,omitempty"`
	Type_ *EnumPullReqReviewerType `json:"type,omitempty"`
	Updated int32 `json:"updated,omitempty"`
}
