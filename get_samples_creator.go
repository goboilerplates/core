package core

import uuid "github.com/satori/go.uuid"

// CreateDefaultGetSamples .
func CreateDefaultGetSamples(names ...string) GetSamplesInteractor {
	var samples []*SampleEntity
	for _, item := range names {
		samples = append(samples, &SampleEntity{
			ID:   uuid.NewV4().String(),
			Name: item,
		})
	}
	return GetSamplesInteractorImpl{Samples: samples}
}
