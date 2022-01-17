/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// This has the Connector details defined in Harness

import (
	"encoding/json"
)

type ConnectorInfo struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier"`
	Description       string                   `json:"description,omitempty"`
	OrgIdentifier     string                   `json:"orgIdentifier,omitempty"`
	ProjectIdentifier string                   `json:"projectIdentifier,omitempty"`
	Tags              map[string]string        `json:"tags,omitempty"`
	Type_             ConnectorType            `json:"type"`
	AppDynamics       *AppDynamicsConnectorDto `json:"-"`
	Artifactory       *ArtifactoryConnector    `json:"-"`
	Aws               *AwsConnector            `json:"-"`
	AwsCC             *CeAwsConnector          `json:"-"`
	AwsKms            *AwsKmsConnector         `json:"-"`
	AwsSecretManager  *AwsSecretManager        `json:"-"`
	BitBucket         *BitbucketConnector      `json:"-"`
	Datadog           *DatadogConnectorDto     `json:"-"`
	DockerRegistry    *DockerConnector         `json:"-"`
	Dynatrace         *DynatraceConnectorDto   `json:"-"`
	Gcp               *GcpConnector            `json:"-"`
	Git               *GitConfig               `json:"-"`
	Github            *GithubConnector         `json:"-"`
	Gitlab            *GitlabConnector         `json:"-"`
	HttpHelm          *HttpHelmConnector       `json:"-"`
	Jira              *JiraConnector           `json:"-"`
	K8sCluster        *KubernetesClusterConfig `json:"-"`
	NewRelic          *NewRelicConnectorDto    `json:"-"`
	Nexus             *NexusConnector          `json:"-"`
	PagerDuty         *PagerDutyConnectorDto   `json:"-"`
	Prometheus        *PrometheusConnectorDto  `json:"-"`
	Splunk            *SplunkConnector         `json:"-"`
	SumoLogic         *SumoLogicConnectorDto   `json:"-"`
	Spec              json.RawMessage          `json:"spec"`
}
