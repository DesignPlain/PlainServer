package types

type Cloudbuild_getTriggerBuild struct {
	// Special options for this build.
	Options []Cloudbuild_getTriggerBuildOption `json:"options,omitempty" yaml:"options,omitempty"`

	/*
	   The location of the source files to build.

	   One of 'storageSource' or 'repoSource' must be provided.
	*/
	Sources []Cloudbuild_getTriggerBuildSource `json:"sources,omitempty" yaml:"sources,omitempty"`

	// Substitutions data for Build resource.
	Substitutions map[string]string `json:"substitutions,omitempty" yaml:"substitutions,omitempty"`

	/*
	   A list of images to be pushed upon the successful completion of all build steps.
	   The images are pushed using the builder service account's credentials.
	   The digests of the pushed images will be stored in the Build resource's results field.
	   If any of the images fail to be pushed, the build status is marked FAILURE.
	*/
	Images []string `json:"images,omitempty" yaml:"images,omitempty"`

	// Secrets and secret environment variables.
	AvailableSecrets []Cloudbuild_getTriggerBuildAvailableSecret `json:"availableSecrets,omitempty" yaml:"availableSecrets,omitempty"`

	/*
	   Google Cloud Storage bucket where logs should be written.
	   Logs file names will be of the format ${logsBucket}/log-${build_id}.txt.
	*/
	LogsBucket string `json:"logsBucket,omitempty" yaml:"logsBucket,omitempty"`

	/*
	   TTL in queue for this build. If provided and the build is enqueued longer than this value,
	   the build will expire and the build status will be EXPIRED.
	   The TTL starts ticking from createTime.
	   A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
	*/
	QueueTtl string `json:"queueTtl,omitempty" yaml:"queueTtl,omitempty"`

	// Secrets to decrypt using Cloud Key Management Service.
	Secrets []Cloudbuild_getTriggerBuildSecret `json:"secrets,omitempty" yaml:"secrets,omitempty"`

	// The operations to be performed on the workspace.
	Steps []Cloudbuild_getTriggerBuildStep `json:"steps,omitempty" yaml:"steps,omitempty"`

	// Tags for annotation of a Build. These are not docker tags.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   Amount of time that this build should be allowed to run, to second granularity.
	   If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
	   This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
	   The expected format is the number of seconds followed by s.
	   Default time is ten minutes (600s).
	*/
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	// Artifacts produced by the build that should be uploaded upon successful completion of all build steps.
	Artifacts []Cloudbuild_getTriggerBuildArtifact `json:"artifacts,omitempty" yaml:"artifacts,omitempty"`
}
