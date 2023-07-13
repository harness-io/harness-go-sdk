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

import "encoding/json"

// This entity contains the details for Service Now Authentication
type ServiceNowAuthentication struct {
	Type_                      ServiceNowAuthType          `json:"type"`
	ServiceNowUserNamePassword *ServiceNowUserNamePassword `json:"-"`
	ServiceNowAdfs             *ServiceNowAdfs             `json:"-"`
	ServiceNowRefreshToken     *ServiceNowRefreshToken     `json:"-"`
	Spec                       json.RawMessage             `json:"spec"`
}
