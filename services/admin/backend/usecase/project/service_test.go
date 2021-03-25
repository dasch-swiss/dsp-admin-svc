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

// import (
// 	"testing"
// 	"time"

// 	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
// 	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project"
// 	projTesting "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project/testing"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"
// )

// func newMockProject() *entity.Project {
// 	return &entity.Project{
// 		ID:          entity.NewID(),
// 		ShortCode:   "aabb",
// 		CreatedBy:   entity.NewID(),
// 		ShortName:   "short name",
// 		LongName:    "long project name",
// 		Description: "this is a mock project",
// 		CreatedAt:   time.Now(),
// 	}
// }

// func Test_Create(t *testing.T) {
// 	repo := projTesting.NewInmem()      // storage
// 	service := project.NewService(repo) // service implementation
// 	proj := newMockProject()
// 	_, err := service.CreateProject("aabb", uuid.New(), "short name", "long project name", "this is a mock project")
// 	assert.Nil(t, err)
// 	assert.False(t, proj.CreatedAt.IsZero())
// 	assert.True(t, proj.UpdatedAt.IsZero())
// }
