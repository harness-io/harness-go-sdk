/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type HelmChartManifest struct {
	Store *StoreConfigWrapper `json:"store,omitempty"`
	ChartName string `json:"chartName,omitempty"`
	ChartVersion string `json:"chartVersion,omitempty"`
	HelmVersion string `json:"helmVersion,omitempty"`
	ValuesPaths []string `json:"valuesPaths,omitempty"`
	SkipResourceVersioning bool `json:"skipResourceVersioning,omitempty"`
	EnableDeclarativeRollback *ParameterFieldBoolean `json:"enableDeclarativeRollback,omitempty"`
	CommandFlags []HelmManifestCommandFlag `json:"commandFlags,omitempty"`
	SubChartPath string `json:"subChartPath,omitempty"`
	Metadata string `json:"metadata,omitempty"`
}
