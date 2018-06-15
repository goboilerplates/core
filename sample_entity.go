package core

import "strings"

// SampleEntity is the sample entity struct.
type SampleEntity struct {
	ID   string
	Name string
}

// NameLower gets lowercase name.
func (sampleEntity SampleEntity) NameLower() string {
	return strings.ToLower(sampleEntity.Name)
}
