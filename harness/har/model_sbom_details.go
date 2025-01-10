/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

type SbomDetails struct {
	AllowListViolations int32  `json:"allowListViolations,omitempty"`
	ArtifactId          string `json:"artifactId,omitempty"`
	ArtifactSourceId    string `json:"artifactSourceId,omitempty"`
	AvgScore            string `json:"avgScore,omitempty"`
	ComponentsCount     int32  `json:"componentsCount,omitempty"`
	DenyListViolations  int32  `json:"denyListViolations,omitempty"`
	MaxScore            string `json:"maxScore,omitempty"`
	OrchestrationId     string `json:"orchestrationId,omitempty"`
	OrgId               string `json:"orgId,omitempty"`
	ProjectId           string `json:"projectId,omitempty"`
}
