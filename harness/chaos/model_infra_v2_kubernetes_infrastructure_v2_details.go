/*
 * Chaos Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InfraV2KubernetesInfrastructureV2Details struct {
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy *InfraV2UserDetails `json:"createdBy,omitempty"`
	DeploymentType string `json:"deploymentType,omitempty"`
	Description string `json:"description,omitempty"`
	DiscoveryAgentID string `json:"discoveryAgentID,omitempty"`
	EnvironmentID string `json:"environmentID,omitempty"`
	HarnessInfraType string `json:"harnessInfraType,omitempty"`
	Identifier *InfraV2Identifiers `json:"identifier,omitempty"`
	Identity string `json:"identity,omitempty"`
	InfraID string `json:"infraID,omitempty"`
	InfraNamespace string `json:"infraNamespace,omitempty"`
	InfraScope *InfraV2InfraScope `json:"infraScope,omitempty"`
	InfraType *InfraV2InfraType `json:"infraType,omitempty"`
	InsecureSkipVerify bool `json:"insecureSkipVerify,omitempty"`
	InstallationType *InfraV2InstallationType `json:"installationType,omitempty"`
	IsChaosEnabled bool `json:"isChaosEnabled,omitempty"`
	K8sConnectorID string `json:"k8sConnectorID,omitempty"`
	LastHeartbeat int32 `json:"lastHeartbeat,omitempty"`
	LastWorkflowTimestamp string `json:"lastWorkflowTimestamp,omitempty"`
	Mtls *InfraV2MtlsConfiguration `json:"mtls,omitempty"`
	Name string `json:"name,omitempty"`
	NoOfSchedules int32 `json:"noOfSchedules,omitempty"`
	NoOfWorkflows int32 `json:"noOfWorkflows,omitempty"`
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	PlatformName string `json:"platformName,omitempty"`
	Proxy *InfraV2ProxyConfiguration `json:"proxy,omitempty"`
	RunAsGroup int32 `json:"runAsGroup,omitempty"`
	RunAsUser int32 `json:"runAsUser,omitempty"`
	ServiceAccount string `json:"serviceAccount,omitempty"`
	Status *InfraV2InfraStatus `json:"status,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Tolerations []V1Toleration `json:"tolerations,omitempty"`
	UpdateStatus *InfraV2UpdateStatus `json:"updateStatus,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy *InfraV2UserDetails `json:"updatedBy,omitempty"`
	Upgrade *InfraV2Upgrade `json:"upgrade,omitempty"`
	Version string `json:"version,omitempty"`
}
