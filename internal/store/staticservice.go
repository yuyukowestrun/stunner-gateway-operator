package store

import (
	"k8s.io/apimachinery/pkg/types"

	stnrgwv1 "github.com/negativefeast/stunner-gateway-operator/api/v1"
)

var StaticServices = NewStaticServiceStore()

type StaticServiceStore struct {
	Store
}

func NewStaticServiceStore() *StaticServiceStore {
	return &StaticServiceStore{
		Store: NewStore(),
	}
}

// GetAll returns all StaticService objects from the global storage
func (s *StaticServiceStore) GetAll() []*stnrgwv1.StaticService {
	ret := make([]*stnrgwv1.StaticService, 0)

	objects := s.Objects()
	for i := range objects {
		r, ok := objects[i].(*stnrgwv1.StaticService)
		if !ok {
			// this is critical: throw up hands and die
			panic("access to an invalid object in the global StaticServiceStore")
		}

		ret = append(ret, r)
	}

	return ret
}

// GetObject returns a named StaticService object from the global storage
func (s *StaticServiceStore) GetObject(nsName types.NamespacedName) *stnrgwv1.StaticService {
	o := s.Get(nsName)
	if o == nil {
		return nil
	}

	r, ok := o.(*stnrgwv1.StaticService)
	if !ok {
		// this is critical: throw up hands and die
		panic("access to an invalid object in the global StaticServiceStore")
	}

	return r
}

// // AddStaticService adds a StaticService object to the the global storage (this is used mainly for testing)
// func (s *StaticServiceStore) AddStaticService(gc *stnrgwv1.StaticService) {
// 	s.Upsert(gc)
// }
