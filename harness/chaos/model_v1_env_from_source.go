/*
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type V1EnvFromSource struct {
	// The ConfigMap to select from +optional
	ConfigMapRef *AllOfv1EnvFromSourceConfigMapRef `json:"configMapRef,omitempty"`
	// An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER. +optional
	Prefix string `json:"prefix,omitempty"`
	// The Secret to select from +optional
	SecretRef *AllOfv1EnvFromSourceSecretRef `json:"secretRef,omitempty"`
}
