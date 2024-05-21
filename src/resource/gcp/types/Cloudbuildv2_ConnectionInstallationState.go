package types

type Cloudbuildv2_ConnectionInstallationState struct {
	/*
	   (Output)
	   Output only. Link to follow for next action. Empty string if the installation is already complete.
	*/
	ActionUri string `json:"actionUri,omitempty" yaml:"actionUri,omitempty"`

	/*
	   (Output)
	   Output only. Message of what the user should do next to continue the installation. Empty string if the installation is already complete.
	*/
	Message string `json:"message,omitempty" yaml:"message,omitempty"`

	/*
	   (Output)
	   Output only. Current step of the installation process.
	*/
	Stage string `json:"stage,omitempty" yaml:"stage,omitempty"`
}
