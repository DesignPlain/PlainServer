package discoveryengine

import types "libds/gcp/types"

type ChatEngine struct {
	/*
	   Configurations for a chat Engine.
	   Structure is documented below.
	*/
	ChatEngineConfig types.Discoveryengine_ChatEngineChatEngineConfig `json:"chatEngineConfig,omitempty" yaml:"chatEngineConfig,omitempty"`

	/*
	   Common config spec that specifies the metadata of the engine.
	   Structure is documented below.
	*/
	CommonConfig types.Discoveryengine_ChatEngineCommonConfig `json:"commonConfig,omitempty" yaml:"commonConfig,omitempty"`

	// The ID to use for chat engine.
	EngineId string `json:"engineId,omitempty" yaml:"engineId,omitempty"`

	/*
	   The industry vertical that the chat engine registers. Vertical on Engine has to match vertical of the DataStore linked to the engine.
	   Default value is `GENERIC`.
	   Possible values are: `GENERIC`.
	*/
	IndustryVertical string `json:"industryVertical,omitempty" yaml:"industryVertical,omitempty"`

	// Location.
	Location string `json:"location,omitempty" yaml:"location,omitempty"`

	// The collection ID.
	CollectionId string `json:"collectionId,omitempty" yaml:"collectionId,omitempty"`

	// The data stores associated with this engine. Multiple DataStores in the same Collection can be associated here. All listed DataStores must be `SOLUTION_TYPE_CHAT`. Adding or removing data stores will force recreation.
	DataStoreIds []string `json:"dataStoreIds,omitempty" yaml:"dataStoreIds,omitempty"`

	// The display name of the engine. Should be human readable. UTF-8 encoded string with limit of 1024 characters.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
