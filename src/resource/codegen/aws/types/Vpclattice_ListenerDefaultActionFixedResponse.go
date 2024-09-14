package types

type Vpclattice_ListenerDefaultActionFixedResponse struct {
	// Custom HTTP status code to return, e.g. a 404 response code. See [Listeners](https://docs.aws.amazon.com/vpc-lattice/latest/ug/listeners.html) in the AWS documentation for a list of supported codes.
	StatusCode int `json:"statusCode,omitempty" yaml:"statusCode,omitempty"`
}
