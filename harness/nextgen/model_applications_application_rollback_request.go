/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type ApplicationsApplicationRollbackRequest struct {
	Name   string `json:"name,omitempty"`
	Id     string `json:"id,omitempty"`
	DryRun bool   `json:"dryRun,omitempty"`
	Prune  bool   `json:"prune,omitempty"`
}
