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

type FileDescriptorProto struct {
	UnknownFields             *UnknownFieldSet                  `json:"unknownFields,omitempty"`
	SerializedSize            int32                             `json:"serializedSize,omitempty"`
	ParserForType             *ParserFileDescriptorProto        `json:"parserForType,omitempty"`
	DefaultInstanceForType    *FileDescriptorProto              `json:"defaultInstanceForType,omitempty"`
	Options                   *FileOptions                      `json:"options,omitempty"`
	Initialized               bool                              `json:"initialized,omitempty"`
	Name                      string                            `json:"name,omitempty"`
	Package_                  string                            `json:"package,omitempty"`
	PublicDependencyCount     int32                             `json:"publicDependencyCount,omitempty"`
	DependencyCount           int32                             `json:"dependencyCount,omitempty"`
	MessageTypeCount          int32                             `json:"messageTypeCount,omitempty"`
	ServiceCount              int32                             `json:"serviceCount,omitempty"`
	Syntax                    string                            `json:"syntax,omitempty"`
	NameBytes                 *ByteString                       `json:"nameBytes,omitempty"`
	PackageBytes              *ByteString                       `json:"packageBytes,omitempty"`
	DependencyList            []string                          `json:"dependencyList,omitempty"`
	PublicDependencyList      []int32                           `json:"publicDependencyList,omitempty"`
	WeakDependencyList        []int32                           `json:"weakDependencyList,omitempty"`
	WeakDependencyCount       int32                             `json:"weakDependencyCount,omitempty"`
	MessageTypeList           []DescriptorProto                 `json:"messageTypeList,omitempty"`
	MessageTypeOrBuilderList  []DescriptorProtoOrBuilder        `json:"messageTypeOrBuilderList,omitempty"`
	EnumTypeList              []EnumDescriptorProto             `json:"enumTypeList,omitempty"`
	EnumTypeOrBuilderList     []EnumDescriptorProtoOrBuilder    `json:"enumTypeOrBuilderList,omitempty"`
	ServiceList               []ServiceDescriptorProto          `json:"serviceList,omitempty"`
	ServiceOrBuilderList      []ServiceDescriptorProtoOrBuilder `json:"serviceOrBuilderList,omitempty"`
	ExtensionList             []FieldDescriptorProto            `json:"extensionList,omitempty"`
	ExtensionOrBuilderList    []FieldDescriptorProtoOrBuilder   `json:"extensionOrBuilderList,omitempty"`
	OptionsOrBuilder          *FileOptionsOrBuilder             `json:"optionsOrBuilder,omitempty"`
	SourceCodeInfo            *SourceCodeInfo                   `json:"sourceCodeInfo,omitempty"`
	SourceCodeInfoOrBuilder   *SourceCodeInfoOrBuilder          `json:"sourceCodeInfoOrBuilder,omitempty"`
	SyntaxBytes               *ByteString                       `json:"syntaxBytes,omitempty"`
	EnumTypeCount             int32                             `json:"enumTypeCount,omitempty"`
	ExtensionCount            int32                             `json:"extensionCount,omitempty"`
	InitializationErrorString string                            `json:"initializationErrorString,omitempty"`
	DescriptorForType         *Descriptor                       `json:"descriptorForType,omitempty"`
	AllFields                 map[string]interface{}            `json:"allFields,omitempty"`
	MemoizedSerializedSize    int32                             `json:"memoizedSerializedSize,omitempty"`
}
