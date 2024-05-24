package datasync

type LocationObjectStorage struct {
	// The access key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `access_key` and `secret_key` to provide the user name and password, respectively.
	AccessKey string `json:"accessKey,omitempty" yaml:"accessKey,omitempty"`

	// A list of DataSync Agent ARNs with which this location will be associated.
	AgentArns []string `json:"agentArns,omitempty" yaml:"agentArns,omitempty"`

	// The bucket on the self-managed object storage server that is used to read data from.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The secret key is used if credentials are required to access the self-managed object storage server. If your object storage requires a user name and password to authenticate, use `access_key` and `secret_key` to provide the user name and password, respectively.
	SecretKey string `json:"secretKey,omitempty" yaml:"secretKey,omitempty"`

	// The name of the self-managed object storage server. This value is the IP address or Domain Name Service (DNS) name of the object storage server. An agent uses this host name to mount the object storage server in a network.
	ServerHostname string `json:"serverHostname,omitempty" yaml:"serverHostname,omitempty"`

	// Key-value pairs of resource tags to assign to the DataSync Location. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Specifies a certificate to authenticate with an object storage system that uses a private or self-signed certificate authority (CA). You must specify a Base64-encoded .pem string. The certificate can be up to 32768 bytes (before Base64 encoding).
	ServerCertificate string `json:"serverCertificate,omitempty" yaml:"serverCertificate,omitempty"`

	// The port that your self-managed object storage server accepts inbound network traffic on. The server port is set by default to TCP 80 (`HTTP`) or TCP 443 (`HTTPS`). You can specify a custom port if your self-managed object storage server requires one.
	ServerPort int `json:"serverPort,omitempty" yaml:"serverPort,omitempty"`

	// The protocol that the object storage server uses to communicate. Valid values are `HTTP` or `HTTPS`.
	ServerProtocol string `json:"serverProtocol,omitempty" yaml:"serverProtocol,omitempty"`

	// A subdirectory in the HDFS cluster. This subdirectory is used to read data from or write data to the HDFS cluster. If the subdirectory isn't specified, it will default to /.
	Subdirectory string `json:"subdirectory,omitempty" yaml:"subdirectory,omitempty"`
}
