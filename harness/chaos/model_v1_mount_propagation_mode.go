/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1MountPropagationMode string

// List of v1.MountPropagationMode
const (
	NONE_V1MountPropagationMode V1MountPropagationMode = "None"
	HOST_TO_CONTAINER_V1MountPropagationMode V1MountPropagationMode = "HostToContainer"
	BIDIRECTIONAL_V1MountPropagationMode V1MountPropagationMode = "Bidirectional"
)
