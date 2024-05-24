package types

type Connect_getUserHierarchyStructureHierarchyStructure struct {
	// Details of level three. See below.
	LevelThrees []Connect_getUserHierarchyStructureHierarchyStructureLevelThree `json:"levelThrees,omitempty" yaml:"levelThrees,omitempty"`

	// Details of level two. See below.
	LevelTwos []Connect_getUserHierarchyStructureHierarchyStructureLevelTwo `json:"levelTwos,omitempty" yaml:"levelTwos,omitempty"`

	// Details of level five. See below.
	LevelFives []Connect_getUserHierarchyStructureHierarchyStructureLevelFife `json:"levelFives,omitempty" yaml:"levelFives,omitempty"`

	// Details of level four. See below.
	LevelFours []Connect_getUserHierarchyStructureHierarchyStructureLevelFour `json:"levelFours,omitempty" yaml:"levelFours,omitempty"`

	// Details of level one. See below.
	LevelOnes []Connect_getUserHierarchyStructureHierarchyStructureLevelOne `json:"levelOnes,omitempty" yaml:"levelOnes,omitempty"`
}
