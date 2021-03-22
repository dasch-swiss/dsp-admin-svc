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

package entity_test

import (
	"testing"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewListNode(t *testing.T) {
	node, err := entity.NewListNode("TEST name", "TEST label", "TEST comment")
	assert.Nil(t, err)
	assert.Equal(t, node.Name, "TEST name")
	assert.Equal(t, node.Label, "TEST label")
	assert.Equal(t, node.Comment, "TEST comment")
	assert.NotNil(t, node.ID)
	assert.False(t, node.CreatedAt.IsZero())
	assert.True(t, node.UpdatedAt.IsZero())
}
