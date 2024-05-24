package types

type Glue_TriggerPredicateCondition struct {
	// The condition crawl state. Currently, the values supported are `RUNNING`, `SUCCEEDED`, `CANCELLED`, and `FAILED`. If this is specified, `crawler_name` must also be specified. Conflicts with `state`.
	CrawlState string `json:"crawlState,omitempty" yaml:"crawlState,omitempty"`

	// The name of the crawler to watch. If this is specified, `crawl_state` must also be specified. Conflicts with `job_name`.
	CrawlerName string `json:"crawlerName,omitempty" yaml:"crawlerName,omitempty"`

	// The name of the job to watch. If this is specified, `state` must also be specified. Conflicts with `crawler_name`.
	JobName string `json:"jobName,omitempty" yaml:"jobName,omitempty"`

	// A logical operator. Defaults to `EQUALS`.
	LogicalOperator string `json:"logicalOperator,omitempty" yaml:"logicalOperator,omitempty"`

	// The condition job state. Currently, the values supported are `SUCCEEDED`, `STOPPED`, `TIMEOUT` and `FAILED`. If this is specified, `job_name` must also be specified. Conflicts with `crawler_state`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
