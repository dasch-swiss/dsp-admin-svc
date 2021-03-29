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

package testing_test

import "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"

//inmem in memory repo
type inmem struct {
	m map[entity.ID]*entity.Project
}

//newInmem create a new in memory repository
func NewInmem() *inmem {
	var m = map[entity.ID]*entity.Project{}
	return &inmem{
		m: m,
	}
}

//Create a project
func (r *inmem) Create(e *entity.Project) (entity.ID, error) {
	r.m[e.ID] = e
	return e.ID, nil
}

//Update a project
func (r *inmem) Update(id entity.ID, data *entity.Project) (*entity.Project, error) {
	r.m[id] = data
	return r.m[id], nil
}

//Get a project
func (r *inmem) Get(id entity.ID) (*entity.Project, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//GetAll get all projects
func (r *inmem) GetAll() ([]*entity.Project, error) {
	ap := make([]*entity.Project, 0, len(r.m))

	for _, val := range r.m {
		ap = append(ap, val)
	}

	return ap, nil
}
