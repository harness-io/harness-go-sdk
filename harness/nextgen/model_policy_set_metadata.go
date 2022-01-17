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

type PolicySetMetadata struct {
	UnknownFields               *UnknownFieldSet          `json:"unknownFields,omitempty"`
	SerializedSize              int32                     `json:"serializedSize,omitempty"`
	ParserForType               *ParserPolicySetMetadata  `json:"parserForType,omitempty"`
	DefaultInstanceForType      *PolicySetMetadata        `json:"defaultInstanceForType,omitempty"`
	IdentifierBytes             *ByteString               `json:"identifierBytes,omitempty"`
	Status                      string                    `json:"status,omitempty"`
	PolicySetId                 string                    `json:"policySetId,omitempty"`
	PolicySetIdBytes            *ByteString               `json:"policySetIdBytes,omitempty"`
	Deny                        bool                      `json:"deny,omitempty"`
	PolicyMetadataList          []PolicyMetadata          `json:"policyMetadataList,omitempty"`
	PolicyMetadataOrBuilderList []PolicyMetadataOrBuilder `json:"policyMetadataOrBuilderList,omitempty"`
	PolicyMetadataCount         int32                     `json:"policyMetadataCount,omitempty"`
	PolicySetName               string                    `json:"policySetName,omitempty"`
	PolicySetNameBytes          *ByteString               `json:"policySetNameBytes,omitempty"`
	Initialized                 bool                      `json:"initialized,omitempty"`
	OrgId                       string                    `json:"orgId,omitempty"`
	StatusBytes                 *ByteString               `json:"statusBytes,omitempty"`
	AccountId                   string                    `json:"accountId,omitempty"`
	AccountIdBytes              *ByteString               `json:"accountIdBytes,omitempty"`
	OrgIdBytes                  *ByteString               `json:"orgIdBytes,omitempty"`
	ProjectId                   string                    `json:"projectId,omitempty"`
	ProjectIdBytes              *ByteString               `json:"projectIdBytes,omitempty"`
	Created                     int64                     `json:"created,omitempty"`
	Identifier                  string                    `json:"identifier,omitempty"`
	InitializationErrorString   string                    `json:"initializationErrorString,omitempty"`
	DescriptorForType           *Descriptor               `json:"descriptorForType,omitempty"`
	AllFields                   map[string]interface{}    `json:"allFields,omitempty"`
	MemoizedSerializedSize      int32                     `json:"memoizedSerializedSize,omitempty"`
}
