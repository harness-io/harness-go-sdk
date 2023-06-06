/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the Pipeline Service. This is under active development. Beware of the breaking change with respect to the generated code stub.
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Pipeline request body object
type PipelineCreateRequestBody struct {
	// Pipeline YAML (to be passed as a String).
	PipelineYaml string `json:"pipeline_yaml"`
	// Pipeline identifier
	Identifier string `json:"identifier"`
	// Pipeline name
	Name string `json:"name"`
	// Pipeline description
	Description string `json:"description,omitempty"`
	// Pipeline tags
	Tags       map[string]string `json:"tags,omitempty"`
	GitDetails *GitCreateDetails `json:"git_details,omitempty"`
}
