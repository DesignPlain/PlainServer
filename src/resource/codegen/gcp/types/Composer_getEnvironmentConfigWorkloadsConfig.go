package types

type Composer_getEnvironmentConfigWorkloadsConfig struct {
	// Configuration for resources used by Airflow web server.
	WebServers []Composer_getEnvironmentConfigWorkloadsConfigWebServer `json:"webServers,omitempty" yaml:"webServers,omitempty"`

	// Configuration for resources used by Airflow workers.
	Workers []Composer_getEnvironmentConfigWorkloadsConfigWorker `json:"workers,omitempty" yaml:"workers,omitempty"`

	// Configuration for resources used by DAG processor.
	DagProcessors []Composer_getEnvironmentConfigWorkloadsConfigDagProcessor `json:"dagProcessors,omitempty" yaml:"dagProcessors,omitempty"`

	// Configuration for resources used by Airflow schedulers.
	Schedulers []Composer_getEnvironmentConfigWorkloadsConfigScheduler `json:"schedulers,omitempty" yaml:"schedulers,omitempty"`

	// Configuration for resources used by Airflow triggerers.
	Triggerers []Composer_getEnvironmentConfigWorkloadsConfigTriggerer `json:"triggerers,omitempty" yaml:"triggerers,omitempty"`
}
