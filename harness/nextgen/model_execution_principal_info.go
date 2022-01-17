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

type ExecutionPrincipalInfo struct {
	UnknownFields             *UnknownFieldSet              `json:"unknownFields,omitempty"`
	SerializedSize            int32                         `json:"serializedSize,omitempty"`
	ParserForType             *ParserExecutionPrincipalInfo `json:"parserForType,omitempty"`
	DefaultInstanceForType    *ExecutionPrincipalInfo       `json:"defaultInstanceForType,omitempty"`
	Initialized               bool                          `json:"initialized,omitempty"`
	Principal                 string                        `json:"principal,omitempty"`
	PrincipalBytes            *ByteString                   `json:"principalBytes,omitempty"`
	PrincipalTypeValue        int32                         `json:"principalTypeValue,omitempty"`
	PrincipalType             string                        `json:"principalType,omitempty"`
	ShouldValidateRbac        bool                          `json:"shouldValidateRbac,omitempty"`
	InitializationErrorString string                        `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                   `json:"descriptorForType,omitempty"`
	AllFields                 map[string]interface{}        `json:"allFields,omitempty"`
	MemoizedSerializedSize    int32                         `json:"memoizedSerializedSize,omitempty"`
}
