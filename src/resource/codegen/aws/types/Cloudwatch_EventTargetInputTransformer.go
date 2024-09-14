package types

type Cloudwatch_EventTargetInputTransformer struct {
	/*
	   Key value pairs specified in the form of JSONPath (for example, time = $.time)
	   - You can have as many as 100 key-value pairs.
	   - You must use JSON dot notation, not bracket notation.
	   - The keys can't start with "AWS".
	*/
	InputPaths map[string]string `json:"inputPaths,omitempty" yaml:"inputPaths,omitempty"`

	// Template to customize data sent to the target. Must be valid JSON. To send a string value, the string value must include double quotes.
	InputTemplate string `json:"inputTemplate,omitempty" yaml:"inputTemplate,omitempty"`
}
