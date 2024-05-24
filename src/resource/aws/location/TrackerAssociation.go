package location

type TrackerAssociation struct {
	// The Amazon Resource Name (ARN) for the geofence collection to be associated to tracker resource. Used when you need to specify a resource across all AWS.
	ConsumerArn string `json:"consumerArn,omitempty" yaml:"consumerArn,omitempty"`

	// The name of the tracker resource to be associated with a geofence collection.
	TrackerName string `json:"trackerName,omitempty" yaml:"trackerName,omitempty"`
}
