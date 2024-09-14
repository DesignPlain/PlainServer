package types

type Sagemaker_HumanTaskUIUiTemplate struct {
	// The content of the Liquid template for the worker user interface.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// The SHA-256 digest of the contents of the template.
	ContentSha256 string `json:"contentSha256,omitempty" yaml:"contentSha256,omitempty"`

	// The URL for the user interface template.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`
}
