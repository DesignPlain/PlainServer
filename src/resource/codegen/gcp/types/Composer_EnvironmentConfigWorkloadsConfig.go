package types

type Composer_EnvironmentConfigWorkloadsConfig struct {
	// Configuration for resources used by DAG processor.
	DagProcessor Composer_EnvironmentConfigWorkloadsConfigDagProcessor `json:"dagProcessor,omitempty" yaml:"dagProcessor,omitempty"`

	// Configuration for resources used by Airflow schedulers.
	Scheduler Composer_EnvironmentConfigWorkloadsConfigScheduler `json:"scheduler,omitempty" yaml:"scheduler,omitempty"`

	// Configuration for resources used by Airflow triggerers.
	Triggerer Composer_EnvironmentConfigWorkloadsConfigTriggerer `json:"triggerer,omitempty" yaml:"triggerer,omitempty"`

	// Configuration for resources used by Airflow web server.
	WebServer Composer_EnvironmentConfigWorkloadsConfigWebServer `json:"webServer,omitempty" yaml:"webServer,omitempty"`

	// Configuration for resources used by Airflow workers.
	Worker Composer_EnvironmentConfigWorkloadsConfigWorker `json:"worker,omitempty" yaml:"worker,omitempty"`
}
