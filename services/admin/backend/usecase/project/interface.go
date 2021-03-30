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

package project

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//Reader interface
//used in db
type Reader interface {
	Get(id entity.ID) (*entity.Project, error)
	GetAll() ([]*entity.Project, error)
}

//Writer interface
//used in db
type Writer interface {
	Create(e *entity.Project) (entity.ID, error)
	Update(id entity.ID, e *entity.Project) (*entity.Project, error)
	Delete(dp *entity.DeletedProject) (*entity.DeletedProject, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
//used in service
type UseCase interface {
	GetProject(id entity.ID) (*entity.Project, error)
	GetAllProjects() ([]*entity.Project, error)
	CreateProject(shortCode string, createdBy entity.ID, shortName string, longName string, description string) (entity.ID, error)
	UpdateProject(id entity.ID, data *entity.Project) (*entity.Project, error)
	DeleteProject(id entity.ID) (*entity.DeletedProject, error)
}
