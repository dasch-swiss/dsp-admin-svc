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

package presenter

import (
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
)

//Project data
type Project struct {
	ID          entity.ID `json:"id"`
	ShortCode   string    `json:"shortCode"`
	CreatedBy   entity.ID `json:"createdBy"`
	ShortName   string    `json:"shortName"`
	LongName    string    `json:"longName"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

//Projects
type Projects struct {
	Projects []Project `json:"project"`
}

//Delete Response
type DeleteProjectResponse struct {
	ID        entity.ID `json:"id"`
	DeletedAt time.Time `json:"deletedAt"`
}
