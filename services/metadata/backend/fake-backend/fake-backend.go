package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Representation of a project
type Project struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Metadata    interface{} `json:"metadata"`
}

// All projects that are being served
var projects []Project

// Full text search of a project.
// Returns a slice of Projects where each project matches the search query.
// Note: The query is a regex pattern and is matched against the JSON representation of the project.
func searchProjects(query string) []Project {
	var res []Project
	for _, project := range projects {
		content, _ := json.Marshal(project.Metadata)
		match, _ := regexp.Match("(?i)"+query, content)
		if match {
			res = append(res, project)
		}
	}
	return res
}

// Searches for a element with type == "http://ns.dasch.swiss/repository#Project"
// in a json-shaped []interface{}
func findProjectNode(list []interface{}) map[string]interface{} {
	for _, item := range list {
		innerMap, ok := item.(map[string]interface{})
		if ok {
			tp := innerMap["type"]
			if tp == "http://ns.dasch.swiss/repository#Project" {
				return innerMap
			}
		} else {
			log.Fatal("Failed to parse node")
		}
	}
	return nil
}

// Loads a project from a JSON file.
// Expects this file to be located in ./data/*.json
func loadProject(path string) Project {
	log.Printf("Loading: %v", path)
	// read json
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal json
	jsonMap := make(map[string]interface{})
	err2 := json.Unmarshal(byteValue, &jsonMap)
	if err2 != nil {
		log.Fatal(err)
	}

	// grab actual metadata from JSON
	projMetadata, ok := jsonMap["projectsMetadata"].([]interface{})
	if ok {
		projectMap := findProjectNode(projMetadata)
		id := projectMap["shortcode"].(string)
		name := projectMap["name"].(string)
		description := projectMap["description"].(string)
		return Project{
			ID:          id,
			Name:        name,
			Description: description,
			Metadata:    projMetadata,
		}
	} else {
		log.Fatal("Could not find project in JSON")
		return Project{}
	}
}

// Load Project Data
func loadProjectData() []Project {
	var res []Project

	pathPrefix := "./services/metadata/backend/fake-backend/data/"
	paths, _ := filepath.Glob(pathPrefix + "*.json")

	for _, path := range paths {
		file := filepath.Base(path)
		if !strings.HasPrefix(file, "_") {
			res = append(res, loadProject(path))
		}
	}

	return res
}

// Get projects
// Route: /projecs
func getProjects(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for: %v", r.URL)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "X-Total-Count")
	// TODO: are any of those needed? json-server has them
	// w.Header().Set("Content-Type", "charset=utf-8")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Expires", "-1")
	// w.Header().Set("Pragma", "no-cache")
	// w.Header().Add("X-Total-Count", "10")
	// TODO: do we need links to previous and next and first an last?

	// Request parameters
	query := r.URL.Query().Get("q")
	// TODO: does page start at 0 or 1?
	page, _ := strconv.Atoi(r.URL.Query().Get("_page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("_limit"))

	matches := make([]Project, len(projects))

	if query == "" {
		// no search query all projects are matches
		copy(matches, projects)
	} else {
		// reduce projects by search
		matches = searchProjects(query)
	}
	w.Header().Set("X-Total-Count", strconv.Itoa(len(matches)))
	// paginate
	if len(matches) > 1 && len(matches) > limit && page > 0 && limit > 0 {
		max := len(matches)
		start := (page - 1) * limit
		if start > max {
			start = max
		}
		end := page * limit
		if end > max {
			end = max
		}
		matches = matches[start:end]
	}
	// returns whatever remains
	json.NewEncoder(w).Encode(matches)
}

// Get a single project
// Route /projects/:id
func getProject(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request for: %v", r.URL)

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, item := range projects {
		for item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Project{})
}

func main() {
	port := 3000

	// Init Router
	router := mux.NewRouter()

	// Set up routes
	// -------------
	// API
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/projects/{id}", getProject).Methods("GET")
	// Serve frontend from `/public`
	dir := "./public"
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))

	// CORS header
	// TODO: is this a security issue?
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))
	// ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}))
	// TODO: do I even still need CORS, now that all is on the same port?

	// Load Data
	projects = loadProjectData()
	log.Printf("Loaded Projects: %v", len(projects))

	addr := fmt.Sprintf(":%v", port)
	srv := &http.Server{
		Handler:      ch(router),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run server
	log.Printf("Serving metadata at %v, on port %v", addr, port)
	log.Fatal(srv.ListenAndServe())
}
