/*
 * Copyright © 2021 the contributors.
 *
 *  This file is part of the DaSCH Service Platform.
 *
 *  The DaSCH Service Platform is free software: you can
 *  redistribute it and/or modify it under the terms of the
 *  GNU Affero General Public License as published by the
 *  Free Software Foundation, either version 3 of the License,
 *  or (at your option) any later version.
 *
 *  The DaSCH Service Platform is distributed in the hope that
 *  it will be useful, but WITHOUT ANY WARRANTY; without even
 *  the implied warranty of MERCHANTABILITY or FITNESS FOR
 *  A PARTICULAR PURPOSE.  See the GNU Affero General Public
 *  License for more details.
 *
 *  You should have received a copy of the GNU Affero General Public
 *  License along with the DaSCH Service Platform.  If not, see
 *  <http://www.gnu.org/licenses/>.
 *
 */

package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/presenter"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/listNode"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func createListNode(service listNode.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding list node"
		var input struct {
			Name    string `json:"name"`
			Label   string `json:"label"`
			Comment string `json:"comment"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateListNode(input.Name, input.Label, input.Comment)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.ListNode{
			ID:      id,
			Name:    input.Name,
			Label:   input.Label,
			Comment: input.Comment,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getListNode(service listNode.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading list node"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		data, err := service.GetListNode(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.ListNode{
			ID:      data.ID,
			Name:    data.Name,
			Label:   data.Label,
			Comment: data.Comment,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakeUserHandlers make url handlers
func MakeListNodeHandlers(r *mux.Router, n negroni.Negroni, service listNode.UseCase) {
	//r.Handle("/v1/user", n.With(
	//	negroni.Wrap(listUsers(service)),
	//)).Methods("GET", "OPTIONS").Name("listUsers")

	r.Handle("/v1/listnode", n.With(
		negroni.Wrap(createListNode(service)),
	)).Methods("POST", "OPTIONS").Name("createListNode")

	r.Handle("/v1/listnode/{id}", n.With(
		negroni.Wrap(getListNode(service)),
	)).Methods("GET", "OPTIONS").Name("getListNode")

	//r.Handle("/v1/user/{id}", n.With(
	//	negroni.Wrap(deleteUser(service)),
	//)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}
