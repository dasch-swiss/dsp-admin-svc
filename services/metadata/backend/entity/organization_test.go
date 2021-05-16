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
	"github.com/dasch-swiss/dasch-service-platform/services/metadata/backend/entity"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	org, err := entity.NewOrganization("TEST Org")
	assert.Nil(t, err)
	assert.Equal(t, map[string]bool{"TEST Org": true}, org.Name)
	assert.NotNil(t, org.ID)
	assert.False(t, org.CreatedAt.IsZero())
	assert.True(t, org.UpdatedAt.IsZero())
}

func TestAddName(t *testing.T) {
	org, err := entity.NewOrganization("new org")
	assert.Nil(t, err)

	err2 := org.AddName("additional org name")
	assert.Nil(t, err2)
	assert.Equal(t, map[string]bool{"new org": true, "additional org name": true}, org.Name)
}

func TestNotAllowAddSameName(t *testing.T) {
	org, err := entity.NewOrganization("new org")
	assert.Nil(t, err)

	err2 := org.AddName("new org")
	assert.NotNil(t, err2)
}

func TestNotAllowAddFourthName(t *testing.T) {
	org, err1 := entity.NewOrganization("first name")
	assert.Nil(t, err1)

	err2 := org.AddName("second name")
	assert.Nil(t, err2)

	err3 := org.AddName("third name")
	assert.Nil(t, err3)

	err4 := org.AddName("fourth name")
	assert.NotNil(t, err4)
}

func TestRemoveName(t *testing.T) {
	org, err := entity.NewOrganization("new org")
	assert.Nil(t, err)

	err2 := org.AddName("additional org name")
	assert.Nil(t, err2)

	err3 := org.RemoveName("new org")
	assert.Nil(t, err3)
	assert.Equal(t, 1, len(org.Name))
}

func TestNotAllowRemoveLastName(t *testing.T) {
	org, err1 := entity.NewOrganization("new org")
	assert.Nil(t, err1)

	err2 := org.RemoveName("new org")
	assert.NotNil(t, err2)
	assert.Equal(t, 1, len(org.Name))
}

func TestAddAddress(t *testing.T) {
	org, err1 := entity.NewOrganization("new org")
	assert.Nil(t, err1)

	err2 := org.AddAddress("neue strasse 123", "4123", "Allschwil")
	assert.Nil(t, err2)
	assert.Equal(t, "neue strasse 123", org.PostalAddresses.StreetAddress)
	assert.Equal(t, "4123", org.PostalAddresses.PostalCode)
	assert.Equal(t, "Allschwil", org.PostalAddresses.AddressLocality)
	assert.False(t, org.UpdatedAt.IsZero())
}

func TestAddEmail(t *testing.T) {
	org, err1 := entity.NewOrganization("new org")
	assert.Nil(t, err1)

	email, _ := valueobject.NewEmail("test@example.org")
	err2 := org.AddEmail(email)
	assert.Nil(t, err2)
	assert.Equal(t, email, org.Email)
}

func TestRemoveEmail(t *testing.T) {
	org, _ := entity.NewOrganization("new org")

	email, _ := valueobject.NewEmail("test@example.org")
	err1 := org.AddEmail(email)
	assert.Nil(t, err1)

	err2 := org.RemoveEmail()
	assert.Nil(t, err2)
	assert.Equal(t, valueobject.ZeroEmail(), org.Email)
}

func TestNotAllowRemoveNotSetEmail(t *testing.T) {
	org, _ := entity.NewOrganization("new org")

	err := org.RemoveEmail()
	assert.NotNil(t, err)
}

func TestAddURL(t *testing.T) {
	org, _ := entity.NewOrganization("new org")

	newUrlToAdd, _ := valueobject.NewURL("http://example.org")
	err := org.AddURL(newUrlToAdd)
	assert.Nil(t, err)

	assert.Equal(t, newUrlToAdd.String(), org.URL.String())
}

func TestRemoveURL(t *testing.T) {
	org, _ := entity.NewOrganization("new org")
	newUrlToAdd, _ := valueobject.NewURL("http://example.org")
	err1 := org.AddURL(newUrlToAdd)
	assert.Nil(t, err1)

	err2 := org.RemoveURL()
	assert.Nil(t, err2)
}
