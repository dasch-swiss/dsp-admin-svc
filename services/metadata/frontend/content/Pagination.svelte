<script lang="ts">
  import type { PaginationData } from "./interfaces";
  import { getProjects, pagedResults } from "./stores";

  export let pagination = {} as PaginationData;
  let currentPage = 1;
  const baseResultsRange = [1, 9];
  let currentResults = baseResultsRange;

  let handlePagination = (event: MouseEvent) => {
    const id = (event.target as HTMLElement).id;
    if (currentPage === Number(id)) {
      return;
    } else if (id === 'first') {
      currentPage = 1;
    } else if (id === 'last') {
      currentPage = pagination.totalPages;
    } else {
      currentPage = Number(id);
    }
    
    document.querySelector('.active').classList.remove('active');
    document.getElementById((currentPage).toString()).classList.add('active');
    getProjects(currentPage);
    currentResults = baseResultsRange.map(v => v + ((currentPage - 1) * baseResultsRange[1]));
  }
</script>

<div class={pagedResults ? 'pagination-container' : 'hidden'}>
  <div class="stats">
    <div>
      <p>
        Showing
        <span>{currentResults[0]}</span>
        to
        <span>{currentResults[1] > pagination.totalCount ? pagination.totalCount : currentResults[1]}</span>
        of
        <span>{pagination.totalCount}</span>
        results
      </p>
    </div>
  </div>
  <div class="pagination">
    <button on:click={handlePagination} id="first" title="First Page" disabled={currentPage === 1}>&laquo;</button>
    {#each Array(pagination.totalPages) as _, i}
      <button on:click={handlePagination} id={(i + 1).toString()} class={i === 0 ? 'active' : ''}>{i + 1}</button>
    {/each}
    <button on:click={handlePagination} id="last" title="Last Page" disabled={currentPage === pagination.totalPages}>&raquo;</button>
  </div>
</div>

<style>
  .pagination-container {
    display: flex;
    justify-content: center;
    align-items: flex-start;
  }
  .pagination {
    display: inline-block;
    margin: 0 20px;
  }
  button {
    color: black;
    background-color: #fff;
    float: left;
    padding: 8px 16px;
    text-decoration: none;
    border: 1px solid #ddd;
  }
  button.active {
    background-color: var(--dasch-violet);
    color: white;
    border: 1px solid var(--dasch-violet);
  }
  button:hover:not(.active), button:hover:not:disabled {
    background-color: var(--dasch-light-violet);
  }
  button:first-child {
    border-top-left-radius: 5px;
    border-bottom-left-radius: 5px;
  }
  button:last-child {
    border-top-right-radius: 5px;
    border-bottom-right-radius: 5px;
  }
</style>
