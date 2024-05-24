package codebuild

type SourceCredential struct {
	// The type of authentication used to connect to a GitHub, GitHub Enterprise, or Bitbucket repository. An OAUTH connection is not supported by the API.
	AuthType string `json:"authType,omitempty" yaml:"authType,omitempty"`

	// The source provider used for this project.
	ServerType string `json:"serverType,omitempty" yaml:"serverType,omitempty"`

	// For `GitHub` or `GitHub Enterprise`, this is the personal access token. For `Bitbucket`, this is the app password.
	Token string `json:"token,omitempty" yaml:"token,omitempty"`

	// The Bitbucket username when the authType is `BASIC_AUTH`. This parameter is not valid for other types of source providers or connections.
	UserName string `json:"userName,omitempty" yaml:"userName,omitempty"`
}
