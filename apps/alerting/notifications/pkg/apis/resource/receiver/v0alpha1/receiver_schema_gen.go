//
// Code generated by grafana-app-sdk. DO NOT EDIT.
//

package v0alpha1

import (
	"fmt"

	"github.com/grafana/grafana-app-sdk/resource"
)

// schema is unexported to prevent accidental overwrites
var (
	schemaReceiver = resource.NewSimpleSchema("notifications.alerting.grafana.app", "v0alpha1", &Receiver{}, &ReceiverList{}, resource.WithKind("Receiver"),
		resource.WithPlural("receivers"), resource.WithScope(resource.NamespacedScope), resource.WithSelectableFields([]resource.SelectableField{{
			FieldSelector: "spec.title",
			FieldValueFunc: func(o resource.Object) (string, error) {
				cast, ok := o.(*Receiver)
				if !ok {
					return "", fmt.Errorf("provided object must be of type *Receiver")
				}
				return cast.Spec.Title, nil
			},
		},
		}))
	kindReceiver = resource.Kind{
		Schema: schemaReceiver,
		Codecs: map[resource.KindEncoding]resource.Codec{
			resource.KindEncodingJSON: &JSONCodec{},
		},
	}
)

// Kind returns a resource.Kind for this Schema with a JSON codec
func Kind() resource.Kind {
	return kindReceiver
}

// Schema returns a resource.SimpleSchema representation of Receiver
func Schema() *resource.SimpleSchema {
	return schemaReceiver
}

// Interface compliance checks
var _ resource.Schema = kindReceiver
