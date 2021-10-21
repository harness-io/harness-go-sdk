/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type JsonNode struct {
	Float bool `json:"float,omitempty"`
	NodeType string `json:"nodeType,omitempty"`
	Short bool `json:"short,omitempty"`
	Number bool `json:"number,omitempty"`
	Binary bool `json:"binary,omitempty"`
	Pojo bool `json:"pojo,omitempty"`
	ValueNode bool `json:"valueNode,omitempty"`
	ContainerNode bool `json:"containerNode,omitempty"`
	MissingNode bool `json:"missingNode,omitempty"`
	IntegralNumber bool `json:"integralNumber,omitempty"`
	FloatingPointNumber bool `json:"floatingPointNumber,omitempty"`
	Long bool `json:"long,omitempty"`
	Double bool `json:"double,omitempty"`
	BigDecimal bool `json:"bigDecimal,omitempty"`
	BigInteger bool `json:"bigInteger,omitempty"`
	Textual bool `json:"textual,omitempty"`
	Boolean bool `json:"boolean,omitempty"`
	Int_ bool `json:"int,omitempty"`
	Object bool `json:"object,omitempty"`
	Array bool `json:"array,omitempty"`
	Null bool `json:"null,omitempty"`
}
