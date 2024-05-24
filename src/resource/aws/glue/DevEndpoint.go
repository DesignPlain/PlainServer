package glue

type DevEndpoint struct {
	// Security group IDs for the security groups to be used by this endpoint.
	SecurityGroupIds []string `json:"securityGroupIds,omitempty" yaml:"securityGroupIds,omitempty"`

	// The subnet ID for the new endpoint to use.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// Key-value map of resource tags. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Path to one or more Java Jars in an S3 bucket that should be loaded in this endpoint.
	ExtraJarsS3Path string `json:"extraJarsS3Path,omitempty" yaml:"extraJarsS3Path,omitempty"`

	// Specifies the versions of Python and Apache Spark to use. Defaults to AWS Glue version 0.9.
	GlueVersion string `json:"glueVersion,omitempty" yaml:"glueVersion,omitempty"`

	// The number of workers of a defined worker type that are allocated to this endpoint. This field is available only when you choose worker type G.1X or G.2X.
	NumberOfWorkers int `json:"numberOfWorkers,omitempty" yaml:"numberOfWorkers,omitempty"`

	// The name of this endpoint. It must be unique in your account.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The type of predefined worker that is allocated to this endpoint. Accepts a value of Standard, G.1X, or G.2X.
	WorkerType string `json:"workerType,omitempty" yaml:"workerType,omitempty"`

	// The name of the Security Configuration structure to be used with this endpoint.
	SecurityConfiguration string `json:"securityConfiguration,omitempty" yaml:"securityConfiguration,omitempty"`

	// Path(s) to one or more Python libraries in an S3 bucket that should be loaded in this endpoint. Multiple values must be complete paths separated by a comma.
	ExtraPythonLibsS3Path string `json:"extraPythonLibsS3Path,omitempty" yaml:"extraPythonLibsS3Path,omitempty"`

	// The number of AWS Glue Data Processing Units (DPUs) to allocate to this endpoint. Conflicts with `worker_type`.
	NumberOfNodes int `json:"numberOfNodes,omitempty" yaml:"numberOfNodes,omitempty"`

	// The IAM role for this endpoint.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of arguments used to configure the endpoint.
	Arguments map[string]string `json:"arguments,omitempty" yaml:"arguments,omitempty"`

	// The public key to be used by this endpoint for authentication.
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`

	// A list of public keys to be used by this endpoint for authentication.
	PublicKeys []string `json:"publicKeys,omitempty" yaml:"publicKeys,omitempty"`
}
