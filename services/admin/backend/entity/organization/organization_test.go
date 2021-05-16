/*
 * Copyright 2021 DaSCH - Data and Service Center for the Humanities.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package organization_test

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity/organization"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	org, err := organization.NewOrganization("TEST Org")
	assert.Nil(t, err)
	assert.Equal(t, org.Name, "TEST Org")
	assert.NotNil(t, org.ID)
	assert.False(t, org.CreatedAt.IsZero())
	assert.True(t, org.UpdatedAt.IsZero())
}

func TestAddPostalAddress(t *testing.T) {
	org, err := organization.NewOrganization("new org")
	assert.Nil(t, err)

	err2 := org.AddPostalAddress("neue strasse 123", "4123", "Allschwil")
	assert.Nil(t, err2)
	assert.Equal(t, org.PostalAddresses.StreetAddress, "neue strasse 123")
	assert.Equal(t, org.PostalAddresses.PostalCode, "4123")
	assert.Equal(t, org.PostalAddresses.AddressLocality, "Allschwil")
	assert.False(t, org.UpdatedAt.IsZero())
}
