package types

type Glue_CrawlerHudiTarget struct {
	// The name of the connection to use to connect to the Hudi target.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []string `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`

	// The maximum depth of Amazon S3 paths that the crawler can traverse to discover the Hudi metadata folder in your Amazon S3 path. Used to limit the crawler run time. Valid values are between `1` and `20`.
	MaximumTraversalDepth int `json:"maximumTraversalDepth,omitempty" yaml:"maximumTraversalDepth,omitempty"`

	// One or more Amazon S3 paths that contains Hudi metadata folders as s3://bucket/prefix.
	Paths []string `json:"paths,omitempty" yaml:"paths,omitempty"`
}
