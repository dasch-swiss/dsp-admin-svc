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

package entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

//Project domain entity
type Project struct {
	ID          ID
	ShortCode   string
	CreatedBy   ID
	ShortName   string
	LongName    string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewProject(shortCode string, createdBy ID, shortName string, longName string, description string) (*Project, error) {
	proj := &Project{
		ID:          NewID(),
		ShortCode:   shortCode,
		CreatedBy:   createdBy,
		ShortName:   shortName,
		LongName:    longName,
		Description: description,
		CreatedAt:   time.Now(),
	}

	err := proj.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return proj, nil
}

func (proj *Project) UpdateProject(updateProjectInfo Project) (*Project, error) {

	// only update these properties if the a non-empty string was provided
	if strings.TrimSpace(updateProjectInfo.ShortCode) != "" {
		proj.ShortCode = updateProjectInfo.ShortCode
	}

	if strings.TrimSpace(updateProjectInfo.ShortName) != "" {
		proj.ShortName = updateProjectInfo.ShortName
	}

	if strings.TrimSpace(updateProjectInfo.LongName) != "" {
		proj.LongName = updateProjectInfo.LongName
	}

	if strings.TrimSpace(updateProjectInfo.Description) != "" {
		proj.Description = updateProjectInfo.Description
	}

	// assign current time to UpdateAt property
	proj.UpdatedAt = time.Now()

	return proj, nil
}

func (proj *Project) Validate() error {

	// switch statement would probably look better

	if proj.ShortCode == "" {
		return ErrInvalidEntity
	}

	if proj.CreatedBy == uuid.Nil {
		return ErrInvalidEntity
	}

	if proj.ShortName == "" {
		return ErrInvalidEntity
	}

	if proj.LongName == "" {
		return ErrInvalidEntity
	}

	if proj.Description == "" {
		return ErrInvalidEntity
	}

	return nil
}
