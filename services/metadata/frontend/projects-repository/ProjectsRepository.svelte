<script lang='ts'>
  import Tile from './Tile.svelte';
  import Category from './Category.svelte';
  import { onMount } from 'svelte';
  import Pagination from './Pagination.svelte';
  import { getProjectsMetadata, pagedResults } from '../stores';
  import { fade } from 'svelte/transition';
  import {createEventDispatcher} from 'svelte'

  const dispatch = createEventDispatcher();
  let message = 'Loading...';

  setTimeout(() => {
    const noData = 'No data retrived. Please check the connection and retry.';
    const noProject = 'No projects found.'
      message = $pagedResults && $pagedResults.length ? noData : noProject;
    }, 3000);
  
  onMount(async () => {
    // preventing go back button to get back out of domain
    dispatch('routeEvent');
    
    if (!$pagedResults) {
      await getProjectsMetadata(1);
    }
  });
</script>

<nav>
  <div class="category-container hidden m-inline-block">
    <Category />
  </div>
</nav>
<main in:fade="{{duration: 200}}">
  <div class=tile-container>
    {#if $pagedResults && $pagedResults.length}
      {#each $pagedResults as project}
        <Tile projectMetadata={project}/>
      {/each}
    {:else}
      <p>{message}</p>
    {/if}
  </div>
  {#if $pagedResults && $pagedResults.length}
    <Pagination />
  {/if}
</main>

<style>
* {
  box-sizing: border-box;
}
nav, main {
  width: 100%;
  min-height: auto;
  padding: 10px;
}
nav {
  flex: 0 0 20%;
  display: flex;
  justify-content: flex-end;
  padding: 0;
  /* TODO: temp hidden faceated search */
  display: none;
}
.category-container {
  padding-top: 45px;
  max-width: 210px;
}
main {
  width: 100%;
  align-items: center;
  justify-content: center;
}
.tile-container {
  padding: 10px 5px;
  display: flex;
  flex-flow: row wrap;
  justify-content: center;
  max-width: 1200px;
}
@media screen and (min-width: 992px) {
  nav, main {
    min-height: 950px;
  }
  nav {
    padding: 10px;
  }
  .tile-container {
    padding: 40px 0;
    min-width: 742px;
  }
}
@media screen and (min-width: 1200px) {
  .tile-container {
    min-width: 940px;
  }
}
@media screen and (min-width: 768px) and (max-width: 1023px) { }
@media screen and (min-width: 1024px) and (max-width: 1365px) { }
@media screen and (min-width: 1366px) {}
</style>
