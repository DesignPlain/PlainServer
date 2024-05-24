package connect

type ContactFlow struct {
	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow source specified with `filename`.
	ContentHash string `json:"contentHash,omitempty" yaml:"contentHash,omitempty"`

	// Specifies the description of the Contact Flow.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The path to the Contact Flow source within the local filesystem. Conflicts with `content`.
	Filename string `json:"filename,omitempty" yaml:"filename,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Specifies the name of the Contact Flow.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Tags to apply to the Contact Flow. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies the type of the Contact Flow. Defaults to `CONTACT_FLOW`. Allowed Values are: `CONTACT_FLOW`, `CUSTOMER_QUEUE`, `CUSTOMER_HOLD`, `CUSTOMER_WHISPER`, `AGENT_HOLD`, `AGENT_WHISPER`, `OUTBOUND_WHISPER`, `AGENT_TRANSFER`, `QUEUE_TRANSFER`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Specifies the content of the Contact Flow, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the `filename` argument cannot be used.
	Content string `json:"content,omitempty" yaml:"content,omitempty"`
}
