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

type TriggerIssuer struct {
	UnknownFields                *UnknownFieldSet       `json:"unknownFields,omitempty"`
	Initialized                  bool                   `json:"initialized,omitempty"`
	AbortPrevConcurrentExecution bool                   `json:"abortPrevConcurrentExecution,omitempty"`
	TriggerRef                   string                 `json:"triggerRef,omitempty"`
	SerializedSize               int32                  `json:"serializedSize,omitempty"`
	ParserForType                *ParserTriggerIssuer   `json:"parserForType,omitempty"`
	DefaultInstanceForType       *TriggerIssuer         `json:"defaultInstanceForType,omitempty"`
	TriggerRefBytes              *ByteString            `json:"triggerRefBytes,omitempty"`
	AllFields                    map[string]interface{} `json:"allFields,omitempty"`
	InitializationErrorString    string                 `json:"initializationErrorString,omitempty"`
	DescriptorForType            *Descriptor            `json:"descriptorForType,omitempty"`
	MemoizedSerializedSize       int32                  `json:"memoizedSerializedSize,omitempty"`
}
