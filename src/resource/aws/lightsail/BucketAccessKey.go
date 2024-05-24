package lightsail

type BucketAccessKey struct {
	// The name of the bucket that the new access key will belong to, and grant access to.
	BucketName string `json:"bucketName,omitempty" yaml:"bucketName,omitempty"`
}
