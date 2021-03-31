<script>

  import { onMount } from "svelte";
  import ProjectInfo from "./ProjectInfo.svelte";

  export let projectsList = [];

  onMount(async () => {
    await fetch('http://localhost:8080/v1/projects')
      .then(res => res.json())
      .then(data => {
        projectsList = data;
      })
  })
</script>

<main>
  <p>Projects</p>

  {#if projectsList.projects}
    {#each projectsList.projects as project}
    <ul>
      <li>
        {project.shortName}
      </li>
    </ul>
    {/each}
  {:else}
  <p>loading...</p>
  {/if}
</main>
