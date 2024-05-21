package types

type Tpu_V2VmNetworkEndpoint struct {
	/*
	   (Output)
	   The access config for the TPU worker.
	   Structure is documented below.
	*/
	AccessConfigs []Tpu_V2VmNetworkEndpointAccessConfig `json:"accessConfigs,omitempty" yaml:"accessConfigs,omitempty"`

	/*
	   (Output)
	   The internal IP address of this network endpoint.
	*/
	IpAddress string `json:"ipAddress,omitempty" yaml:"ipAddress,omitempty"`

	/*
	   (Output)
	   The port of this network endpoint.
	*/
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
