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

import (
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//inmem in memory repo
type inmem struct {
	projects map[entity.ID]*entity.Project
}

//newInmem create a new in memory repository
func NewInmem() *inmem {
	var newProjectsMap = map[entity.ID]*entity.Project{}
	return &inmem{
		projects: newProjectsMap,
	}
}

//Create a project
func (repository *inmem) Create(project *entity.Project) (entity.ID, error) {
	repository.projects[project.ID] = project
	return project.ID, nil
}

//Update a project
//project contains the updated project
func (respository *inmem) Update(id entity.ID, project *entity.Project) (*entity.Project, error) {
	respository.projects[id] = project
	return respository.projects[id], nil
}

//Get a project
func (respository *inmem) Get(id entity.ID) (*entity.Project, error) {
	if respository.projects[id] == nil {
		return nil, entity.ErrNotFound
	}
	return respository.projects[id], nil
}

//GetAll get all projects
func (respository *inmem) GetAll() ([]*entity.Project, error) {
	allProjects := make([]*entity.Project, 0, len(respository.projects))

	for _, val := range respository.projects {
		allProjects = append(allProjects, val)
	}

	return allProjects, nil
}

//Delete a project
//deletedProject is the project to be deleted
func (respository *inmem) Delete(deletedProject *entity.DeletedProject) (*entity.DeletedProject, error) {
	//make sure the key exists
	if respository.projects[deletedProject.ID] == nil {
		return nil, entity.ErrNotFound
	}

	delete(respository.projects, deletedProject.ID)
	return deletedProject, nil
}
