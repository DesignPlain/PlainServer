package lex

import types "DesignSphere_Server/src/resource/aws/types"

type Intent struct {
	/*
	   When the user answers "no" to the question defined in
	   `confirmation_prompt`, Amazon Lex responds with this statement to acknowledge that the intent was
	   canceled. You must provide both the `rejection_statement` and the `confirmation_prompt`, or neither.
	   Attributes are documented under statement.
	*/
	RejectionStatement types.Lex_IntentRejectionStatement `json:"rejectionStatement,omitempty" yaml:"rejectionStatement,omitempty"`

	/*
	   An array of utterances (strings) that a user might say to signal
	   the intent. For example, "I want {PizzaSize} pizza", "Order {Quantity} {PizzaSize} pizzas".
	   In each utterance, a slot name is enclosed in curly braces. Must have between 1 and 10 items in the list, and each item must be less than or equal to 200 characters in length.
	*/
	SampleUtterances []string `json:"sampleUtterances,omitempty" yaml:"sampleUtterances,omitempty"`

	/*
	   Determines if a new slot type version is created when the initial
	   resource is created and on each update. Defaults to `false`.
	*/
	CreateVersion bool `json:"createVersion,omitempty" yaml:"createVersion,omitempty"`

	/*
	   Specifies a Lambda function to invoke for each user input. You can
	   invoke this Lambda function to personalize user interaction. Attributes are documented under code_hook.
	*/
	DialogCodeHook types.Lex_IntentDialogCodeHook `json:"dialogCodeHook,omitempty" yaml:"dialogCodeHook,omitempty"`

	/*
	   Describes how the intent is fulfilled. For example, after a
	   user provides all of the information for a pizza order, `fulfillment_activity` defines how the bot
	   places an order with a local pizza store. Attributes are documented under fulfillment_activity.
	*/
	FulfillmentActivity types.Lex_IntentFulfillmentActivity `json:"fulfillmentActivity,omitempty" yaml:"fulfillmentActivity,omitempty"`

	// The name of the intent, not case sensitive. Must be less than or equal to 100 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   A unique identifier for the built-in intent to base this
	   intent on. To find the signature for an intent, see
	   [Standard Built-in Intents](https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/standard-intents)
	   in the Alexa Skills Kit.
	*/
	ParentIntentSignature string `json:"parentIntentSignature,omitempty" yaml:"parentIntentSignature,omitempty"`

	/*
	   An list of intent slots. At runtime, Amazon Lex elicits required slot values
	   from the user using prompts defined in the slots. Attributes are documented under slot.
	*/
	Slots []types.Lex_IntentSlot `json:"slots,omitempty" yaml:"slots,omitempty"`

	/*
	   The statement that you want Amazon Lex to convey to the user
	   after the intent is successfully fulfilled by the Lambda function. This element is relevant only if
	   you provide a Lambda function in the `fulfillment_activity`. If you return the intent to the client
	   application, you can't specify this element. The `follow_up_prompt` and `conclusion_statement` are
	   mutually exclusive. You can specify only one. Attributes are documented under statement.
	*/
	ConclusionStatement types.Lex_IntentConclusionStatement `json:"conclusionStatement,omitempty" yaml:"conclusionStatement,omitempty"`

	/*
	   Prompts the user to confirm the intent. This question should
	   have a yes or no answer. You you must provide both the `rejection_statement` and `confirmation_prompt`,
	   or neither. Attributes are documented under prompt.
	*/
	ConfirmationPrompt types.Lex_IntentConfirmationPrompt `json:"confirmationPrompt,omitempty" yaml:"confirmationPrompt,omitempty"`

	// A description of the intent. Must be less than or equal to 200 characters in length.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Amazon Lex uses this prompt to solicit additional activity after
	   fulfilling an intent. For example, after the OrderPizza intent is fulfilled, you might prompt the
	   user to order a drink. The `follow_up_prompt` field and the `conclusion_statement` field are mutually
	   exclusive. You can specify only one. Attributes are documented under follow_up_prompt.
	*/
	FollowUpPrompt types.Lex_IntentFollowUpPrompt `json:"followUpPrompt,omitempty" yaml:"followUpPrompt,omitempty"`
}
