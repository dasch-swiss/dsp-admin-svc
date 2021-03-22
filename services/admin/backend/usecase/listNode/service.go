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

//Service interface
type Service struct {
	repo Repository
}

//NewService create a new organization use case
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateListNode create a ListNode
func (s *Service) CreateListNode(name string, label string, comment string) (entity.ID, error) {
	e, err := entity.NewListNode(name, label, comment)
	if err != nil {
		return e.ID, err
	}
	return s.repo.Create(e)
}

//GetListNode get a ListNode
func (s *Service) GetListNode(id entity.ID) (*entity.ListNode, error) {
	return s.repo.Get(id)
}
