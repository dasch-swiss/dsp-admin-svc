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

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project"
	projTesting "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project/testing"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	// storage
	repo := projTesting.NewInmem()

	// service implementation
	service := project.NewService(repo)

	// add to Inmem db
	uuid, err := service.CreateProject("aabb", entity.NewID(), "short name", "long project name", "this is a test project")

	// assert there is no error
	assert.Nil(t, err)

	// assert that a uuid was returned
	// maybe this should check if the uuid is actually valid
	assert.NotNil(t, uuid)
}

// func Test_Create_Failure(t *testing.T) {
// 	// storage
// 	repo := projTesting.NewInmem()

// 	// service implementation
// 	service := project.NewService(repo)

// 	// add invalid project to Inmem db
// 	_, err := service.CreateProject("", entity.NewID(), "short name", "long project name", "this is a test project")

// 	// assert there is an error
// 	assert.NotNil(t, err)
// }

func Test_Update(t *testing.T) {
	repo := projTesting.NewInmem()
	service := project.NewService(repo)
	userID := entity.NewID()
	uuid, err := service.CreateProject("aabb", userID, "short name", "long project name", "this is a test project")
	assert.Nil(t, err)

	proj, err2 := service.UpdateProject(uuid, &entity.Project{
		ShortCode:   "ccdd",
		ShortName:   "new short name",
		LongName:    "new long project name",
		Description: "new description",
	})

	assert.Nil(t, err2)
	assert.Equal(t, proj.ID, uuid)
	assert.Equal(t, proj.ShortCode, "ccdd")
	assert.Equal(t, proj.CreatedBy, userID)
	assert.Equal(t, proj.ShortName, "new short name")
	assert.Equal(t, proj.LongName, "new long project name")
	assert.Equal(t, proj.Description, "new description")
	assert.False(t, proj.CreatedAt.IsZero())
	assert.False(t, proj.UpdatedAt.IsZero())
}

func Test_Get(t *testing.T) {
	repo := projTesting.NewInmem()      // storage
	service := project.NewService(repo) // service implementation
	uuid, err := service.CreateProject("aabb", entity.NewID(), "short name", "long project name", "this is a test project")
	assert.Nil(t, err)
	p, err2 := service.GetProject(uuid)
	assert.Nil(t, err2)
	assert.Equal(t, p.LongName, "long project name")
}

func Test_GetAll(t *testing.T) {
	repo := projTesting.NewInmem()      // storage
	service := project.NewService(repo) // service implementation

	_, err := service.CreateProject("aabb", entity.NewID(), "short name", "long project name", "this is a test project")
	assert.Nil(t, err)

	_, err2 := service.CreateProject("ccdd", entity.NewID(), "short name 2", "long project name 2", "this is a test project 2")
	assert.Nil(t, err2)

	ap, err3 := service.GetAllProjects()
	assert.Nil(t, err3)
	assert.Len(t, ap, 2)
	assert.Equal(t, ap[0].ShortCode, "aabb")
	assert.Equal(t, ap[1].ShortCode, "ccdd")
}

func Test_Delete(t *testing.T) {
	repo := projTesting.NewInmem()      // storage
	service := project.NewService(repo) // service implementation

	uuid, err := service.CreateProject("aabb", entity.NewID(), "short name", "long project name", "this is a test project")
	assert.Nil(t, err)

	deleteRes, err2 := service.DeleteProject(uuid)
	assert.Nil(t, err2)
	assert.Equal(t, deleteRes.ID, uuid)
	assert.NotNil(t, deleteRes.DeletedAt)
}
