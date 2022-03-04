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

type FieldOptionsOrBuilder struct {
	Packed                           bool                           `json:"packed,omitempty"`
	Jstype                           string                         `json:"jstype,omitempty"`
	Ctype                            string                         `json:"ctype,omitempty"`
	Deprecated                       bool                           `json:"deprecated,omitempty"`
	Lazy                             bool                           `json:"lazy,omitempty"`
	UninterpretedOptionList          []UninterpretedOption          `json:"uninterpretedOptionList,omitempty"`
	UninterpretedOptionCount         int32                          `json:"uninterpretedOptionCount,omitempty"`
	UninterpretedOptionOrBuilderList []UninterpretedOptionOrBuilder `json:"uninterpretedOptionOrBuilderList,omitempty"`
	Weak                             bool                           `json:"weak,omitempty"`
	DefaultInstanceForType           *Message                       `json:"defaultInstanceForType,omitempty"`
	AllFields                        map[string]interface{}         `json:"allFields,omitempty"`
	InitializationErrorString        string                         `json:"initializationErrorString,omitempty"`
	DescriptorForType                *Descriptor                    `json:"descriptorForType,omitempty"`
	UnknownFields                    *UnknownFieldSet               `json:"unknownFields,omitempty"`
	Initialized                      bool                           `json:"initialized,omitempty"`
}
