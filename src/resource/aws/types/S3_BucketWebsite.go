package types

type S3_BucketWebsite struct {
	// An absolute path to the document to return in case of a 4XX error.
	ErrorDocument string `json:"errorDocument,omitempty" yaml:"errorDocument,omitempty"`

	// Amazon S3 returns this index document when requests are made to the root domain or any of the subfolders.
	IndexDocument string `json:"indexDocument,omitempty" yaml:"indexDocument,omitempty"`

	// A hostname to redirect all website requests for this bucket to. Hostname can optionally be prefixed with a protocol (`http://` or `https://`) to use when redirecting requests. The default is the protocol that is used in the original request.
	RedirectAllRequestsTo string `json:"redirectAllRequestsTo,omitempty" yaml:"redirectAllRequestsTo,omitempty"`

	/*
	   A json array containing [routing rules](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html)
	   describing redirect behavior and when redirects are applied.

	   The `CORS` object supports the following:
	*/
	RoutingRules string `json:"routingRules,omitempty" yaml:"routingRules,omitempty"`
}
