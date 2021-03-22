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

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.ListNode, error)
	// Search(query string) ([]*entity.ListNode, error)
	// List() ([]*entity.ListNode, error)
}

//Writer interface
type Writer interface {
	Create(e *entity.ListNode) (entity.ID, error)
	// Update(e *entity.ListNode) error
	// Delete(e *entity.ID) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetListNode(id entity.ID) (*entity.ListNode, error)
	CreateListNode(name string, label string, comment string) (entity.ID, error)
}
