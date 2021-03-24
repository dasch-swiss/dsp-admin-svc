package repository

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//inmem in memory repo
type inmemdb struct {
	m map[entity.ID]*entity.ListNode
}

//NewInmem create a new in memory repository
func NewInmemDB() *inmemdb {
	var m = map[entity.ID]*entity.ListNode{}
	return &inmemdb{
		m: m,
	}
}

//Create a ListNode
func (r *inmemdb) Create(e *entity.ListNode) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Get a ListNode
func (r *inmemdb) Get(id entity.ID) (*entity.ListNode, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}
