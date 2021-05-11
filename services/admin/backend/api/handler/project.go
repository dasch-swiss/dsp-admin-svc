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

package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/presenter"
	projectEntity "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity/project"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/service/project"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func createProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error creating project"
		var input struct {
			ShortCode   string `json:"shortCode"`
			ShortName   string `json:"shortName"`
			LongName    string `json:"longName"`
			Description string `json:"description"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		id, err := service.CreateProject(ctx, input.ShortCode, input.ShortName, input.LongName, input.Description)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Project{
			ID:          id,
			ShortCode:   input.ShortCode,
			ShortName:   input.ShortName,
			LongName:    input.LongName,
			Description: input.Description,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProjectShortCode(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating project short code"
		var input struct {
			ShortCode string `json:"shortCode"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Project has been deleted"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		sc := p.ShortCode()

		if input.ShortCode != "" && sc.String() != input.ShortCode {
			usc, err := service.UpdateProjectShortCode(ctx, uuid, input.ShortCode)
			if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Project has been deleted"))
				return
			}
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			sc = usc.ShortCode()
		}

		toJ := &presenter.UpdatedProjectShortCode{
			ID:        p.ID(),
			ShortCode: sc.String(),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProjectShortName(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating project short name"
		var input struct {
			ShortName string `json:"shortName"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Project has been deleted"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		sn := p.ShortName()

		if input.ShortName != "" && sn.String() != input.ShortName {
			usn, err := service.UpdateProjectShortName(ctx, uuid, input.ShortName)
			if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Project has been deleted"))
				return
			}
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			sn = usn.ShortName()
		}

		toJ := &presenter.UpdatedProjectShortName{
			ID:        p.ID(),
			ShortName: sn.String(),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProjectLongName(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating project long name"
		var input struct {
			LongName string `json:"longName"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Project has been deleted"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		ln := p.LongName()

		if input.LongName != "" && ln.String() != input.LongName {
			uln, err := service.UpdateProjectLongName(ctx, uuid, input.LongName)
			if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Project has been deleted"))
				return
			}
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			ln = uln.LongName()
		}

		toJ := &presenter.UpdatedProjectLongName{
			ID:       p.ID(),
			LongName: ln.String(),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProjectDescription(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating project"
		var input struct {
			Description string `json:"description"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Project has been deleted"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		desc := p.Description()

		if input.Description != "" && desc.String() != input.Description {
			ud, err := service.UpdateProjectDescription(ctx, uuid, input.Description)
			if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Project has been deleted"))
				return
			}
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}

			desc = ud.Description()
		}

		toJ := &presenter.UpdatedProjectDescription{
			ID:          p.ID(),
			Description: desc.String(),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		data, err := service.GetProject(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		toJ := &presenter.Project{
			ID:          data.ID(),
			ShortCode:   data.ShortCode().String(),
			ShortName:   data.ShortName().String(),
			LongName:    data.LongName().String(),
			Description: data.Description().String(),
			CreatedAt:   data.CreatedAt().String(),
			CreatedBy:   data.CreatedBy().String(),
			ChangedAt:   data.ChangedAt().String(),
			ChangedBy:   data.ChangedBy().String(),
			DeletedAt:   data.DeletedAt().String(),
			DeletedBy:   data.DeletedBy().String(),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed encoding data to JSON"))
		}
	})
}

func deleteProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get variables from request url
		vars := mux.Vars(r)

		// create empty Identifier
		uuid := valueobject.Identifier{}

		// create byte array from the provided id string
		b := []byte(vars["id"])

		// assign the value of the Identifier
		uuid.UnmarshalText(b)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// delete the project
		data, err := service.DeleteProject(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Project has already been deleted"))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No data was returned"))
			return
		}

		toJ := &presenter.DeleteProject{
			ID:        data.ID(),
			DeletedAt: data.ChangedAt().String(),
			DeletedBy: data.ChangedBy().String(),
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed encoding data to JSON"))
		}
	})
}

func listProjects(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get all project ids
		ids, err := service.ListProjects(ctx)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No project found for this uuid"))
			return
		}

		if err != nil && err != projectEntity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("The server is not responding"))
			return
		}
		if ids == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No ids was returned"))
			return
		}

		var toJ []presenter.Project

		for _, id := range ids {
			p, err := service.GetProject(ctx, id)
			if err != nil && err == projectEntity.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("No project found for this uuid"))
				return
			}

			if err != nil && err != projectEntity.ErrNotFound {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("The server is not responding"))
				return
			}
			if p == nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("No data was returned"))
				return
			}

			toJ = append(toJ, presenter.Project{
				ID:          p.ID(),
				ShortCode:   p.ShortCode().String(),
				ShortName:   p.ShortName().String(),
				LongName:    p.LongName().String(),
				Description: p.Description().String(),
			})
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed encoding ids to JSON"))
		}
	})
}

//MakeProjectHandlers make url handlers
func MakeProjectHandlers(r *mux.Router, n negroni.Negroni, service project.UseCase) {

	r.Handle("/v1/project", n.With(
		negroni.Wrap(createProject(service)),
	)).Methods("POST", "OPTIONS").Name("createProject")

	r.Handle("/v1/project/{id}/shortcode", n.With(
		negroni.Wrap(updateProjectShortCode(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProjectShortCode")

	r.Handle("/v1/project/{id}/shortname", n.With(
		negroni.Wrap(updateProjectShortName(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProjectShortName")

	r.Handle("/v1/project/{id}/longname", n.With(
		negroni.Wrap(updateProjectLongName(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProjectLongName")

	r.Handle("/v1/project/{id}/description", n.With(
		negroni.Wrap(updateProjectDescription(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProjectDescription")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(deleteProject(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteProject")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(getProject(service)),
	)).Methods("GET", "OPTIONS").Name("getProject")

	r.Handle("/v1/projects/all", n.With(
		negroni.Wrap(listProjects(service)),
	)).Methods("GET", "OPTIONS").Name("listProjects")
}
