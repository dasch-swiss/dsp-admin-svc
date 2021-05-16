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

package user_test

import (
	"context"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/service/user"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUser_SignupUser(t *testing.T) {

	expectedAggregateType := "http://ns.dasch.swiss/admin#User"
	expectedUsername := "dduck"
	expectedEmail := "dduck@example.com"
	expectedPassword := "secret"
	expectedGivenName := "Donald"
	expectedFamilyName := "Duck"

	repo := NewInMemRepo()
	service := user.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	userId, err := service.SignupUser(ctx, expectedUsername, expectedEmail, expectedPassword, expectedGivenName, expectedFamilyName)
	assert.Nil(t, err)

	foundUser, err := service.GetUser(ctx, userId)
	assert.Nil(t, err)
	assert.Equal(t, expectedAggregateType, foundUser.AggregateType().String())
	assert.Equal(t, expectedUsername, foundUser.Username().String())
	assert.Equal(t, expectedEmail, foundUser.Email().String())
	assert.Equal(t, expectedPassword, foundUser.Password().String())
	assert.Equal(t, expectedGivenName, foundUser.GivenName().String())
	assert.Equal(t, expectedFamilyName, foundUser.FamilyName().String())
}
