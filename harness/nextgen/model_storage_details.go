/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type StorageDetails struct {
	Id                      string  `json:"id,omitempty"`
	InstanceId              string  `json:"instanceId,omitempty"`
	InstanceName            string  `json:"instanceName,omitempty"`
	ClaimName               string  `json:"claimName,omitempty"`
	ClaimNamespace          string  `json:"claimNamespace,omitempty"`
	ClusterName             string  `json:"clusterName,omitempty"`
	ClusterId               string  `json:"clusterId,omitempty"`
	StorageClass            string  `json:"storageClass,omitempty"`
	VolumeType              string  `json:"volumeType,omitempty"`
	CloudProvider           string  `json:"cloudProvider,omitempty"`
	Region                  string  `json:"region,omitempty"`
	StorageCost             float64 `json:"storageCost,omitempty"`
	StorageActualIdleCost   float64 `json:"storageActualIdleCost,omitempty"`
	StorageUnallocatedCost  float64 `json:"storageUnallocatedCost,omitempty"`
	Capacity                float64 `json:"capacity,omitempty"`
	StorageRequest          float64 `json:"storageRequest,omitempty"`
	StorageUtilizationValue float64 `json:"storageUtilizationValue,omitempty"`
	CreateTime              int64   `json:"createTime,omitempty"`
	DeleteTime              int64   `json:"deleteTime,omitempty"`
}
