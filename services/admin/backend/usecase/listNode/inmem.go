/*
 * Copyright © 2021 the contributors.
 *
 *  This file is part of the DaSCH Service Platform.
 *
 *  The DaSCH Service Platform is free software: you can
 *  redistribute it and/or modify it under the terms of the
 *  GNU Affero General Public License as published by the
 *  Free Software Foundation, either version 3 of the License,
 *  or (at your option) any later version.
 *
 *  The DaSCH Service Platform is distributed in the hope that
 *  it will be useful, but WITHOUT ANY WARRANTY; without even
 *  the implied warranty of MERCHANTABILITY or FITNESS FOR
 *  A PARTICULAR PURPOSE.  See the GNU Affero General Public
 *  License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public
 *  License along with the DaSCH Service Platform.  If not, see
 *  <http://www.gnu.org/licenses/>.
 *
 */

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
