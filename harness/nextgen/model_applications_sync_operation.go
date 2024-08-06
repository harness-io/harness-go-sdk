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

// SyncOperation contains details about a sync operation.
type ApplicationsSyncOperation struct {
	// Revision is the revision (Git) or chart version (Helm) which to sync the application to If omitted, will use the revision specified in app spec.
	Revision     string                              `json:"revision,omitempty"`
	Prune        bool                                `json:"prune,omitempty"`
	DryRun       bool                                `json:"dryRun,omitempty"`
	SyncStrategy *ApplicationsSyncStrategy           `json:"syncStrategy,omitempty"`
	Resources    []ApplicationsSyncOperationResource `json:"resources,omitempty"`
	Source       *ApplicationsApplicationSource      `json:"source,omitempty"`
	Manifests    []string                            `json:"manifests,omitempty"`
	SyncOptions  []string                            `json:"syncOptions,omitempty"`
	Sources []ApplicationsApplicationSource `json:"sources,omitempty"`
	// Revisions is the list of revision (Git) or chart version (Helm) which to sync each source in sources field for the application to If omitted, will use the revision specified in app spec.
	Revisions []string `json:"revisions,omitempty"`
}
