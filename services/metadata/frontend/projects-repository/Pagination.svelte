<script lang="ts">
  import { getProjectsMetadata, pagedResults, pagination, query } from '../stores';

  let handlePagination = (event: MouseEvent) => {
    const id = (event.target as HTMLElement).id;
    if ($pagination.currentPage === Number(id)) {
      return;
    } else if (id === 'first') {
      $pagination.currentPage = 1;
    } else if (id === 'last') {
      $pagination.currentPage = $pagination.totalPages;
    } else {
      $pagination.currentPage = Number(id);
    }
    
    document.querySelector('.active').classList.remove('active');
    document.getElementById(($pagination.currentPage).toString()).classList.add('active');
    console.log('curr',$pagination.currentPage);
    getProjectsMetadata($pagination.currentPage, $query);
    window.scrollTo(0,0);
  }
</script>

<div class={pagedResults ? 'pagination-container' : 'hidden'}>
  <div class="stats">
    <div>
      <p>
        Showing
        <span>{$pagination.currentResultsRange[0]}</span>
        to
        <span>{$pagination.currentResultsRange[1] > $pagination.totalCount ? $pagination.totalCount : $pagination.currentResultsRange[1]}</span>
        of
        <span>{$pagination.totalCount}</span>
        results
      </p>
    </div>
  </div>
  <div class="pagination">
    <button on:click={handlePagination} id="first" title="First Page" disabled={$pagination.currentPage === 1}>&laquo;</button>
    {#each Array($pagination.totalPages) as _, i}
      <button on:click={handlePagination} id={(i + 1).toString()} class={i + 1 === $pagination.currentPage ? 'active' : ''}>{i + 1}</button>
    {/each}
    <button on:click={handlePagination} id="last" title="Last Page" disabled={$pagination.currentPage === $pagination.totalPages}>&raquo;</button>
  </div>
</div>

<style>
  .pagination-container {
    display: grid;
    text-align: center;
    grid-template-columns: repeat(1, 1fr);
  }
  .pagination {
    display: inline-block;
    margin: 0 auto;
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
    background-color: var(--lead-colour);
    color: white;
    border: 1px solid var(--lead-colour);
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
  @media screen and (min-width: 768px) {
    .pagination-container {
      /* display: flex;
      justify-content: center;
      align-items: flex-start; */
      text-align: right;
      grid-template-columns: repeat(2, 1fr);
    }
    .pagination {
      margin: 0 20px;
    }
  }
</style>
