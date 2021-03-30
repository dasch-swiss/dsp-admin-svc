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

package project_repository

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//inmem in memory repo
type inmemdb struct {
	m map[entity.ID]*entity.Project
}

//NewInmem create a new in memory repository
func NewInmemDB() *inmemdb {
	var m = map[entity.ID]*entity.Project{}
	return &inmemdb{
		m: m,
	}
}

//Create a project
func (r *inmemdb) Create(e *entity.Project) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Update a project
func (r *inmemdb) Update(id entity.ID, updatedProject *entity.Project) (*entity.Project, error) {
	// find project using the provided id and replace it with the provided updatedProject
	r.m[id] = updatedProject
	return r.m[id], nil
}

//Get a project
func (r *inmemdb) Get(id entity.ID) (*entity.Project, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//GetAll get all projects
func (r *inmemdb) GetAll() ([]*entity.Project, error) {
	ap := make([]*entity.Project, 0, len(r.m))

	for _, val := range r.m {
		ap = append(ap, val)
	}

	return ap, nil
}

//Delete a project
func (r *inmemdb) Delete(projectToDelete *entity.DeletedProject) (*entity.DeletedProject, error) {
	if r.m[projectToDelete.ID] == nil {
		return nil, entity.ErrNotFound
	}
	delete(r.m, projectToDelete.ID)
	return projectToDelete, nil
}
