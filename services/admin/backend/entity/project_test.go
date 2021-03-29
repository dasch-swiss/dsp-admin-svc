/*
 * Copyright 2021 Data and Service Center for the Humanities - DaSCH

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

package entity_test

import (
	"testing"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	proj, err := entity.NewProject("aabb", uuid.New(), "short name", "long project name", "this is a test project")
	assert.Nil(t, err)
	assert.NotNil(t, proj.ID)
	assert.Equal(t, proj.ShortCode, "aabb")
	assert.NotNil(t, proj.CreatedBy)
	assert.Equal(t, proj.ShortName, "short name")
	assert.Equal(t, proj.LongName, "long project name")
	assert.Equal(t, proj.Description, "this is a test project")
	assert.False(t, proj.CreatedAt.IsZero())
	assert.True(t, proj.UpdatedAt.IsZero())
}

func TestUpdateProject(t *testing.T) {
	proj, err := entity.NewProject("aabb", uuid.New(), "short name", "long project name", "this is a test project")
	assert.Nil(t, err)

	updatedProj, err2 := proj.UpdateProject(entity.Project{proj.ID, "ccdd", proj.CreatedBy, proj.ShortName, proj.LongName, proj.Description, proj.CreatedAt, proj.UpdatedAt})
	assert.Nil(t, err2)
	assert.Equal(t, updatedProj.ShortCode, "ccdd")
	assert.NotNil(t, updatedProj.CreatedBy)
	assert.Equal(t, updatedProj.ShortName, "short name")
	assert.Equal(t, updatedProj.LongName, "long project name")
	assert.Equal(t, updatedProj.Description, "this is a test project")
	assert.False(t, updatedProj.CreatedAt.IsZero())
	assert.False(t, updatedProj.UpdatedAt.IsZero())
}
