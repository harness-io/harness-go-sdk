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

// Contains information of filters for Resource Group
type ResourceGroupFilter struct {
	// Filter by account identifier
	AccountIdentifier string `json:"accountIdentifier"`
	// Filter by organization identifier
	OrgIdentifier string `json:"orgIdentifier,omitempty"`
	// Filter by project identifier
	ProjectIdentifier string `json:"projectIdentifier,omitempty"`
	// Filter resource group matching by identifier/name
	SearchTerm string `json:"searchTerm,omitempty"`
	// Filter by resource group identifiers
	IdentifierFilter []string `json:"identifierFilter,omitempty"`
	// Filter based on whether it has a particular resource
	ResourceSelectorFilterList []ResourceSelectorFilter `json:"resourceSelectorFilterList,omitempty"`
	// Filter based on whether the resource group is Harness managed
	ManagedFilter string `json:"managedFilter,omitempty"`
}
