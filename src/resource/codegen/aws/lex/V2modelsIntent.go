package lex

import types "libds/aws/types"

type V2modelsIntent struct {
	// Configuration block for information required to use the AMAZON.KendraSearchIntent intent to connect to an Amazon Kendra index. The AMAZON.KendraSearchIntent intent is called when Amazon Lex can't determine another intent to invoke. See `kendra_configuration`.
	KendraConfiguration types.Lex_V2modelsIntentKendraConfiguration `json:"kendraConfiguration,omitempty" yaml:"kendraConfiguration,omitempty"`

	// Identifier for the built-in intent to base this intent on.
	ParentIntentSignature string `json:"parentIntentSignature,omitempty" yaml:"parentIntentSignature,omitempty"`

	// Configuration block for a new list of slots and their priorities that are contained by the intent. This is ignored on create and only valid for updates. See `slot_priority`.
	SlotPriorities []types.Lex_V2modelsIntentSlotPriority `json:"slotPriorities,omitempty" yaml:"slotPriorities,omitempty"`

	// Description of the intent. Use the description to help identify the intent in lists.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Configuration block for invoking the alias Lambda function for each user input. You can invoke this Lambda function to personalize user interaction. See `dialog_code_hook`.
	DialogCodeHook types.Lex_V2modelsIntentDialogCodeHook `json:"dialogCodeHook,omitempty" yaml:"dialogCodeHook,omitempty"`

	// Configuration blocks for contexts that the intent activates when it is fulfilled. You can use an output context to indicate the intents that Amazon Lex should consider for the next turn of the conversation with a customer. When you use the outputContextsList property, all of the contexts specified in the list are activated when the intent is fulfilled. You can set up to 10 output contexts. You can also set the number of conversation turns that the context should be active, or the length of time that the context should be active. See `output_context`.
	OutputContexts []types.Lex_V2modelsIntentOutputContext `json:"outputContexts,omitempty" yaml:"outputContexts,omitempty"`

	// Configuration block for strings that a user might say to signal the intent. See `sample_utterance`.
	SampleUtterances []types.Lex_V2modelsIntentSampleUtterance `json:"sampleUtterances,omitempty" yaml:"sampleUtterances,omitempty"`

	// Version of the bot associated with this intent.
	BotVersion string `json:"botVersion,omitempty" yaml:"botVersion,omitempty"`

	// Configuration block for the response that Amazon Lex sends to the user when the intent is closed. See `closing_setting`.
	ClosingSetting types.Lex_V2modelsIntentClosingSetting `json:"closingSetting,omitempty" yaml:"closingSetting,omitempty"`

	/*
	   Name of the intent. Intent names must be unique in the locale that contains the intent and cannot match the name of any built-in intent.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Timeouts types.Lex_V2modelsIntentTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Identifier of the bot associated with this intent.
	BotId string `json:"botId,omitempty" yaml:"botId,omitempty"`

	//
	ConfirmationSetting types.Lex_V2modelsIntentConfirmationSetting `json:"confirmationSetting,omitempty" yaml:"confirmationSetting,omitempty"`

	// Configuration block for invoking the alias Lambda function when the intent is ready for fulfillment. You can invoke this function to complete the bot's transaction with the user. See `fulfillment_code_hook`.
	FulfillmentCodeHook types.Lex_V2modelsIntentFulfillmentCodeHook `json:"fulfillmentCodeHook,omitempty" yaml:"fulfillmentCodeHook,omitempty"`

	// Configuration block for the response that is sent to the user at the beginning of a conversation, before eliciting slot values. See `initial_response_setting`.
	InitialResponseSetting types.Lex_V2modelsIntentInitialResponseSetting `json:"initialResponseSetting,omitempty" yaml:"initialResponseSetting,omitempty"`

	// Configuration blocks for contexts that must be active for this intent to be considered by Amazon Lex. When an intent has an input context list, Amazon Lex only considers using the intent in an interaction with the user when the specified contexts are included in the active context list for the session. If the contexts are not active, then Amazon Lex will not use the intent. A context can be automatically activated using the outputContexts property or it can be set at runtime. See `input_context`.
	InputContexts []types.Lex_V2modelsIntentInputContext `json:"inputContexts,omitempty" yaml:"inputContexts,omitempty"`

	// Identifier of the language and locale where this intent is used. All of the bots, slot types, and slots used by the intent must have the same locale.
	LocaleId string `json:"localeId,omitempty" yaml:"localeId,omitempty"`
}
