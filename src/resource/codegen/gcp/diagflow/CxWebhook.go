package diagflow

import types "libds/gcp/types"

type CxWebhook struct {
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging bool `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`

	/*
	   Configuration for a Service Directory service.
	   Structure is documented below.
	*/
	ServiceDirectory types.Diagflow_CxWebhookServiceDirectory `json:"serviceDirectory,omitempty" yaml:"serviceDirectory,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings string `json:"securitySettings,omitempty" yaml:"securitySettings,omitempty"`

	// Webhook execution timeout.
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Indicates whether the webhook is disabled.
	Disabled bool `json:"disabled,omitempty" yaml:"disabled,omitempty"`

	/*
	   The human-readable name of the webhook, unique within the agent.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection bool `json:"enableSpellCorrection,omitempty" yaml:"enableSpellCorrection,omitempty"`

	/*
	   Configuration for a generic web service.
	   Structure is documented below.
	*/
	GenericWebService types.Diagflow_CxWebhookGenericWebService `json:"genericWebService,omitempty" yaml:"genericWebService,omitempty"`

	/*
	   The agent to create a webhook for.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
