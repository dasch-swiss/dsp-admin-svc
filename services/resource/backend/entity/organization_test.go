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
	"github.com/dasch-swiss/dasch-service-platform/services/resource/backend/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	org, err := entity.NewOrganization("TEST Org")
	assert.Nil(t, err)
	assert.Equal(t, org.Name, "TEST Org")
	assert.NotNil(t, org.ID)
	assert.False(t, org.CreatedAt.IsZero())
	assert.True(t, org.UpdatedAt.IsZero())
}

func TestAddPostalAddress(t *testing.T) {
	org, err := entity.NewOrganization("new org")
	assert.Nil(t, err)

	err2 := org.AddPostalAddress("neue strasse 123", "4123", "Allschwil")
	assert.Nil(t, err2)
	assert.Equal(t, org.PostalAddresses.StreetAddress, "neue strasse 123")
	assert.Equal(t, org.PostalAddresses.PostalCode, "4123")
	assert.Equal(t, org.PostalAddresses.AddressLocality, "Allschwil")
	assert.False(t, org.UpdatedAt.IsZero())
}
