package types

type Connect_UserHierarchyGroupHierarchyPath struct {
	// A block that defines the details of level five. The level block is documented below.
	LevelFives []Connect_UserHierarchyGroupHierarchyPathLevelFife `json:"levelFives,omitempty" yaml:"levelFives,omitempty"`

	// A block that defines the details of level four. The level block is documented below.
	LevelFours []Connect_UserHierarchyGroupHierarchyPathLevelFour `json:"levelFours,omitempty" yaml:"levelFours,omitempty"`

	// A block that defines the details of level one. The level block is documented below.
	LevelOnes []Connect_UserHierarchyGroupHierarchyPathLevelOne `json:"levelOnes,omitempty" yaml:"levelOnes,omitempty"`

	// A block that defines the details of level three. The level block is documented below.
	LevelThrees []Connect_UserHierarchyGroupHierarchyPathLevelThree `json:"levelThrees,omitempty" yaml:"levelThrees,omitempty"`

	// A block that defines the details of level two. The level block is documented below.
	LevelTwos []Connect_UserHierarchyGroupHierarchyPathLevelTwo `json:"levelTwos,omitempty" yaml:"levelTwos,omitempty"`
}
