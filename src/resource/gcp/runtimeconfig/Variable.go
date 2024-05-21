package runtimeconfig

type Variable struct {
	/*
	   The ID of the project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   or `value` - (Required) The content to associate with the variable.
	   Exactly one of `text` or `variable` must be specified. If `text` is specified,
	   it must be a valid UTF-8 string and less than 4096 bytes in length. If `value`
	   is specified, it must be base64 encoded and less than 4096 bytes in length.

	   - - -
	*/
	Text string `json:"text,omitempty" yaml:"text,omitempty"`

	//
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	/*
	   The name of the variable to manage. Note that variable
	   names can be hierarchical using slashes (e.g. "prod-variables/hostname").
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The name of the RuntimeConfig resource containing this
	   variable.
	*/
	Parent string `json:"parent,omitempty" yaml:"parent,omitempty"`
}
