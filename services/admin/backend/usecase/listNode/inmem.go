package listNode

import "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.ListNode
}

//newInmem create a new in memory repository
func newInmem() *inmem {
	var m = map[entity.ID]*entity.ListNode{}
	return &inmem{
		m: m,
	}
}

//Create a ListNode
func (r *inmem) Create(e *entity.ListNode) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a ListNode
func (r *inmem) Get(id entity.ID) (*entity.ListNode, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}
