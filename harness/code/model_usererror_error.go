/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type UsererrorError struct {
	Message string `json:"message,omitempty"`
	// manually changed object to interface
	Values map[string]interface{} `json:"values,omitempty"`
}
