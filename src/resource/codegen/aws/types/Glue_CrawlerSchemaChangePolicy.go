package types

type Glue_CrawlerSchemaChangePolicy struct {
	// The deletion behavior when the crawler finds a deleted object. Valid values: `LOG`, `DELETE_FROM_DATABASE`, or `DEPRECATE_IN_DATABASE`. Defaults to `DEPRECATE_IN_DATABASE`.
	DeleteBehavior string `json:"deleteBehavior,omitempty" yaml:"deleteBehavior,omitempty"`

	// The update behavior when the crawler finds a changed schema. Valid values: `LOG` or `UPDATE_IN_DATABASE`. Defaults to `UPDATE_IN_DATABASE`.
	UpdateBehavior string `json:"updateBehavior,omitempty" yaml:"updateBehavior,omitempty"`
}
