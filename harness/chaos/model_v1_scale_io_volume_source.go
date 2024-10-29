/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1ScaleIoVolumeSource struct {
	// Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. \"ext4\", \"xfs\", \"ntfs\". Default is \"xfs\". +optional
	FsType string `json:"fsType,omitempty"`
	// The host address of the ScaleIO API Gateway.
	Gateway string `json:"gateway,omitempty"`
	// The name of the ScaleIO Protection Domain for the configured storage. +optional
	ProtectionDomain string `json:"protectionDomain,omitempty"`
	// Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
	SecretRef *AllOfv1ScaleIoVolumeSourceSecretRef `json:"secretRef,omitempty"`
	// Flag to enable/disable SSL communication with Gateway, default false +optional
	SslEnabled bool `json:"sslEnabled,omitempty"`
	// Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned. +optional
	StorageMode string `json:"storageMode,omitempty"`
	// The ScaleIO Storage Pool associated with the protection domain. +optional
	StoragePool string `json:"storagePool,omitempty"`
	// The name of the storage system as configured in ScaleIO.
	System string `json:"system,omitempty"`
	// The name of a volume already created in the ScaleIO system that is associated with this volume source.
	VolumeName string `json:"volumeName,omitempty"`
}
