package types

type Appengine_StandardAppVersionLibrary struct {
	// Name of the library. Example "django".
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Version of the library to select, or "latest".
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
}
