/*
 *  Copyright 2021 Data and Service Center for the Humanities - DaSCH.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package project

import "errors"

//ErrProjectNotFound project not found
var ErrProjectNotFound = errors.New("no project found with the provided uuid")

//ErrNoProjectDataReturned no project data returned
var ErrNoProjectDataReturned = errors.New("no project data was returned")

//ErrServerNotResponding server not responding
var ErrServerNotResponding = errors.New("the server is not responding")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("invalid entity")

//ErrInvalidUUID invalid UUID
var ErrInvalidUUID = errors.New("invalid uuid provided")

//ErrNoPropertiesChanged invalid update values
var ErrNoPropertiesChanged = errors.New("no new value for any property provided")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

//ErrProjectHasBeenDeleted project has been marked as deleted
var ErrProjectHasBeenDeleted = errors.New("project has been marked as deleted")

//ErrShortCodeAlreadyExists provided short code already exists
var ErrShortCodeAlreadyExists = errors.New("provided short code already exists")

//ErrUserDoesNotHaveCreateProjectsPermission user does not have permission to create projects
var ErrUserDoesNotHaveCreateProjectsPermission = errors.New("user does not have permission to create projects")

//ErrUserDoesNotHaveProjectPermission user does not have permission to view this project
var ErrUserDoesNotHaveReadProjectPermission = errors.New("user does not have permission to view this project")

//ErrUserDoesNotHaveReadAllProjectsPermission user does not have permission to view all projects
var ErrUserDoesNotHaveReadAllProjectsPermission = errors.New("user does not have permission to view all projects")

//ErrUserDoesNotHaveUpdateProjectPermission user does not have permission to update projects
var ErrUserDoesNotHaveUpdateProjectPermission = errors.New("user does not have permission to update projects")

//ErrUserDoesNotHaveDeleteProjectPermission user does not have permission to delete projects
var ErrUserDoesNotHaveDeleteProjectPermission = errors.New("user does not have permission to delete projects")
