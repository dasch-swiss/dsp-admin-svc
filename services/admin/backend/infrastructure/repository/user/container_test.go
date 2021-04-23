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

package user_test

import (
	"crypto/tls"
	"fmt"
	"github.com/EventStore/EventStore-Client-Go/client"
	"github.com/ory/dockertest/v3"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"testing"
)

const EVENTSTORE_DOCKER_REPOSITORY = "eventstore/eventstore"
const EVENTSTORE_DOCKER_TAG = "latest"
const EVENTSTORE_DOCKER_PORT = "2113"

// Container structure.
type Container struct {
	Endpoint string
	Resource *dockertest.Resource
}

// Close shuts down the Docker container.
func (container *Container) Close() {
	container.Resource.Close()
}

func GetEmptyDatabase() *Container {
	options := &dockertest.RunOptions{
		Repository:   EVENTSTORE_DOCKER_REPOSITORY,
		Tag:          EVENTSTORE_DOCKER_TAG,
		ExposedPorts: []string{EVENTSTORE_DOCKER_PORT},
		Env:          []string{"EVENTSTORE_INSECURE=true"},
	}
	return getDatabase(options)
}

func GetPrePopulatedDatabase() *Container {
	options := &dockertest.RunOptions{
		Repository:   EVENTSTORE_DOCKER_REPOSITORY,
		Tag:          EVENTSTORE_DOCKER_TAG,
		ExposedPorts: []string{EVENTSTORE_DOCKER_PORT},
		Env:          []string{"EVENTSTORE_INSECURE=true", "EVENTSTORE_RUN_PROJECTIONS=All"},
	}
	return getDatabase(options)
}

func getDatabase(options *dockertest.RunOptions) *Container {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker. Reason: %v", err)
	}

	//err = setTLSContext(options)
	//if err != nil {
	//	log.Fatal(err)
	//}

	resource, err := pool.RunWithOptions(options)
	if err != nil {
		log.Fatalf("Could not start resource. Reason: %v", err)
	}

	endpoint := fmt.Sprintf("localhost:%s", resource.GetPort("2113/tcp"))

	// Disable certificate verification
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	err = pool.Retry(func() error {
		if resource != nil && resource.Container != nil {
			c, e := pool.Client.InspectContainer(resource.Container.ID)
			if e == nil && c.State.Running == false {
				return fmt.Errorf("unexpected exit of container check the container logs for more information, container ID: %v", resource.Container.ID)
			}
		}

		healthCheckEndpoint := fmt.Sprintf("http://%s", endpoint)
		_, err := http.Get(healthCheckEndpoint)
		return err
	})

	if err != nil {
		log.Fatalf("HealthCheck failed. Reason: %v", err)
	}

	return &Container{
		Endpoint: endpoint,
		Resource: resource,
	}
}

func setTLSContext(options *dockertest.RunOptions) error {
	err := verifyCertificatesExist()
	if err != nil {
		return err
	}

	options.Env = append(options.Env, []string{
		"EVENTSTORE_CERTIFICATE_FILE=/etc/eventstore/certs/node/node.crt",
		"EVENTSTORE_CERTIFICATE_PRIVATE_KEY_FILE=/etc/eventstore/certs/node/node.key",
		"EVENTSTORE_TRUSTED_ROOT_CERTIFICATES_PATH=/etc/eventstore/certs/ca",
	}...)

	certsDir, err := getCertificatesDir()
	if err != nil {
		return err
	}
	options.Mounts = []string{
		fmt.Sprintf("%v:/etc/eventstore/certs", certsDir),
	}
	return nil
}

func verifyCertificatesExist() error {
	certs := []string{
		path.Join("ca", "ca.crt"),
		path.Join("ca", "ca.key"),
		path.Join("node", "node.crt"),
		path.Join("node", "node.key"),
	}

	certsDir, err := getCertificatesDir()
	if err != nil {
		return err
	}

	for _, f := range certs {
		if _, err := os.Stat(path.Join(certsDir, f)); os.IsNotExist(err) {
			return fmt.Errorf("could not locate the certificates needed to run EventStoreDB and the tests. Please run 'docker-compose up' for generating the certificates")
		}
	}
	return nil
}

func getCertificatesDir() (string, error) {
	rootDir, err := getRootDir()
	if err != nil {
		return "", err
	}
	return path.Join(rootDir, "certs"), nil
}

func getRootDir() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	currentDir = strings.Replace(currentDir, "\\", "/", -1)
	return path.Clean(path.Join(currentDir, "../")), nil
}

func CreateTestClient(container *Container, t *testing.T) *client.Client {
	config, err := client.ParseConnectionString(fmt.Sprintf("esdb://%s?tls=false", container.Endpoint))
	if err != nil {
		t.Fatalf("Unexpected configuration error: %s", err.Error())
	}

	client, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("Unexpected failure setting up test connection: %s", err.Error())
	}
	err = client.Connect()
	if err != nil {
		t.Fatalf("Unexpected failure connecting: %s", err.Error())
	}

	return client
}

func CreateLocalhostTestClient(t *testing.T) *client.Client {
	config, err := client.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		t.Fatalf("Unexpected configuration error: %s", err.Error())
	}

	client, err := client.NewClient(config)
	if err != nil {
		t.Fatalf("Unexpected failure setting up test connection: %s", err.Error())
	}
	err = client.Connect()
	if err != nil {
		t.Fatalf("Unexpected failure connecting: %s", err.Error())
	}

	return client
}