package api

import (
	"fmt"

	"github.com/harness-io/harness-go-sdk/harness/api/graphql"
)

func (c *CloudProviderClient) GetPcfCloudProviderByName(name string) (*graphql.PcfCloudProvider, error) {
	cp := &graphql.PcfCloudProvider{}
	err := c.getCloudProviderByName(name, getPcfCloudProviderFields(), &cp)
	if err != nil {
		return nil, err
	}

	return cp, nil
}

func (c *CloudProviderClient) CreatePcfCloudProvider(provider *graphql.PcfCloudProvider) (*graphql.PcfCloudProvider, error) {
	input := &graphql.CreateCloudProviderInput{
		CloudProviderType: graphql.CloudProviderTypes.Pcf,
		PcfCloudProvider:  provider,
	}

	resp := &graphql.PcfCloudProvider{}
	err := c.createCloudProvider(input, getPcfCloudProviderFields(), resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getPcfCloudProviderFields() string {
	return fmt.Sprintf(`
	... on PcfCloudProvider {
		%[1]s
	}
	`, commonCloudProviderFields)
}
