package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/presenter"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/listNode"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func CreateListNode(service listNode.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding list node"
		var input struct {
			Label     string    `json:"label"`
			Comment   string    `json:"comment"`
			CreatedAt time.Time `json:"createdAt"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateListNode(input.Label, input.Comment)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.ListNode{
			ID:      id,
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

func GetListNode(service listNode.UseCase) http.Handler {
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
			ID:        data.ID,
			Label:     data.Label,
			Comment:   data.Comment,
			CreatedAt: data.CreatedAt,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakeListNodeHandlers make url handlers
func MakeListNodeHandlers(r *mux.Router, n negroni.Negroni, service listNode.UseCase) {
	//r.Handle("/v1/user", n.With(
	//	negroni.Wrap(listUsers(service)),
	//)).Methods("GET", "OPTIONS").Name("listUsers")

	r.Handle("/v1/listnode", n.With(
		negroni.Wrap(CreateListNode(service)),
	)).Methods("POST", "OPTIONS").Name("createListNode")

	r.Handle("/v1/listnode/{id}", n.With(
		negroni.Wrap(GetListNode(service)),
	)).Methods("GET", "OPTIONS").Name("getListNode")

	//r.Handle("/v1/user/{id}", n.With(
	//	negroni.Wrap(deleteUser(service)),
	//)).Methods("DELETE", "OPTIONS").Name("deleteUser")
}
