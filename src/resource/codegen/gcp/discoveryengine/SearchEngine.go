package discoveryengine

import types "libds/gcp/types"

type SearchEngine struct {
	// The collection ID.
	CollectionId string `json:"collectionId,omitempty" yaml:"collectionId,omitempty"`

	// The data stores associated with this engine. For SOLUTION_TYPE_SEARCH type of engines, they can only associate with at most one data store.
	DataStoreIds []string `json:"dataStoreIds,omitempty" yaml:"dataStoreIds,omitempty"`

	// Location.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   Configurations for a Search Engine.
	   Structure is documented below.
	*/
	SearchEngineConfig types.Discoveryengine_SearchEngineSearchEngineConfig `json:"searchEngineConfig,omitempty" yaml:"searchEngineConfig,omitempty"`

	/*
	   Common config spec that specifies the metadata of the engine.
	   Structure is documented below.
	*/
	CommonConfig types.Discoveryengine_SearchEngineCommonConfig `json:"commonConfig,omitempty" yaml:"commonConfig,omitempty"`

	// Required. The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	// Unique ID to use for Search Engine App.
	EngineId string `json:"engineId,omitempty" yaml:"engineId,omitempty"`

	/*
	   The industry vertical that the engine registers. The restriction of the Engine industry vertical is based on DataStore: If unspecified, default to GENERIC. Vertical on Engine has to match vertical of the DataStore liniked to the engine.
	   Default value is `GENERIC`.
	   Possible values are: `GENERIC`, `MEDIA`.
	*/
	IndustryVertical string `json:"industryVertical,omitempty" yaml:"industryVertical,omitempty"`
}
