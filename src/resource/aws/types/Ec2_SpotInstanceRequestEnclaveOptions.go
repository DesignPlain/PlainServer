package types

type Ec2_SpotInstanceRequestEnclaveOptions struct {
	/*
	   Whether Nitro Enclaves will be enabled on the instance. Defaults to `false`.

	   For more information, see the documentation on [Nitro Enclaves](https://docs.aws.amazon.com/enclaves/latest/user/nitro-enclave.html).
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
