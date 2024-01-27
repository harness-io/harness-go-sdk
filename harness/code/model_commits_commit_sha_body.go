/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CommitsCommitShaBody struct {
	CheckUid string `json:"check_uid,omitempty"`
	Link string `json:"link,omitempty"`
	Payload *TypesCheckPayload `json:"payload,omitempty"`
	Status *EnumCheckStatus `json:"status,omitempty"`
	Summary string `json:"summary,omitempty"`
}
