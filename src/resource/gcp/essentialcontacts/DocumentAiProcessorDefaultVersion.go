package essentialcontacts

type DocumentAiProcessorDefaultVersion struct {
	/*
	   The processor to set the version on.


	   - - -
	*/
	Processor string `json:"processor,omitempty" yaml:"processor,omitempty"`

	/*
	   The version to set. Using `stable` or `rc` will cause the API to return the latest version in that release channel.
	   Apply `lifecycle.ignore_changes` to the `version` field to suppress this diff.
	*/
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
