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

// Describes a distribution rule
type Distribution struct {
	// The attribute to use when distributing targets across buckets
	BucketBy string `json:"bucketBy"`
	// A list of variations and the weight that should be given to each
	Variations []WeightedVariation `json:"variations"`
}
