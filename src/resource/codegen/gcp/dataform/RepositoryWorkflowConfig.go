package dataform

import types "libds/gcp/types"

type RepositoryWorkflowConfig struct {
	// The workflow's name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A reference to the region
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The name of the release config whose releaseCompilationResult should be executed. Must be in the format projects/-/locations/-/repositories/-/releaseConfigs/-.


	   - - -
	*/
	ReleaseConfig string `json:"releaseConfig,omitempty" yaml:"releaseConfig,omitempty"`

	// A reference to the Dataform repository
	Repository string `json:"repository,omitempty" yaml:"repository,omitempty"`

	// Optional. Specifies the time zone to be used when interpreting cronSchedule. Must be a time zone name from the time zone database (https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). If left unspecified, the default is UTC.
	TimeZone string `json:"timeZone,omitempty" yaml:"timeZone,omitempty"`

	// Optional. Optional schedule (in cron format) for automatic creation of compilation results.
	CronSchedule string `json:"cronSchedule,omitempty" yaml:"cronSchedule,omitempty"`

	/*
	   Optional. If left unset, a default InvocationConfig will be used.
	   Structure is documented below.
	*/
	InvocationConfig types.Dataform_RepositoryWorkflowConfigInvocationConfig `json:"invocationConfig,omitempty" yaml:"invocationConfig,omitempty"`
}
