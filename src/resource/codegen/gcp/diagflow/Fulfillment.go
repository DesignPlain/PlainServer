package diagflow

import types "libds/gcp/types"

type Fulfillment struct {
	/*
	   The field defines whether the fulfillment is enabled for certain features.
	   Structure is documented below.
	*/
	Features []types.Diagflow_FulfillmentFeature `json:"features,omitempty" yaml:"features,omitempty"`

	/*
	   Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	   Structure is documented below.
	*/
	GenericWebService types.Diagflow_FulfillmentGenericWebService `json:"genericWebService,omitempty" yaml:"genericWebService,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   The human-readable name of the fulfillment, unique within the agent.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Whether fulfillment is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
