package macie

type CustomDataIdentifier struct {
	// An array that lists specific character sequences (ignore words) to exclude from the results. If the text matched by the regular expression is the same as any string in this array, Amazon Macie ignores it. The array can contain as many as 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words are case sensitive.
	IgnoreWords []string `json:"ignoreWords,omitempty" yaml:"ignoreWords,omitempty"`

	// An array that lists specific character sequences (keywords), one of which must be within proximity (`maximum_match_distance`) of the regular expression to match. The array can contain as many as 50 keywords. Each keyword can contain 3 - 90 characters. Keywords aren't case sensitive.
	Keywords []string `json:"keywords,omitempty" yaml:"keywords,omitempty"`

	// The maximum number of characters that can exist between text that matches the regex pattern and the character sequences specified by the keywords array. Macie includes or excludes a result based on the proximity of a keyword to text that matches the regex pattern. The distance can be 1 - 300 characters. The default value is 50.
	MaximumMatchDistance int `json:"maximumMatchDistance,omitempty" yaml:"maximumMatchDistance,omitempty"`

	// A custom name for the custom data identifier. The name can contain as many as 128 characters. If omitted, the provider will assign a random, unique name. Conflicts with `name_prefix`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Creates a unique name beginning with the specified prefix. Conflicts with `name`.
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`

	// The regular expression (regex) that defines the pattern to match. The expression can contain as many as 512 characters.
	Regex string `json:"regex,omitempty" yaml:"regex,omitempty"`

	// A map of key-value pairs that specifies the tags to associate with the custom data identifier.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A custom description of the custom data identifier. The description can contain as many as 512 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}
