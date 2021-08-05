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

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/middleware"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/presenter"
	projectEntity "github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity/project"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/service/project"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/valueobject"
	"github.com/gorilla/mux"
)

// RequestBody provides a reusable struct to use when decoding the JSON request body.
type RequestBody struct {
	ShortCode   string `json:"shortCode"`
	ShortName   string `json:"shortName"`
	LongName    string `json:"longName"`
	Description string `json:"description"`
}

// createProject creates a project with the provided RequestBody.
func createProject(service project.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token to make sure user is authenticated
		// an object containing the users info is returned by ExtractTokenMetadata
		userInfo, tokenErr := middleware.ExtractTokenMetadata(r)
		if tokenErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(tokenErr.Error()))
			return
		}

		// ensure the user has the required role for the action
		if userInfo.Roles == nil || !checkRoles("Role:SystemAdmin", userInfo.Roles) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(projectEntity.ErrUserDoesNotHaveReadPermission.Error()))
			return
		}

		var input RequestBody
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// convert input strings to value objects
		sc, err := valueobject.NewShortCode(input.ShortCode)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		sn, err := valueobject.NewShortName(input.ShortName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		ln, err := valueobject.NewLongName(input.LongName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		desc, err := valueobject.NewDescription(input.Description)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		id, err := service.CreateProject(ctx, sc, sn, ln, desc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// get the project
		p, err := service.GetProject(ctx, id)
		if err != nil && err == projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(projectEntity.ErrServerNotResponding.Error()))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(projectEntity.ErrNoProjectDataReturned.Error()))
			return
		}

		res := &presenter.Project{
			ID:          id,
			ShortCode:   p.ShortCode().String(),
			ShortName:   p.ShortName().String(),
			LongName:    p.LongName().String(),
			Description: p.Description().String(),
			CreatedAt:   p.CreatedAt().String(),
			CreatedBy:   p.CreatedBy().String(),
			ChangedAt:   p.ChangedAt().String(),
			ChangedBy:   p.ChangedBy().String(),
			DeletedAt:   p.DeletedAt().String(),
			DeletedBy:   p.DeletedBy().String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// updateProject updates a project with the provided RequestBody.
// Updating a project that has been marked as deleted is not possible.
// All fields of the RequestBody must be provided.
// At least one of the values of the provided RequestBody must differ from the current value of the corresponding project field.
// If a value of a field is identical to what it already is, the update will not be performed for that field.
func updateProject(service project.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token to make sure user is authenticated
		// an object containing the users info is returned by ExtractTokenMetadata
		userInfo, tokenErr := middleware.ExtractTokenMetadata(r)
		if tokenErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(tokenErr.Error()))
			return
		}

		// ensure the user has the required role for the action
		if userInfo.Roles == nil || !checkRoles("Role:SystemAdmin", userInfo.Roles) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(projectEntity.ErrUserDoesNotHaveReadPermission.Error()))
			return
		}

		var input RequestBody
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
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
		if err != nil && err == projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err != projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(projectEntity.ErrServerNotResponding.Error()))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(projectEntity.ErrNoProjectDataReturned.Error()))
			return
		}

		// convert input strings to value objects
		sc, err := valueobject.NewShortCode(input.ShortCode)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		sn, err := valueobject.NewShortName(input.ShortName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		ln, err := valueobject.NewLongName(input.LongName)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		desc, err := valueobject.NewDescription(input.Description)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		// update the project
		up, err := service.UpdateProject(ctx, uuid, sc, sn, ln, desc)
		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil && err == projectEntity.ErrNoPropertiesChanged {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(err.Error()))
			return
		}
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(projectEntity.ErrServerNotResponding.Error()))
			return
		}

		res := &presenter.Project{
			ID:          up.ID(),
			ShortCode:   up.ShortCode().String(),
			ShortName:   up.ShortName().String(),
			LongName:    up.LongName().String(),
			Description: up.Description().String(),
			CreatedAt:   up.CreatedAt().String(),
			CreatedBy:   up.CreatedBy().String(),
			ChangedAt:   up.ChangedAt().String(),
			ChangedBy:   up.ChangedBy().String(),
			DeletedAt:   up.DeletedAt().String(),
			DeletedBy:   up.DeletedBy().String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

// getProject gets a project with the provided UUID.
func getProject(service project.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token to make sure user is authenticated
		// an object containing the users info is returned by ExtractTokenMetadata
		userInfo, tokenErr := middleware.ExtractTokenMetadata(r)
		if tokenErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(tokenErr.Error()))
			return
		}

		// ensure the user has the required role for the action
		if userInfo.Roles == nil || !checkRoles("Role:SystemAdmin", userInfo.Roles) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(projectEntity.ErrUserDoesNotHaveReadPermission.Error()))
			return
		}

		// get variables from request url
		vars := mux.Vars(r)

		uuid, err := valueobject.IdentifierFromBytes([]byte(vars["id"]))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get the project
		p, err := service.GetProject(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err != nil && err != projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(projectEntity.ErrNoProjectDataReturned.Error()))
			return
		}

		res := &presenter.Project{
			ID:          p.ID(),
			ShortCode:   p.ShortCode().String(),
			ShortName:   p.ShortName().String(),
			LongName:    p.LongName().String(),
			Description: p.Description().String(),
			CreatedAt:   p.CreatedAt().String(),
			CreatedBy:   p.CreatedBy().String(),
			ChangedAt:   p.ChangedAt().String(),
			ChangedBy:   p.ChangedBy().String(),
			DeletedAt:   p.DeletedAt().String(),
			DeletedBy:   p.DeletedBy().String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

// deleteProject deletes a project with the provided UUID.
func deleteProject(service project.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token to make sure user is authenticated
		// an object containing the users info is returned by ExtractTokenMetadata
		userInfo, tokenErr := middleware.ExtractTokenMetadata(r)
		if tokenErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(tokenErr.Error()))
			return
		}

		// ensure the user has the required role for the action
		if userInfo.Roles == nil || !checkRoles("Role:SystemAdmin", userInfo.Roles) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(projectEntity.ErrUserDoesNotHaveReadPermission.Error()))
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

		// delete the project
		p, err := service.DeleteProject(ctx, uuid)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err != nil && err == projectEntity.ErrProjectHasBeenDeleted {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(projectEntity.ErrServerNotResponding.Error()))
			return
		}
		if p == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(projectEntity.ErrNoProjectDataReturned.Error()))
			return
		}

		res := &presenter.Project{
			ID:          p.ID(),
			ShortCode:   p.ShortCode().String(),
			ShortName:   p.ShortName().String(),
			LongName:    p.LongName().String(),
			Description: p.Description().String(),
			CreatedAt:   p.CreatedAt().String(),
			CreatedBy:   p.CreatedBy().String(),
			ChangedAt:   p.ChangedAt().String(),
			ChangedBy:   p.ChangedBy().String(),
			DeletedAt:   p.DeletedAt().String(),
			DeletedBy:   p.DeletedBy().String(),
		}

		// replace null-values with "null"
		*res = res.NullifyJsonProps()

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

