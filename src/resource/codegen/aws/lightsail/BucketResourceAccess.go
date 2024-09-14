package lightsail

type BucketResourceAccess struct {
	// The name of the bucket to grant access to.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`

	// The name of the resource to be granted bucket access.
	ResourceName string `json:"resourceName,omitempty" yaml:"resourceName,omitempty"`
}
