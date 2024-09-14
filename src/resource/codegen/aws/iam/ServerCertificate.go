package iam

type ServerCertificate struct {
	/*
	   The IAM path for the server certificate.  If it is not
	   included, it defaults to a slash (/). If this certificate is for use with
	   AWS CloudFront, the path must be in format `/cloudfront/your_path_here`.
	   See [IAM Identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html) for more details on IAM Paths.
	*/
	Path string `json:"path,omitempty" yaml:"path,omitempty"`

	// The contents of the private key in PEM-encoded format.
	PrivateKey string `json:"privateKey,omitempty" yaml:"privateKey,omitempty"`

	/*
	   Map of resource tags for the server certificate. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.

	   > --NOTE:-- AWS performs behind-the-scenes modifications to some certificate files if they do not adhere to a specific format. These modifications will result in this provider forever believing that it needs to update the resources since the local and AWS file contents will not match after theses modifications occur. In order to prevent this from happening you must ensure that all your PEM-encoded files use UNIX line-breaks and that `certificate_body` contains only one certificate. All other certificates should go in `certificate_chain`. It is common for some Certificate Authorities to issue certificate files that have DOS line-breaks and that are actually multiple certificates concatenated together in order to form a full certificate chain.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   The contents of the public key certificate in
	   PEM-encoded format.
	*/
	CertificateBody string `json:"certificateBody,omitempty" yaml:"certificateBody,omitempty"`

	/*
	   The contents of the certificate chain.
	   This is typically a concatenation of the PEM-encoded public key certificates
	   of the chain.
	*/
	CertificateChain string `json:"certificateChain,omitempty" yaml:"certificateChain,omitempty"`

	/*
	   The name of the Server Certificate. Do not include the
	   path in this value. If omitted, the provider will assign a random, unique name.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   Creates a unique name beginning with the specified
	   prefix. Conflicts with `name`.
	*/
	NamePrefix string `json:"namePrefix,omitempty" yaml:"namePrefix,omitempty"`
}
