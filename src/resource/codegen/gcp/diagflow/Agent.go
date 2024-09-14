package diagflow

type Agent struct {
	/*
	   The URI of the agent's avatar, which are used throughout the Dialogflow console. When an image URL is entered
	   into this field, the Dialogflow will save the image in the backend. The address of the backend image returned
	   from the API will be shown in the [avatarUriBackend] field.
	*/
	AvatarUri string `json:"avatarUri,omitempty" yaml:"avatarUri,omitempty"`

	/*
	   The default language of the agent as a language tag. [See Language Support](https://cloud.google.com/dialogflow/docs/reference/language)
	   for a list of the currently supported language codes. This field cannot be updated after creation.
	*/
	DefaultLanguageCode string `json:"defaultLanguageCode,omitempty" yaml:"defaultLanguageCode,omitempty"`

	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of this agent.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Determines whether this agent should log conversation queries.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`

	/*
	   Determines how intents are detected from user queries.
	   - MATCH_MODE_HYBRID: Best for agents with a small number of examples in intents and/or wide use of templates
	   syntax and composite entities.
	   - MATCH_MODE_ML_ONLY: Can be used for agents with a large number of examples in intents, especially the ones
	   using @sys.any or very large developer entities.
	   Possible values are: `MATCH_MODE_HYBRID`, `MATCH_MODE_ML_ONLY`.
	*/
	MatchMode string `json:"matchMode,omitempty" yaml:"matchMode,omitempty"`

	/*
	   The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	   Europe/Paris.


	   - - -
	*/
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	/*
	   API version displayed in Dialogflow console. If not specified, V2 API is assumed. Clients are free to query
	   different service endpoints for different API versions. However, bots connectors and webhook calls will follow
	   the specified API version.
	   - API_VERSION_V1: Legacy V1 API.
	   - API_VERSION_V2: V2 API.
	   - API_VERSION_V2_BETA_1: V2beta1 API.
	   Possible values are: `API_VERSION_V1`, `API_VERSION_V2`, `API_VERSION_V2_BETA_1`.
	*/
	ApiVersion string `json:"apiVersion,omitempty" yaml:"apiVersion,omitempty"`

	/*
	   To filter out false positive results and still get variety in matched natural language inputs for your agent,
	   you can tune the machine learning classification threshold. If the returned score value is less than the threshold
	   value, then a fallback intent will be triggered or, if there are no fallback intents defined, no intent will be
	   triggered. The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the
	   default of 0.3 is used.
	*/
	ClassificationThreshold float64 `json:"classificationThreshold,omitempty" yaml:"classificationThreshold,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The list of all languages supported by this agent (except for the defaultLanguageCode).
	SupportedLanguageCodes []string `json:"supportedLanguageCodes,omitempty" yaml:"supportedLanguageCodes,omitempty"`

	/*
	   The agent tier. If not specified, TIER_STANDARD is assumed.
	   - TIER_STANDARD: Standard tier.
	   - TIER_ENTERPRISE: Enterprise tier (Essentials).
	   - TIER_ENTERPRISE_PLUS: Enterprise tier (Plus).
	   NOTE: Due to consistency issues, the provider will not read this field from the API. Drift is possible between
	   the the provider state and Dialogflow if the agent tier is changed outside of the provider.
	*/
	Tier string `json:"tier,omitempty" yaml:"tier,omitempty"`
}
