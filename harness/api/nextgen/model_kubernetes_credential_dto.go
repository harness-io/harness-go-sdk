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

import (
	"encoding/json"
	"fmt"
)

type KubernetesCredentialDto struct {
	ManualCredentials *KubernetesClusterDetailsDto `json:"-"`
	Type_             string                       `json:"type"`
	Spec              json.RawMessage              `json:"spec,omitempty"`
}

func (a *KubernetesCredentialDto) UnmarshalJSON(data []byte) error {

	type Alias KubernetesCredentialDto

	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	switch a.Type_ {
	case KubernetesCredentialTypes.InheritFromDelegate.String():
		// do nothing
	case KubernetesCredentialTypes.ManualConfig.String():
		err = json.Unmarshal(a.Spec, &a.ManualCredentials)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	return err
}

func (a *KubernetesCredentialDto) MarshalJSON() ([]byte, error) {
	type Alias KubernetesCredentialDto

	var spec []byte
	var err error

	switch a.Type_ {
	case KubernetesCredentialTypes.InheritFromDelegate.String():
		// do nothing
	case KubernetesCredentialTypes.ManualConfig.String():
		spec, err = json.Marshal(a.ManualCredentials)
	default:
		panic(fmt.Sprintf("unknown connector type %s", a.Type_))
	}

	if err != nil {
		return nil, err
	}

	a.Spec = json.RawMessage(spec)

	return json.Marshal((*Alias)(a))
}
