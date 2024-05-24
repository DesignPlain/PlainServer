package codepipeline

import types "DesignSphere_Server/src/resource/aws/types"

type Webhook struct {
	// The name of the pipeline.
	TargetPipeline string `json:"targetPipeline,omitempty" yaml:"targetPipeline,omitempty"`

	// The type of authentication  to use. One of `IP`, `GITHUB_HMAC`, or `UNAUTHENTICATED`.
	Authentication string `json:"authentication,omitempty" yaml:"authentication,omitempty"`

	// An `auth` block. Required for `IP` and `GITHUB_HMAC`. Auth blocks are documented below.
	AuthenticationConfiguration types.Codepipeline_WebhookAuthenticationConfiguration `json:"authenticationConfiguration,omitempty" yaml:"authenticationConfiguration,omitempty"`

	// One or more `filter` blocks. Filter blocks are documented below.
	Filters []types.Codepipeline_WebhookFilter `json:"filters,omitempty" yaml:"filters,omitempty"`

	// The name of the webhook.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The name of the action in a pipeline you want to connect to the webhook. The action must be from the source (first) stage of the pipeline.
	TargetAction string `json:"targetAction,omitempty" yaml:"targetAction,omitempty"`
}
