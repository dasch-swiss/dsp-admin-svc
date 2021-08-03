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
import { writable } from 'svelte/store'
import type { Project, User} from './interfaces';

export const projectsList = writable([] as Project[]);
export const currentProject = writable({} as Project);
export const currentUser = writable({} as User);

const protocol = window.location.protocol;
const port = protocol === 'https:' ? '' : ':8080';
const baseUrl = `${protocol}//${window.location.hostname}${port}/`;

export async function getProjects(jwt: string, returnDeletedProjects?: boolean): Promise<void> {

  const response = await fetch(`${baseUrl}v1/projects`, {
    headers: {'Authorization': 'Bearer ' + jwt}
  });
  
  response.json().then(res => {
      projectsList.set(res);
  });
}

export async function getProject(jwt: string, uuid: string): Promise<void> {

  const response = await fetch(`${baseUrl}v1/projects/${uuid}`, {
    headers: {'Authorization': 'Bearer ' + jwt}
  });
  
  response.json().then(res => {
      currentProject.set(res);
  });
}

export async function createProject(jwt: string, sc: string, sn: string, ln: string, desc: string): Promise<void> {

  const p = {
    shortCode: sc,
    shortName: sn,
    longName: ln,
    description: desc
  }

  const response = await fetch(`${baseUrl}v1/projects`, {
    method: 'POST',
    headers: {'Authorization': 'Bearer ' + jwt},
    body: JSON.stringify(p)
  });

  response.json().then(() => {
    getProjects(jwt);
  })
}

export async function editProject(jwt: string, uuid: string, sc: string, sn: string, ln: string, desc: string): Promise<void> {

  const p = {
    shortCode: sc,
    shortName: sn,
    longName: ln,
    description: desc
  }

  const response = await fetch(`${baseUrl}v1/projects/${uuid}`, {
    method: 'PUT',
    headers: {'Authorization': 'Bearer ' + jwt},
    body: JSON.stringify(p)
  });

  response.json().then(() => {
    getProject(jwt, uuid);
  })
}

export async function deleteProject(jwt: string, uuid: string): Promise<void> {

  const response = await fetch(`${baseUrl}v1/projects/${uuid}`, {
    method: 'DELETE',
    headers: {'Authorization': 'Bearer ' + jwt}
  });

  response.json().then(() => {
    getProjects(jwt);
  });
}

export function getUser(user: User) {
  currentUser.set(user);
}