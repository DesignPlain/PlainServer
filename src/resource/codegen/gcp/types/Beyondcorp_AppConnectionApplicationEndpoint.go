package types

type Beyondcorp_AppConnectionApplicationEndpoint struct {
	// Hostname or IP address of the remote application endpoint.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	/*
	   Port of the remote application endpoint.

	   - - -
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
