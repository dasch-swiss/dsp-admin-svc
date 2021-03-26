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

package project_test

import (
	"testing"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project"
	projTesting "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project/testing"
	"github.com/stretchr/testify/assert"
)

func newMockProject() *entity.Project {
	return &entity.Project{
		ID:          entity.NewID(),
		ShortCode:   "aabb",
		CreatedBy:   entity.NewID(),
		ShortName:   "short name",
		LongName:    "long project name",
		Description: "this is a mock project",
		CreatedAt:   time.Now(),
	}
}

func newMockProject2() *entity.Project {
	return &entity.Project{
		ID:          entity.NewID(),
		ShortCode:   "ccdd",
		CreatedBy:   entity.NewID(),
		ShortName:   "short name 2",
		LongName:    "long project name 2",
		Description: "this is a mock project 2",
		CreatedAt:   time.Now(),
	}
}

func Test_Create(t *testing.T) {
	repo := projTesting.NewInmem()      // storage
	service := project.NewService(repo) // service implementation
	proj := newMockProject()
	_, err := service.CreateProject(proj.ShortCode, proj.CreatedBy, proj.ShortName, proj.LongName, proj.Description)
	assert.Nil(t, err)
	assert.False(t, proj.CreatedAt.IsZero())
	assert.True(t, proj.UpdatedAt.IsZero())
}

func Test_Get(t *testing.T) {
	repo := projTesting.NewInmem()      // storage
	service := project.NewService(repo) // service implementation
	proj := newMockProject()
	uuid, err := service.CreateProject(proj.ShortCode, proj.CreatedBy, proj.ShortName, proj.LongName, proj.Description)
	assert.Nil(t, err)
	p, err2 := service.GetProject(uuid)
	assert.Nil(t, err2)
	assert.Equal(t, p.LongName, proj.LongName)
}

func Test_GetAll(t *testing.T) {
	repo := projTesting.NewInmem()      // storage
	service := project.NewService(repo) // service implementation

	proj := newMockProject()
	_, err := service.CreateProject(proj.ShortCode, proj.CreatedBy, proj.ShortName, proj.LongName, proj.Description)
	assert.Nil(t, err)

	proj2 := newMockProject2()
	_, err2 := service.CreateProject(proj2.ShortCode, proj2.CreatedBy, proj2.ShortName, proj2.LongName, proj2.Description)
	assert.Nil(t, err2)

	ap, err3 := service.GetAllProjects()
	assert.Nil(t, err3)
	assert.Len(t, ap, 2)
	assert.Equal(t, ap[0].ShortCode, "aabb")
	assert.Equal(t, ap[1].ShortCode, "ccdd")
}
