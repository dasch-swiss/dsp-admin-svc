import {push} from 'svelte-spa-router';
import { writable } from 'svelte/store';
import type { PaginationData } from './interfaces';

export const pagination = writable({} as PaginationData);
export const pagedResults = writable(undefined as any[]);
export const currentProjectMetadata = writable(undefined);
export const query = writable('');
export const previousRoute = writable('');

export async function getProjectsMetadata(page: number, q?: string): Promise<void> {
  // const baseUrl = process.env.BASE_URL;
  const protocol = window.location.protocol;
  const port = protocol === 'https:' ? '' : ':3000';
  const baseUrl = `${protocol}//${window.location.hostname}${port}/`;
  const baseResultsRange = [1, 9];
  let route: string;
  let currentResultsRange = baseResultsRange.map(v => v + ((page - 1) * baseResultsRange[1]));
  
  if (q) {
    query.set(q);
    route = `projects?q=${q}&_page=${page}&_limit=${baseResultsRange[1]}`;
  } else {
    query.set('');
    route = `projects?_page=${page}&_limit=${baseResultsRange[1]}`;
  }

  // console.log(baseUrl, route);
  push(`/${route}`);

  await fetch(`${baseUrl}${route}`)
    .then(r => {
      const totalCount = parseInt(r.headers.get('X-Total-Count'));
      let totalPages = Math.floor(totalCount/baseResultsRange[1]);
      if (!Number.isInteger(totalCount/baseResultsRange[1])) {
        totalPages++;
      };
      // console.log(totalCount, totalPages)
      pagination.set({currentPage: page, currentResultsRange, totalCount, totalPages});
      return r.json();
    })
    .then(data => {pagedResults.set(data), console.log(data)})
}
