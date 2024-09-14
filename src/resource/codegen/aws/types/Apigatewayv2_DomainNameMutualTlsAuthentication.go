package types

type Apigatewayv2_DomainNameMutualTlsAuthentication struct {
	// Amazon S3 URL that specifies the truststore for mutual TLS authentication, for example, `s3://bucket-name/key-name`. The truststore can contain certificates from public or private certificate authorities. To update the truststore, upload a new version to S3, and then update your custom domain name to use the new version.
	TruststoreUri string `json:"truststoreUri,omitempty" yaml:"truststoreUri,omitempty"`

	// Version of the S3 object that contains the truststore. To specify a version, you must have versioning enabled for the S3 bucket.
	TruststoreVersion string `json:"truststoreVersion,omitempty" yaml:"truststoreVersion,omitempty"`
}
