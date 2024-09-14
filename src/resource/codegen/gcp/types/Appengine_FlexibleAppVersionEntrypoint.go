package types

type Appengine_FlexibleAppVersionEntrypoint struct {
	// The format should be a shell command that can be fed to bash -c.
	Shell string `json:"shell,omitempty" yaml:"shell,omitempty"`
}
