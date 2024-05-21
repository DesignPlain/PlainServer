package types

type Securesourcemanager_InstanceHostConfig struct {
	/*
	   (Output)
	   API hostname.
	*/
	Api string `json:"api,omitempty" yaml:"api,omitempty"`

	/*
	   (Output)
	   Git HTTP hostname.
	*/
	GitHttp string `json:"gitHttp,omitempty" yaml:"gitHttp,omitempty"`

	/*
	   (Output)
	   Git SSH hostname.
	*/
	GitSsh string `json:"gitSsh,omitempty" yaml:"gitSsh,omitempty"`

	/*
	   (Output)
	   HTML hostname.
	*/
	Html string `json:"html,omitempty" yaml:"html,omitempty"`
}
