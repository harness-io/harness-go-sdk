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

type AccessPointMeta struct {
	Error_           string              `json:"error,omitempty"`
	CertificateId    string              `json:"certificate_id,omitempty"`
	SecurityGroups   []string            `json:"security_groups,omitempty"`
	Dns              *AccessPointMetaDns `json:"dns,omitempty"`
	AlbArn           string              `json:"albArn,omitempty"`
	ResourceGroup    string              `json:"resource_group,omitempty"`
	FeIpId           string              `json:"fe_ip_id,omitempty"`
	SubnetId         string              `json:"subnet_id,omitempty"`
	Size             string              `json:"size,omitempty"`
	AppGatewayId     string              `json:"app_gateway_id,omitempty"`
	SubnetName       string              `json:"subnet_name,omitempty"`
	FeIpName         string              `json:"fe_ip_name,omitempty"`
	Certificate      *CertificateData    `json:"certificate,omitempty"`
	FuncRegion       string              `json:"func_region,omitempty"`
	Zone             string              `json:"zone,omitempty"`
	MachineType      string              `json:"machine_type,omitempty"`
	ApiKey           string              `json:"api-key,omitempty"`
	AllocateStaticIp bool                `json:"allocate_static_ip,omitempty"`
	Keypair          string              `json:"keypair,omitempty"`
	Certificates     *CertificatesData   `json:"certificates,omitempty"`
}
