package types

type Iot_ThingGroupPropertiesAttributePayload struct {
	// Key-value map.
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}
