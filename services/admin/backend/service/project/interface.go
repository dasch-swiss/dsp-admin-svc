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

package project

import (
	"context"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity/project"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
)

//Reader interface
type Reader interface {
	Load(ctx context.Context, id valueobject.Identifier) (*project.Aggregate, error)
	GetProjectIds(ctx context.Context, returnDeletedProjects bool) ([]valueobject.Identifier, error)
	// Search(query string) ([]*project.Aggregate, error)
	// List() ([]*project.Aggregate, error)
}

//Writer interface
type Writer interface {
	Save(ctx context.Context, e *project.Aggregate) (valueobject.Identifier, error)
	// UpdateShortCode(ctx context.Context, id valueobject.Identifier, shortCode valueobject.ShortCode) error
	// Delete(e *entity.ID) error
}

//Repository interface which should be implemented by repositories.
type Repository interface {
	Reader
	Writer
}

//UseCase interface which should be implemented by services.
type UseCase interface {
	GetProject(ctx context.Context, id valueobject.Identifier) (*project.Aggregate, error)
	ListProjects(ctx context.Context, returnDeletedProjects bool) ([]project.Aggregate, error)
	CreateProject(ctx context.Context, shortName valueobject.ShortName, longName valueobject.LongName, description valueobject.Description) (valueobject.Identifier, error)
	UpdateProject(ctx context.Context, id valueobject.Identifier, shortCode valueobject.ShortCode, shortName valueobject.ShortName, longName valueobject.LongName, description valueobject.Description) (*project.Aggregate, error)
	DeleteProject(ctx context.Context, id valueobject.Identifier) (*project.Aggregate, error)
	MigrateProject(ctx context.Context, shortCode valueobject.ShortCode, shortName valueobject.ShortName, longName valueobject.LongName, description valueobject.Description) (valueobject.Identifier, error)
}
