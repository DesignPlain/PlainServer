package types

type Tpu_NodeNetworkEndpoint struct {
	/*
	   (Output)
	   The IP address of this network endpoint.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   (Output)
	   The port of this network endpoint.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
