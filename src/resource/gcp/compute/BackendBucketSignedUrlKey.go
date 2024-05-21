package compute

type BackendBucketSignedUrlKey struct {
	/*
	   The backend bucket this signed URL key belongs.


	   - - -
	*/
	BackendBucket string `json:"backendBucket,omitempty" yaml:"backendBucket,omitempty"`

	/*
	   128-bit key value used for signing the URL. The key value must be a
	   valid RFC 4648 Section 5 base64url encoded string.
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	KeyValue string `json:"keyValue,omitempty" yaml:"keyValue,omitempty"`

	// Name of the signed URL key.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
