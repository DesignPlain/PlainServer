package types

type Elasticsearch_getDomainSnapshotOption struct {
	// Hour during which the service takes an automated daily snapshot of the indices in the domain.
	AutomatedSnapshotStartHour int `json:"automatedSnapshotStartHour,omitempty" yaml:"automatedSnapshotStartHour,omitempty"`
}
