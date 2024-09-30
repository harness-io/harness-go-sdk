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

type ClusterOrchDistributionSelector string
type ClusterOrchNodeDistributionStrategy string

type ClusterOrchConfig struct {
	SpotDistribution     ClusterOrchDistributionSelector     `json:"spot_distribution"`
	NodeDeletionDelay    int                                 `json:"node_deletion_delay"`
	DistributionStrategy ClusterOrchNodeDistributionStrategy `json:"distribution_strategy"`
	BaseOnDemandCapacity int                                 `json:"base_on_demand_capacity"`
	SpotSplit            int                                 `json:"spot_split"`
	OnDemandSplit        int                                 `json:"on_demand_split"`
}

type ClusterOrchestratorUserConfig struct {
	ClusterEndPoint string `json:"cluster_endpoint"`
	K8sConnID       string `json:"k8s_connector_id"`
}

type ClusterOrchestrator struct {
	ID                  string                         `json:"id" db:"id"`
	AccountID           string                         `json:"account_id" db:"account_id"`
	K8sConnectorID      string                         `json:"k8s_connector_id" db:"k8s_connector_id"`
	CloudAccID          string                         `json:"cloud_account_id" db:"cloud_account_id"`
	Name                string                         `json:"name" db:"name"`
	Config              *ClusterOrchConfig             `json:"config" db:"config"`
	UserConfig          *ClusterOrchestratorUserConfig `json:"user_config,omitempty" db:"user_config"`
	OptimizationEnabled bool                           `json:"optimization_enabled" db:"optimization_enabled"`
	Disabled            bool                           `json:"-" db:"disabled"`
}
