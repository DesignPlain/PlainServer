package types

type Appstream_StackStreamingExperienceSettings struct {
	/*
	   The preferred protocol that you want to use while streaming your application.
	   Valid values are `TCP` and `UDP`.
	*/
	PreferredProtocol string `json:"preferredProtocol,omitempty" yaml:"preferredProtocol,omitempty"`
}
