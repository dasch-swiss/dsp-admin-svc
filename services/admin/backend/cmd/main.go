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

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/handler"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/api/middleware"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/config"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/infrastructure/repository"
	"github.com/dasch-swiss/dasch-service-platform/services/admin/backend/usecase/listNode"
	"github.com/dasch-swiss/dasch-service-platform/shared/go/pkg/metric"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
)

func main() {

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	listNodeRepository := repository.NewInmemDB()
	listNodeService := listNode.NewService(listNodeRepository)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		log.Fatal(err.Error())
	}
	r := mux.NewRouter()
	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)

	//organization
	handler.MakeListNodeHandlers(r, *n, listNodeService)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

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

	/*
		// create file server handler to serve public folder relative to workspace root
		fs := http.FileServer(http.Dir("./public"))
		http.Handle("/*", fs)

		// add spa handler to serve for calls to root
		http.HandleFunc("/", spaHandler)

		// add db route handler to serve db.json
		http.HandleFunc("/db", dbHandler)

		// add projects route handler to serve projects
		http.HandleFunc("/projects", projectsHandler)

		// start HTTP server with all the previous attached handlers
		log.Fatal(http.ListenAndServe(":8080", nil))
	*/
}

/*
func spaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./public/index.html")
}

func dbHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./services/metadata/backend/data/db.json")
}

func projectsHandler(responseWriter http.ResponseWriter, request *http.Request) {

	resp, err := http.Get("https://api.staging.dasch.swiss/admin/projects")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	fmt.Fprintf(responseWriter, string(body))
}
*/
