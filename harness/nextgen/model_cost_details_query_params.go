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

// Common query parameters for all cost details APIs
type CostDetailsQueryParams struct {
	// Filters to be applied on the response.
	Filters []FieldFilter `json:"filters,omitempty"`
	// Fields on which the response will be grouped by.
	GroupBy []string `json:"groupBy,omitempty"`
	// Only applicable for Time Series Endpoints, defaults to DAY
	TimeResolution string `json:"timeResolution,omitempty"`
	// Limit on the number of cost values returned, 0 by default.
	Limit int32 `json:"limit,omitempty"`
	// Order of sorting on cost, Descending by default.
	SortOrder string `json:"sortOrder,omitempty"`
	// Offset on the cost values returned, 10 by default.
	Offset int32 `json:"offset,omitempty"`
	// Skip Rounding off the cost values returned, false by default.
	SkipRoundOff bool `json:"skipRoundOff,omitempty"`
}
