package types

type Storage_getBucketLifecycleRuleCondition struct {
	// Storage Class of objects to satisfy this condition. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE, STANDARD, DURABLE_REDUCED_AVAILABILITY.
	MatchesStorageClasses []string `json:"matchesStorageClasses,omitempty" yaml:"matchesStorageClasses,omitempty"`

	// One or more matching name suffixes to satisfy this condition.
	MatchesSuffixes []string `json:"matchesSuffixes,omitempty" yaml:"matchesSuffixes,omitempty"`

	// Relevant only for versioned objects. The number of newer versions of an object to satisfy this condition.
	NumNewerVersions int `json:"numNewerVersions,omitempty" yaml:"numNewerVersions,omitempty"`

	// Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: "LIVE", "ARCHIVED", "ANY".
	WithState string `json:"withState,omitempty" yaml:"withState,omitempty"`

	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	CustomTimeBefore string `json:"customTimeBefore,omitempty" yaml:"customTimeBefore,omitempty"`

	// Number of days elapsed since the user-specified timestamp set on an object.
	DaysSinceCustomTime int `json:"daysSinceCustomTime,omitempty" yaml:"daysSinceCustomTime,omitempty"`

	/*
	   Number of days elapsed since the noncurrent timestamp of an object. This
	   										condition is relevant only for versioned objects.
	*/
	DaysSinceNoncurrentTime int `json:"daysSinceNoncurrentTime,omitempty" yaml:"daysSinceNoncurrentTime,omitempty"`

	// One or more matching name prefixes to satisfy this condition.
	MatchesPrefixes []string `json:"matchesPrefixes,omitempty" yaml:"matchesPrefixes,omitempty"`

	// Minimum age of an object in days to satisfy this condition.
	Age int `json:"age,omitempty" yaml:"age,omitempty"`

	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	CreatedBefore string `json:"createdBefore,omitempty" yaml:"createdBefore,omitempty"`

	// While set true, age value will be omitted.Required to set true when age is unset in the config file.
	NoAge bool `json:"noAge,omitempty" yaml:"noAge,omitempty"`

	// Creation date of an object in RFC 3339 (e.g. 2017-06-13) to satisfy this condition.
	NoncurrentTimeBefore string `json:"noncurrentTimeBefore,omitempty" yaml:"noncurrentTimeBefore,omitempty"`
}
