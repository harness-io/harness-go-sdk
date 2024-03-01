/*
 * API Specification
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package code

type RulesRuleUidBody struct {
	Definition  *OpenapiRuleDefinition `json:"definition,omitempty"`
	Description string                 `json:"description,omitempty"`
	Identifier  string                 `json:"identifier,omitempty"`
	Pattern     *ProtectionPattern     `json:"pattern,omitempty"`
	State       *EnumRuleState         `json:"state,omitempty"`
	Type_       *OpenapiRuleType       `json:"type,omitempty"`
	Uid         string                 `json:"uid,omitempty"`
}
