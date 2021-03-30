package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/presenter"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/entity"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/project"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func createProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding project"
		var input struct {
			ShortCode   string    `json:"shortCode"`
			CreatedBy   entity.ID `json:"createdBy"`
			ShortName   string    `json:"shortName"`
			LongName    string    `json:"longName"`
			Description string    `json:"description"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := service.CreateProject(input.ShortCode, input.CreatedBy, input.ShortName, input.LongName, input.Description)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		toJ := &presenter.Project{
			ID:          id,
			ShortCode:   input.ShortCode,
			CreatedBy:   input.CreatedBy,
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

func updateProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating project"

		// get id from url
		vars := mux.Vars(r)
		id, paramsErr := entity.StringToID(vars["id"])
		if paramsErr != nil {
			log.Println(paramsErr.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		// get project info from the body
		var input struct {
			UpdateProjectInfo entity.Project `json:"project"`
		}
		inputErr := json.NewDecoder(r.Body).Decode(&input)
		if inputErr != nil {
			log.Println(inputErr.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		updatedProject, inputErr := service.UpdateProject(id, &input.UpdateProjectInfo)
		if inputErr != nil {
			log.Println(inputErr.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		toJ := &presenter.Project{
			ID:          id,
			ShortCode:   updatedProject.ShortCode,
			CreatedBy:   updatedProject.CreatedBy,
			ShortName:   updatedProject.ShortName,
			LongName:    updatedProject.LongName,
			Description: updatedProject.Description,
			CreatedAt:   updatedProject.CreatedAt,
			UpdatedAt:   updatedProject.UpdatedAt,
		}

		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func getProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading project"
		notFoundErrorMessage := "Project not found"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		project, err := service.GetProject(id)
		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			if err == entity.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(notFoundErrorMessage))
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
		}

		toJ := &presenter.Project{
			ID:          project.ID,
			ShortCode:   project.ShortCode,
			CreatedBy:   project.CreatedBy,
			ShortName:   project.ShortName,
			LongName:    project.LongName,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func getAllProjects(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading projects"
		projects, err := service.GetAllProjects()

		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		if projects == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}

		// initialize struct from presenter to later encode as JSON to return
		toJ := &presenter.Projects{}

		// loop through each project in 'data' and add it to the array of projects
		for _, project := range projects {
			toJ.Projects = append(toJ.Projects, presenter.Project(*project))
		}

		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func deleteProject(service project.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error deleting project"
		notFoundErrorMessage := "Project not found"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		deleteResponse, err := service.DeleteProject(id)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			if err == entity.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(notFoundErrorMessage))
				return
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
		}

		toJ := &presenter.DeleteProjectResponse{
			ID:        deleteResponse.ID,
			DeletedAt: deleteResponse.DeletedAt,
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

//MakeProjectHandlers make url handlers
func MakeProjectHandlers(r *mux.Router, n negroni.Negroni, service project.UseCase) {

	r.Handle("/v1/project", n.With(
		negroni.Wrap(createProject(service)),
	)).Methods("POST", "OPTIONS").Name("createProject")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(updateProject(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProject")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(getProject(service)),
	)).Methods("GET", "OPTIONS").Name("getProject")

	r.Handle("/v1/projects", n.With(
		negroni.Wrap(getAllProjects(service)),
	)).Methods("GET", "OPTIONS").Name("getAllProjects")

	r.Handle("/v1/project/{id}", n.With(
		negroni.Wrap(deleteProject(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteProject")
}
