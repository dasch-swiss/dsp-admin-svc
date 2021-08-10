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

export interface Project {
    id: string;
    shortCode: string;
    shortName: string;
    longName: string;
    description: string;
    createdAt: string;
    createdBy: string;
    changedAt: string;
    changedBy: string;
    deletedAt: string;
    deletedBy: string;
}

export interface User {
    sub: string;
    email_verified: boolean;
    lastName: string;
    jwt: string;
    name: string;
    preferred_username: string;
    given_name: string;
    family_name: string;
    email: string;
    token: string;
    roles: string[];
    groups: string[];
}
