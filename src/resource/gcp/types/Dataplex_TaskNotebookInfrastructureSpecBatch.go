package types

type Dataplex_TaskNotebookInfrastructureSpecBatch struct {
	// Max configurable executors. If maxExecutorsCount > executorsCount, then auto-scaling is enabled. Max Executor Count should be between 2 and 1000. [Default=1000]
	MaxExecutorsCount int `json:"maxExecutorsCount,omitempty" yaml:"maxExecutorsCount,omitempty"`

	// Total number of job executors. Executor Count should be between 2 and 100. [Default=2]
	ExecutorsCount int `json:"executorsCount,omitempty" yaml:"executorsCount,omitempty"`
}
