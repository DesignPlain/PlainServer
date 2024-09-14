package networkservices

type Mesh struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Optional. If set to a valid TCP port (1-65535), instructs the SIDECAR proxy to listen on the
	   specified port of localhost (127.0.0.1) address. The SIDECAR proxy will expect all traffic to
	   be redirected to this port regardless of its actual ip:port destination. If unset, a port
	   '15001' is used as the interception port. This will is applicable only for sidecar proxy
	   deployments.
	*/
	InterceptionPort int `json:"interceptionPort,omitempty" yaml:"interceptionPort,omitempty"`

	/*
	   Set of label tags associated with the Mesh resource.
	   --Note--: This field is non-authoritative, and will only manage the labels present in your configuration.
	   Please refer to the field `effective_labels` for all of the labels present on the resource.
	*/
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	/*
	   Short name of the Mesh resource to be created.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