// listProjects gets a list of all projects.
// By default, this only returns active projects.
// ReturnDeletedProjects can be provided in the request body to also return projects marked as deleted.
func listProjects(service project.UseCase) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// check JWT token to make sure user is authenticated
		// an object containing the users info is returned by ExtractTokenMetadata
		userInfo, tokenErr := middleware.ExtractTokenMetadata(r)
		if tokenErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(tokenErr.Error()))
			return
		}

		log.Print(userInfo)

		// ensure the user has the required role for the action
		if userInfo.Roles == nil || !checkRoles("Role:SystemAdmin", userInfo.Roles) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(projectEntity.ErrUserDoesNotHaveReadPermission.Error()))
			return
		}

		var input struct {
			ReturnDeletedProjects bool `json:"returnDeletedProjects"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			input.ReturnDeletedProjects = false // default to false if decoding fails (likely because it wasn't provided)
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
		defer cancel()

		// get all project ids
		projects, err := service.ListProjects(ctx, input.ReturnDeletedProjects)
		w.Header().Set("Content-Type", "application/json")

		if err != nil && err == projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err != nil && err != projectEntity.ErrProjectNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(projectEntity.ErrServerNotResponding.Error()))
			return
		}
		if projects == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(projectEntity.ErrNoProjectDataReturned.Error()))
			return
		}

		var res []presenter.Project

		for _, p := range projects {

			projToAppend := presenter.Project{
				ID:          p.ID(),
				ShortCode:   p.ShortCode().String(),
				ShortName:   p.ShortName().String(),
				LongName:    p.LongName().String(),
				Description: p.Description().String(),
				CreatedAt:   p.CreatedAt().String(),
				CreatedBy:   p.CreatedBy().String(),
				ChangedAt:   p.ChangedAt().String(),
				ChangedBy:   p.ChangedBy().String(),
				DeletedAt:   p.DeletedAt().String(),
				DeletedBy:   p.DeletedBy().String(),
			}

			// replace null-values with "null"
			projToAppend = projToAppend.NullifyJsonProps()

			res = append(res, projToAppend)

		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}
}

func checkRoles(role string, usersRoleList[]interface{}) bool {
	for _, r := range usersRoleList {
		if r == role {
			return true
		}
	}
	return false
}

// MakeProjectHandlers make url handlers for creating, updating, deleting, and getting projects
func MakeProjectHandlers(r *mux.Router, service project.UseCase) {

	r.HandleFunc("/v1/projects", createProject(service)).Methods("POST", "OPTIONS")

	r.HandleFunc("/v1/projects/{id}", updateProject(service)).Methods("PUT", "OPTIONS")

	r.HandleFunc("/v1/projects/{id}", deleteProject(service)).Methods("DELETE", "OPTIONS")

	r.HandleFunc("/v1/projects/{id}", getProject(service)).Methods("GET", "OPTIONS")

	r.HandleFunc("/v1/projects", listProjects(service)).Methods("GET", "OPTIONS")
}
