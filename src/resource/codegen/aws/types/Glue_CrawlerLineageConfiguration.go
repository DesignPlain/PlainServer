package types

type Glue_CrawlerLineageConfiguration struct {
	// Specifies whether data lineage is enabled for the crawler. Valid values are: `ENABLE` and `DISABLE`. Default value is `DISABLE`.
	CrawlerLineageSettings string `json:"crawlerLineageSettings,omitempty" yaml:"crawlerLineageSettings,omitempty"`
}
