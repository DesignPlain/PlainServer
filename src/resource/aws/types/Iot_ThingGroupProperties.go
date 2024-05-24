package types

type Iot_ThingGroupProperties struct {
	// The Thing Group attributes. Defined below.
	AttributePayload Iot_ThingGroupPropertiesAttributePayload `json:"attributePayload,omitempty" yaml:"attributePayload,omitempty"`

	// A description of the Thing Group.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
