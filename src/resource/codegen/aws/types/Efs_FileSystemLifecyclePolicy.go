package types

type Efs_FileSystemLifecyclePolicy struct {
	// Indicates how long it takes to transition files to the archive storage class. Requires transition_to_ia, Elastic Throughput and General Purpose performance mode. Valid values: `AFTER_1_DAY`, `AFTER_7_DAYS`, `AFTER_14_DAYS`, `AFTER_30_DAYS`, `AFTER_60_DAYS`, `AFTER_90_DAYS`, `AFTER_180_DAYS`, `AFTER_270_DAYS`, or `AFTER_365_DAYS`.
	TransitionToArchive string `json:"transitionToArchive,omitempty" yaml:"transitionToArchive,omitempty"`

	// Indicates how long it takes to transition files to the IA storage class. Valid values: `AFTER_1_DAY`, `AFTER_7_DAYS`, `AFTER_14_DAYS`, `AFTER_30_DAYS`, `AFTER_60_DAYS`, `AFTER_90_DAYS`, `AFTER_180_DAYS`, `AFTER_270_DAYS`, or `AFTER_365_DAYS`.
	TransitionToIa string `json:"transitionToIa,omitempty" yaml:"transitionToIa,omitempty"`

	// Describes the policy used to transition a file from infequent access storage to primary storage. Valid values: `AFTER_1_ACCESS`.
	TransitionToPrimaryStorageClass string `json:"transitionToPrimaryStorageClass,omitempty" yaml:"transitionToPrimaryStorageClass,omitempty"`
}
