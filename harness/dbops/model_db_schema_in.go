/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the DB Service
 *
 * API version: 1.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package dbops

// Database Schema Request
type DbSchemaIn struct {
	// identifier of the database schema
	Identifier string `json:"identifier"`
	// name of the database schema
	Name string `json:"name"`
	// tags attached to the database schema
	Tags      map[string]string `json:"tags,omitempty"`
	Changelog *Changelog        `json:"changelog"`
	// harness service corresponding to database schema
	Service string `json:"service,omitempty"`
}
