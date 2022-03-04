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

type RerunInfoOrBuilder struct {
	PrevExecutionId           string                 `json:"prevExecutionId,omitempty"`
	PrevExecutionIdBytes      *ByteString            `json:"prevExecutionIdBytes,omitempty"`
	PrevTriggerTypeValue      int32                  `json:"prevTriggerTypeValue,omitempty"`
	PrevTriggerType           string                 `json:"prevTriggerType,omitempty"`
	RootExecutionId           string                 `json:"rootExecutionId,omitempty"`
	RootTriggerType           string                 `json:"rootTriggerType,omitempty"`
	RootTriggerTypeValue      int32                  `json:"rootTriggerTypeValue,omitempty"`
	RootExecutionIdBytes      *ByteString            `json:"rootExecutionIdBytes,omitempty"`
	AllFields                 map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor            `json:"descriptorForType,omitempty"`
	UnknownFields             *UnknownFieldSet       `json:"unknownFields,omitempty"`
	DefaultInstanceForType    *Message               `json:"defaultInstanceForType,omitempty"`
	Initialized               bool                   `json:"initialized,omitempty"`
}
