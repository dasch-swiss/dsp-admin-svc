# dsp-repository
Service for browsing, searching, and editing of project metadata


## Services

### Metadata

Service for providing users with project and dataset specific metadata.

#### Front-end

The front-end part is basing on [Svelte](https://svelte.dev). To run it, `yarn` and/or `make` need to be installed. 

1. Clone the repoistory:

```
https://github.com/dasch-swiss/dasch-service-platform.git
```


2. Install the dependencies:

```bash
yarn install
```

or

```bash
make yarn
```

3. Start the application:

```bash
yarn run dev
```

This starts the application on the [localhost:5000](http://localhost:5000).

*Note that you will also need to have [Node.js](https://nodejs.org) installed.*

#### Server

For now, metadata is served from a simple server written in Go.

The server serves the frontend (static file serving on `./public/`) and the metadata on `http://localhost:3000/`.

The route `/` serves the frontend.

The routes `/projects` and `/projects/:id` form a simple metadata API.  
The server serves all data found in `./services/metadata/backend/data/*.json`, where the JSON file follows the data structure as currently provided by DSP-JS-LIB.  
__Note:__ Files starting with underscore (`_`) are excluded. This provides a simple means to leave out files that are not supposed to be public.  
The server supports pagination and full text search.

To run the server locally, use the command `make metadata`.  
To run, build and publish a docker image of the server, use the commands `metadata-docker-run`, `metadata-docker-build` and `metadata-docker-publish`respectively. (`...-run` will build first.)

To use the legacy metadata json-server, use the make targets `make metadata-json-server`, `make metadata-json-server-docker-build`, `make metadata-json-server-docker-run` and `make metadata-json-server-docker-publish`.

### Admin

#### Server

This service depends on a running event store.

For testing, you can create a local event store by running:

```docker run --name esdb-node -it -p 2113:2113 -p 1113:1113 eventstore/eventstore:latest --insecure --run-projections=All --enable-atom-pub-over-http=true```

You can view the event store on http://localhost:2113/

Then run:

```make admin-service-run```

The terminal will hang on:

```/private/var/tmp/_bazel_username/6ae9f9aa2327923c1b3eb247cfa2ec4b/execroot/swiss_dasch_dsp_repository/bazel-out/darwin-fastbuild/bin/services/admin/backend/cmd/cmd_/cmd.runfiles/swiss_dasch_dsp_repository```

This is because of Negroni, which will log out the results of your API requests.

Now you can send requests to the API via Postman or your preferred application/method.

Example create project request:

URL:
```POST http://localhost:8080/v1/projects```

JSON request body:

```json
{
  "shortName": "my proj",
  "longName": "my projects name",
  "description": "description of my project"
}
```

You will then see this project creation event in your event store on http://localhost:2113 under the Stream Browser tab (you may need to refresh the page if you're currently on it).


Example update project request:

URL:
```PUT http://localhost:8080/v1/projects/[uuid]```

JSON request body:

```json
{
  "shortCode": "ffff",
  "shortName": "short",
  "longName": "updated project name",
  "description": "updated description of my project"
}
```

You will then see this project update event in your event store on http://localhost:2113 under the Stream Browser tab (you may need to refresh the page if you're currently on it).

Example delete project request:

URL:
```DELETE http://localhost:8080/v1/projects/[uuid]```

You will then see this project deletion event in your event store on http://localhost:2113 under the Stream Browser tab (you may need to refresh the page if you're currently on it).

To get a list of all the projects:

URL:
```GET http://localhost:8080/v1/projects```

To get a project:

URL:
```GET http://localhost:8080/v1/projects/[uuid]```

## Go dependencies

The Go dependencies are defined inside the `go.mod` and the corresponding `go.sum` files.
To be able to use external dependencies with Bazel, all external dependencies need to be registered with Bazel
so that they can be referenced in the `BUILD.bazel` files.

The steps to add an external dependency are as follows:
1. to add repository and version to `go.mod`, run `go get gihub.com/stretchr/testify`
  (exchange the name of the package with the one you would like to add)
1. from inside this directory, run `go mod download gihub.com/stretchr/testify`
  (exchange the name of the package with the one you would like to add)
1. from the root of the repository, run `make gen-go-deps`

Running the `make gen-go-deps` will regenerate the `deps.bzl` file found
in the root of the repository. This file is loaded inside the `WORKSPACE` file, so that Bazel
can go and fetch the external repositories and make them available to be used in rules.

From then on, you should be able to use the external dependency in a `BUILD.bazel` file like so:

```bazel
go_test(
    name = "entity_test",
    srcs = ["organization_test.go"],
    embed = [":entity"],
    deps = ["@com_github_stretchr_testify//assert"],
)
```

## Docs
Three make commands are available for the docs

`make docs-build` - builds the docs

`make docs-serve` - serves the docs locally; useful for when you want to work on the docs locally

`make docs-publish` - publishes the docs to Github Pages which can then be accessed via https://dasch-swiss.github.io/dasch-service-platform
