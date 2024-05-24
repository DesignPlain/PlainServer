package connect

type ContactFlowModule struct {
	// Specifies the content of the Contact Flow Module, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the `filename` argument cannot be used.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow Module source specified with `filename`.
	ContentHash string `json:"contentHash,omitempty" yaml:"contentHash,omitempty"`

	// Specifies the description of the Contact Flow Module.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The path to the Contact Flow Module source within the local filesystem. Conflicts with `content`.
	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Specifies the name of the Contact Flow Module.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Tags to apply to the Contact Flow Module. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
