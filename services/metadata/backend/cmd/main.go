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
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)
	/*
		organizationRepository := repository.NewInmemDB()
		organizationService := organization.NewService(organizationRepository)

		metricService, err := metric.NewPrometheusService()
		if err != nil {
			log.Fatal(err.Error())
		}

		r := mux.NewRouter()

		//handlers
		n := negroni.New(
			negroni.HandlerFunc(middleware.Cors),
			middleware.Metrics(metricService),
			negroni.NewLogger(),
		)

		//organization
		handler.MakeOrganizationHandlers(r, *n, organizationService)

		//spa
		handler.MakeSpaHandlers(r, *n)

		http.Handle("/", r)
		http.Handle("/metrics", promhttp.Handler())
		r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		// start HTTP server with all the previous attached handlers
		logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
		srv := &http.Server{
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			Addr:         ":" + strconv.Itoa(config.API_PORT),
			// FIXME: get rid of deprecated github.com/gorilla/context library
			Handler:  context.ClearHandler(http.DefaultServeMux),
			ErrorLog: logger,
		}
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatal(err.Error())
		}
	*/

	r := mux.NewRouter()

	// add spa handler to serve for calls to root
	r.HandleFunc("/", spaHandler)

	// add db route handler to serve db.json
	r.HandleFunc("/db", dbHandler)

	//if a path not found until now, e.g. "/image/tiny.png"
	//this will look at "./public/image/tiny.png" at filesystem
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// add the router
	http.Handle("/", r)

	// start HTTP server with all the previous attached handlers
	log.Fatal(http.ListenAndServe(":3000", nil))

}

func spaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./public/index.html")
}

func dbHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./services/metadata/backend/data/db.json")
}
