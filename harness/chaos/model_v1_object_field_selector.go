/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1ObjectFieldSelector struct {
	// Version of the schema the FieldPath is written in terms of, defaults to \"v1\". +optional
	ApiVersion string `json:"apiVersion,omitempty"`
	// Path of the field to select in the specified API version.
	FieldPath string `json:"fieldPath,omitempty"`
}
