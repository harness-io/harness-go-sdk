/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package chaos

type V1PersistentVolumeClaimSpec struct {
	// AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1 +optional
	AccessModes []V1PersistentVolumeAccessMode `json:"accessModes,omitempty"`
	// This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) * An existing custom resource that implements data population (Alpha) In order to use custom resource types that implement data population, the AnyVolumeDataSource feature gate must be enabled. If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. +optional
	DataSource *AllOfv1PersistentVolumeClaimSpecDataSource `json:"dataSource,omitempty"`
	// Resources represents the minimum resources the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources +optional
	Resources *AllOfv1PersistentVolumeClaimSpecResources `json:"resources,omitempty"`
	// A label query over volumes to consider for binding. +optional
	Selector *AllOfv1PersistentVolumeClaimSpecSelector `json:"selector,omitempty"`
	// Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1 +optional
	StorageClassName string `json:"storageClassName,omitempty"`
	// volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec. +optional
	VolumeMode *AllOfv1PersistentVolumeClaimSpecVolumeMode `json:"volumeMode,omitempty"`
	// VolumeName is the binding reference to the PersistentVolume backing this claim. +optional
	VolumeName string `json:"volumeName,omitempty"`
}
