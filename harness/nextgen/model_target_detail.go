/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub  # Authentication  <!-- ReDoc-Inject: <security-definitions> -->
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

// Details of which Target Groups (Segments) a target is included in or excluded from
type TargetDetail struct {
	// A list of target groups (segments) that the target is excluded from.
	ExcludedSegments []TargetDetailSegment `json:"excludedSegments,omitempty"`
	// The unique identifier for the target
	Identifier string `json:"identifier"`
	// A list of target groups (segments) that the target is included in.
	IncludedSegments []TargetDetailSegment `json:"includedSegments,omitempty"`
	// A list of target groups (segments) that the target is included in via group rules.
	RuleSegments []TargetDetailSegment `json:"ruleSegments,omitempty"`
}
