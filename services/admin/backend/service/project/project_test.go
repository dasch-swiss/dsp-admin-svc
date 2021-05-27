/*
 *  Copyright 2021 Data and Service Center for the Humanities - DaSCH.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package project_test

import (
	"context"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"testing"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/service/project"
	"github.com/stretchr/testify/assert"
)

// TODO: split this into different methods and use a mock project
func TestProject_CreateProject(t *testing.T) {

	expectedAggregateType := "http://ns.dasch.swiss/admin#Project"
	expectedShortName := "short name"
	expectedLongName := "project long name"
	expectedDescription := "project description"

	repo := NewInMemRepo()
	service := project.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create short name value object
	sn, err := valueobject.NewShortName(expectedShortName)
	assert.Nil(t, err)

	// create long name value object
	ln, err := valueobject.NewLongName(expectedLongName)
	assert.Nil(t, err)

	// create short code value object
	desc, err := valueobject.NewDescription(expectedDescription)
	assert.Nil(t, err)

	projectId, err := service.CreateProject(ctx, sn, ln, desc)
	assert.Nil(t, err)

	// get project
	foundProject, err := service.GetProject(ctx, projectId)
	assert.Nil(t, err)
	assert.Equal(t, expectedAggregateType, foundProject.AggregateType().String())
	assert.Len(t, foundProject.ShortCode().String(), 4)
	assert.Equal(t, expectedShortName, foundProject.ShortName().String())
	assert.Equal(t, expectedLongName, foundProject.LongName().String())
	assert.Equal(t, expectedDescription, foundProject.Description().String())

	expectedUpdatedShortCode := "11AA"
	expectedUpdatedShortName := "new short name"
	expectedUpdatedLongName := "new project long name"
	expectedUpdatedDescription := "new project description"

	// create new short code value object
	nsc, err := valueobject.NewShortCode("11AA")
	assert.Nil(t, err)

	// update short code
	usc, err := service.UpdateProject(ctx, foundProject.ID(), nsc, foundProject.ShortName(), foundProject.LongName(), foundProject.Description())
	assert.Nil(t, err)

	// assert short code was updated
	assert.Equal(t, expectedUpdatedShortCode, usc.ShortCode().String())

	// the remaining fields should remain unchanged
	assert.Equal(t, expectedShortName, usc.ShortName().String())
	assert.Equal(t, expectedLongName, usc.LongName().String())
	assert.Equal(t, expectedDescription, usc.Description().String())
	assert.NotZero(t, usc.ChangedAt())
	// TODO: assert ChangedBy is not an empty UUID

	// create new short name value object
	nsn, err := valueobject.NewShortName("new short name")
	assert.Nil(t, err)

	// update short name
	usn, err := service.UpdateProject(ctx, foundProject.ID(), nsc, nsn, foundProject.LongName(), foundProject.Description())
	assert.Nil(t, err)

	// short code should remain the updated short code
	assert.Equal(t, expectedUpdatedShortCode, usn.ShortCode().String())

	// assert short code was updated
	assert.Equal(t, expectedUpdatedShortName, usn.ShortName().String())

	// the remaining fields should remain unchanged
	assert.Equal(t, expectedLongName, usn.LongName().String())
	assert.Equal(t, expectedDescription, usn.Description().String())

	// create new long name value object
	nln, err := valueobject.NewLongName("new project long name")
	assert.Nil(t, err)

	// update long name
	uln, err := service.UpdateProject(ctx, foundProject.ID(), nsc, nsn, nln, foundProject.Description())
	assert.Nil(t, err)

	// short code should remain the updated short code
	assert.Equal(t, expectedUpdatedShortCode, uln.ShortCode().String())

	// short name should remain the updated short name
	assert.Equal(t, expectedUpdatedShortName, uln.ShortName().String())

	// assert long name was updated
	assert.Equal(t, expectedUpdatedLongName, uln.LongName().String())

	// description should remain the expectedDescription
	assert.Equal(t, expectedDescription, uln.Description().String())

	// create new description value object
	nd, err := valueobject.NewDescription("new project description")
	assert.Nil(t, err)

	// update description
	ud, err := service.UpdateProject(ctx, foundProject.ID(), nsc, nsn, nln, nd)
	assert.Nil(t, err)

	// short code should remain the updated short code
	assert.Equal(t, expectedUpdatedShortCode, ud.ShortCode().String())

	// short name should remain the updated short name
	assert.Equal(t, expectedUpdatedShortName, ud.ShortName().String())

	// long name should remain the updated long name
	assert.Equal(t, expectedUpdatedLongName, ud.LongName().String())

	// assert description was updated
	assert.Equal(t, expectedUpdatedDescription, ud.Description().String())

	// get a list of projects
	projectsList, err := service.ListProjects(ctx, false)
	assert.Nil(t, err)
	assert.Len(t, projectsList, 1)
	assert.Equal(t, projectsList[0].ID(), projectId)

	// delete a project
	deletedProject, err := service.DeleteProject(ctx, projectId)
	assert.Nil(t, err)
	assert.NotZero(t, deletedProject.DeletedAt())
	// TODO: assert DeletedBy is not an empty UUID
}

func TestProject_MigrateProject(t *testing.T) {

	expectedAggregateType := "http://ns.dasch.swiss/admin#Project"
	expectedShortCode := "00FF"
	expectedShortName := "short name"
	expectedLongName := "project long name"
	expectedDescription := "project description"

	repo := NewInMemRepo()
	service := project.NewService(repo)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	// create short code value object
	sc, err := valueobject.NewShortCode(expectedShortCode)
	assert.Nil(t, err)

	// create short name value object
	sn, err := valueobject.NewShortName(expectedShortName)
	assert.Nil(t, err)

	// create long name value object
	ln, err := valueobject.NewLongName(expectedLongName)
	assert.Nil(t, err)

	// create short code value object
	desc, err := valueobject.NewDescription(expectedDescription)
	assert.Nil(t, err)

	// migrate (create) project
	projectId, err := service.MigrateProject(ctx, sc, sn, ln, desc)
	assert.Nil(t, err)

	// get migrated project
	migratedProject, err := service.GetProject(ctx, projectId)
	assert.Nil(t, err)
	assert.Equal(t, expectedAggregateType, migratedProject.AggregateType().String())
	assert.Equal(t, expectedShortCode, migratedProject.ShortCode().String())
	assert.Equal(t, expectedShortName, migratedProject.ShortName().String())
	assert.Equal(t, expectedLongName, migratedProject.LongName().String())
	assert.Equal(t, expectedDescription, migratedProject.Description().String())
}
