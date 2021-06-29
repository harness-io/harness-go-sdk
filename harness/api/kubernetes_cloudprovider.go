package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetKubernetesCloudProviderByName(name string) (*graphql.KubernetesCloudProvider, error) {
	cp := &graphql.KubernetesCloudProvider{}
	err := c.getCloudProviderByName(name, getKubernetesCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreateKubernetesCloudProvider(provider *graphql.KubernetesCloudProvider) (*graphql.KubernetesCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.KubernetesCluster,
		K8sCloudProvider:  provider,
	}

	resp := &graphql.KubernetesCloudProvider{}
	err := c.createCloudProvider(input, getKubernetesCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getKubernetesCloudProviderFields() string {
	return fmt.Sprintf(`
		%[1]s
		... on KubernetesCloudProvider {
			%[2]s
			skipK8sEventCollection
		}
`, commonCloudProviderFields, ceHealthStatusFields)
}
