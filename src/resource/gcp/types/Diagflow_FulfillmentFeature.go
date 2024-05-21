package types

type Diagflow_FulfillmentFeature struct {
	/*
	   The type of the feature that enabled for fulfillment.
	   - SMALLTALK: Fulfillment is enabled for SmallTalk.
	   Possible values are: `SMALLTALK`.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
