/*
 * Harness Artifact Registry API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package har

// RegistryType : refers to type of registry i.e virtual or upstream
type RegistryType string

// List of RegistryType
const (
	VIRTUAL_RegistryType  RegistryType = "VIRTUAL"
	UPSTREAM_RegistryType RegistryType = "UPSTREAM"
)
