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

package presenter

import (
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
)

//Project data
type Project struct {
	ID          valueobject.Identifier `json:"id"`
	ShortCode   string                 `json:"shortCode"`
	ShortName   string                 `json:"shortName"`
	LongName    string                 `json:"longName"`
	Description string                 `json:"description"`
	CreatedAt   string                 `json:"createdAt"`
	CreatedBy   string                 `json:"createdBy"`
	ChangedAt   string                 `json:"changedAt"`
	ChangedBy   string                 `json:"changedBy"`
	DeletedAt   string                 `json:"deletedAt"`
	DeletedBy   string                 `json:"deletedBy"`
}

//UpdatedProjectShortCode
type UpdatedProjectShortCode struct {
	ID        valueobject.Identifier `json:"id"`
	ShortCode string                 `json:"shortCode"`
}

//UpdatedProjectShortName
type UpdatedProjectShortName struct {
	ID        valueobject.Identifier `json:"id"`
	ShortName string                 `json:"shortName"`
}

//UpdatedProjectLongName
type UpdatedProjectLongName struct {
	ID       valueobject.Identifier `json:"id"`
	LongName string                 `json:"longName"`
}

//UpdatedProjectDescription
type UpdatedProjectDescription struct {
	ID          valueobject.Identifier `json:"id"`
	Description string                 `json:"description"`
}

//DeleteProject data
type DeleteProject struct {
	ID        valueobject.Identifier `json:"id"`
	DeletedAt string                 `json:"deletedAt"`
	DeletedBy string                 `json:"deletedBy"`
}
