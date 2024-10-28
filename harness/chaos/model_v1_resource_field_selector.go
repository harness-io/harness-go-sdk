/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1ResourceFieldSelector struct {
	// Container name: required for volumes, optional for env vars +optional
	ContainerName string `json:"containerName,omitempty"`
	// Specifies the output format of the exposed resources, defaults to \"1\" +optional
	Divisor *AllOfv1ResourceFieldSelectorDivisor `json:"divisor,omitempty"`
	// Required: resource to select
	Resource string `json:"resource,omitempty"`
}
