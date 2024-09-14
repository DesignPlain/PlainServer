package types

type Glue_CrawlerJdbcTarget struct {
	// The name of the connection to use to connect to the JDBC target.
	ConnectionName string `json:"connectionName,omitempty" yaml:"connectionName,omitempty"`

	// Specify a value of `RAWTYPES` or `COMMENTS` to enable additional metadata intable responses. `RAWTYPES` provides the native-level datatype. `COMMENTS` provides comments associated with a column or table in the database.
	EnableAdditionalMetadatas []string `json:"enableAdditionalMetadatas,omitempty" yaml:"enableAdditionalMetadatas,omitempty"`

	// A list of glob patterns used to exclude from the crawl.
	Exclusions []string `json:"exclusions,omitempty" yaml:"exclusions,omitempty"`

	// The path of the JDBC target.
	Path string `json:"path,omitempty" yaml:"path,omitempty"`
}
