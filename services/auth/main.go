/*
 * Copyright 2021 DaSCH - Data and Service Center for the Humanities.
 *
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

package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	port := 3001

	// Init Router
	router := mux.NewRouter()

	// Set up routes
	// -------------
	// API
	// router.HandleFunc("/projects", getProjects).Methods("GET")
	// router.HandleFunc("/projects/{id}", getProject).Methods("GET")
	// Serve frontend from `/public`
	dir := "./services/auth/static"
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(dir)))

	// CORS header
	// TODO: is this a security issue?
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))
	// ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3001"}))
	// TODO: do I even still need CORS, now that all is on the same port?

	addr := fmt.Sprintf(":%v", port)
	srv := &http.Server{
		Handler:      ch(router),
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Run server
	log.Printf("Serving auth at %v, on port %v", addr, port)
	log.Fatal(srv.ListenAndServe())
}
