package compute

import types "DesignSphere_Server/src/resource/gcp/types"

type BackendBucket struct {
	// The security policy associated with this backend bucket.
	EdgeSecurityPolicy string `json:"edgeSecurityPolicy,omitempty" yaml:"edgeSecurityPolicy,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Cloud CDN configuration for this Backend Bucket.
	   Structure is documented below.
	*/
	CdnPolicy types.Compute_BackendBucketCdnPolicy `json:"cdnPolicy,omitempty" yaml:"cdnPolicy,omitempty"`

	/*
	   Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
	   Possible values are: `AUTOMATIC`, `DISABLED`.
	*/
	CompressionMode string `json:"compressionMode,omitempty" yaml:"compressionMode,omitempty"`

	/*
	   An optional textual description of the resource; provided by the
	   client when the resource is created.
	*/
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   Name of the resource. Provided by the client when the resource is
	   created. The name must be 1-63 characters long, and comply with
	   RFC1035.  Specifically, the name must be 1-63 characters long and
	   match the regular expression `a-z?` which means
	   the first character must be a lowercase letter, and all following
	   characters must be a dash, lowercase letter, or digit, except the
	   last character, which cannot be a dash.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Cloud Storage bucket name.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// Headers that the HTTP/S load balancer should add to proxied responses.
	CustomResponseHeaders []string `json:"customResponseHeaders,omitempty" yaml:"customResponseHeaders,omitempty"`

	// If true, enable Cloud CDN for this BackendBucket.
	EnableCdn bool `json:"enableCdn,omitempty" yaml:"enableCdn,omitempty"`
}
