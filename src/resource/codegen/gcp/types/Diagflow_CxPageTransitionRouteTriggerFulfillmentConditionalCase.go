package types

type Diagflow_CxPageTransitionRouteTriggerFulfillmentConditionalCase struct {
	/*
	   A JSON encoded list of cascading if-else conditions. Cases are mutually exclusive. The first one with a matching condition is selected, all the rest ignored.
	   See [Case](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/Fulfillment#case) for the schema.
	*/
	Cases string `json:"cases,omitempty" yaml:"cases,omitempty"`
}
