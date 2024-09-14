package diagflow

import types "libds/gcp/types"

type CxIntent struct {
	/*
	   The human-readable name of the intent, unique within the agent.


	   - - -
	*/
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	   Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	   To manage the fallback intent, set `is_default_negative_intent = true`
	*/
	IsFallback bool `json:"isFallback,omitempty" yaml:"isFallback,omitempty"`

	/*
	   The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	   Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: - sys-head - sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	   An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.

	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   The priority of this intent. Higher numbers represent higher priorities.
	   If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	   If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	*/
	Priority int `json:"priority,omitempty" yaml:"priority,omitempty"`

	/*
	   The collection of training phrases the agent is trained on to identify the intent.
	   Structure is documented below.
	*/
	TrainingPhrases []types.Diagflow_CxIntentTrainingPhrase `json:"trainingPhrases,omitempty" yaml:"trainingPhrases,omitempty"`

	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Marks this as the [Default Negative Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#negative) for an agent. When you create an agent, a Default Negative Intent is created automatically.
	   The Default Negative Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.

	   > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `is_default_negative_intent = true` because they will compete to control a single Default Negative Intent resource in GCP.
	*/
	IsDefaultNegativeIntent bool `json:"isDefaultNegativeIntent,omitempty" yaml:"isDefaultNegativeIntent,omitempty"`

	/*
	   Marks this as the [Default Welcome Intent](https://cloud.google.com/dialogflow/cx/docs/concept/intent#welcome) for an agent. When you create an agent, a Default Welcome Intent is created automatically.
	   The Default Welcome Intent cannot be deleted; deleting the `gcp.diagflow.CxIntent` resource does nothing to the underlying GCP resources.

	   > Avoid having multiple `gcp.diagflow.CxIntent` resources linked to the same agent with `is_default_welcome_intent = true` because they will compete to control a single Default Welcome Intent resource in GCP.
	*/
	IsDefaultWelcomeIntent bool `json:"isDefaultWelcomeIntent,omitempty" yaml:"isDefaultWelcomeIntent,omitempty"`

	/*
	   The language of the following fields in intent:
	   Intent.training_phrases.parts.text
	   If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	*/
	LanguageCode string `json:"languageCode,omitempty" yaml:"languageCode,omitempty"`

	/*
	   The collection of parameters associated with the intent.
	   Structure is documented below.
	*/
	Parameters []types.Diagflow_CxIntentParameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`

	/*
	   The agent to create an intent for.
	   Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
