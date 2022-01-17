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

type PolicyMetadata struct {
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	SerializedSize            int32                  `json:"serializedSize,omitempty"`
	ParserForType             *ParserPolicyMetadata  `json:"parserForType,omitempty"`
	DefaultInstanceForType    *PolicyMetadata        `json:"defaultInstanceForType,omitempty"`
	IdentifierBytes           *ByteString            `json:"identifierBytes,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
	Error_                    string                 `json:"error,omitempty"`
	PolicyId                  string                 `json:"policyId,omitempty"`
	OrgId                     string                 `json:"orgId,omitempty"`
	PolicyIdBytes             *ByteString            `json:"policyIdBytes,omitempty"`
	PolicyName                string                 `json:"policyName,omitempty"`
	PolicyNameBytes           *ByteString            `json:"policyNameBytes,omitempty"`
	Severity                  string                 `json:"severity,omitempty"`
	SeverityBytes             *ByteString            `json:"severityBytes,omitempty"`
	DenyMessagesList          []string               `json:"denyMessagesList,omitempty"`
	DenyMessagesCount         int32                  `json:"denyMessagesCount,omitempty"`
	StatusBytes               *ByteString            `json:"statusBytes,omitempty"`
	AccountId                 string                 `json:"accountId,omitempty"`
	AccountIdBytes            *ByteString            `json:"accountIdBytes,omitempty"`
	OrgIdBytes                *ByteString            `json:"orgIdBytes,omitempty"`
	ProjectId                 string                 `json:"projectId,omitempty"`
	ProjectIdBytes            *ByteString            `json:"projectIdBytes,omitempty"`
	Created                   int64                  `json:"created,omitempty"`
	Updated                   int64                  `json:"updated,omitempty"`
	ErrorBytes                *ByteString            `json:"errorBytes,omitempty"`
	Identifier                string                 `json:"identifier,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	MemoizedSerializedSize    int32                  `json:"memoizedSerializedSize,omitempty"`
}
