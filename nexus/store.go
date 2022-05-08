package nexus

import "gozealous/service"

func NewStore(
	nativeService service.Native,
	railwayService service.Railway,
) *Store {
	return &Store{
		nativeSvc:  nativeService,
		railwaySvc: railwayService,
	}
}
