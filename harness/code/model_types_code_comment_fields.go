/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type TypesCodeCommentFields struct {
	LineNew int32 `json:"line_new,omitempty"`
	LineOld int32 `json:"line_old,omitempty"`
	MergeBaseSha string `json:"merge_base_sha,omitempty"`
	Outdated bool `json:"outdated,omitempty"`
	Path string `json:"path,omitempty"`
	SourceSha string `json:"source_sha,omitempty"`
	SpanNew int32 `json:"span_new,omitempty"`
	SpanOld int32 `json:"span_old,omitempty"`
}
