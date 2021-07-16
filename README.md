## Services

### Admin

#### Server

This service requires a running event store.

For testing, you can create a local event store by running:

```docker run --name esdb-node -it -p 2113:2113 -p 1113:1113 eventstore/eventstore:latest --insecure --run-projections=All --enable-atom-pub-over-http=true```

You can view the event store on http://localhost:2113/

Then run:
```make admin-service-run```

The terminal will hang on:

```2021/07/08 14:01:33 Starting server...```

```2021/07/08 14:01:33 Serving on port: :8080```

This is because of Negroni, which will log out the results of your API requests.
Now you can send requests to the API via Postman or your preferred application/method.

**A valid JWT token must be provided with each API request**

Example create project request:
URL:
```POST http://localhost:8080/v1/projects```

Headers:
```json
{
  "Authorization": "bearer [JWT]"
}
```

JSON request body:
```json
{
  "shortCode": "0000",
  "shortName": "my proj",
  "longName": "my projects name",
  "description": "description of my project"
}
```

You will then see this project creation event in your event store on http://localhost:2113 under the Stream Browser tab (you may need to refresh the page if you're currently on it).


Example update project request:

URL:
```PUT http://localhost:8080/v1/projects/[uuid]```

Headers:
```json
{
  "Authorization": "bearer [JWT]"
}
```

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

Headers:
```json
{
  "Authorization": "bearer [JWT]"
}
```

You will then see this project deletion event in your event store on http://localhost:2113 under the Stream Browser tab (you may need to refresh the page if you're currently on it).

To get a list of all the projects:

URL:
```GET http://localhost:8080/v1/projects```

Headers:
```json
{
  "Authorization": "bearer [JWT]"
}
```

To get a project:

URL:
```GET http://localhost:8080/v1/projects/[uuid]```

Headers:
```json
{
  "Authorization": "bearer [JWT]"
}
```

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
