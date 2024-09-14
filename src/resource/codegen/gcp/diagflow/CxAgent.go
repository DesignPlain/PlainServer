package diagflow

import types "libds/gcp/types"

type CxAgent struct {
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty" yaml:"supportedLanguageCodes,omitempty"`

	/*
	   Settings related to speech synthesizing.
	   Structure is documented below.
	*/
	TextToSpeechSettings types.Diagflow_CxAgentTextToSpeechSettings `json:"textToSpeechSettings,omitempty" yaml:"textToSpeechSettings,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Git integration settings for this agent.
	   Structure is documented below.
	*/
	GitIntegrationSettings types.Diagflow_CxAgentGitIntegrationSettings `json:"gitIntegrationSettings,omitempty" yaml:"gitIntegrationSettings,omitempty"`

	/*
	   The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/cx/docs/reference/language)
	   for a list of the currently supported language codes. This field cannot be updated after creation.
	*/
	DefaultLanguageCode string `json:"defaultLanguageCode,omitempty" yaml:"defaultLanguageCode,omitempty"`

	// The human-readable name of the agent, unique within the location.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection bool `json:"enableSpellCorrection,omitempty" yaml:"enableSpellCorrection,omitempty"`

	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging bool `json:"enableStackdriverLogging,omitempty" yaml:"enableStackdriverLogging,omitempty"`

	/*
	   The name of the location this agent is located in.
	   > --Note:-- The first time you are deploying an Agent in your project you must configure location settings.
	   This is a one time step but at the moment you can only [configure location settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	   Another options is to use global location so you don't need to manually configure location settings.
	*/
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   Settings related to speech recognition.
	   Structure is documented below.
	*/
	SpeechToTextSettings types.Diagflow_CxAgentSpeechToTextSettings `json:"speechToTextSettings,omitempty" yaml:"speechToTextSettings,omitempty"`

	/*
	   Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at the higher level.
	   Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	   Structure is documented below.
	*/
	AdvancedSettings types.Diagflow_CxAgentAdvancedSettings `json:"advancedSettings,omitempty" yaml:"advancedSettings,omitempty"`

	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo integration.
	AvatarUri string `json:"avatarUri,omitempty" yaml:"avatarUri,omitempty"`

	/*
	   The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	   Europe/Paris.


	   - - -
	*/
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings string `json:"securitySettings,omitempty" yaml:"securitySettings,omitempty"`
}
