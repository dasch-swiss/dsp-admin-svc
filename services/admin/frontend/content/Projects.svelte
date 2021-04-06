<script>
  import { onMount } from "svelte";

  let projectsList = [];
  let currentProject;

  onMount(async () => {
    await fetch('http://localhost:8080/v1/projects')
      .then(res => res.json())
      .then(data => {
        projectsList = data;
      })
  })

  function handleProjectClick(id) {
    currentProject = projectsList.projects.filter(p => p.id.indexOf(id) !== -1)[0];
  }
</script>

<main>
  <p>Projects</p>

  {#if projectsList.projects}
    {#each projectsList.projects as project}
    <ul>
      <li on:click="{() => handleProjectClick(project.id)}">
        {project.shortName}
      </li>
    </ul>
    {/each}
  {:else}
  <p>loading...</p>
  {/if}

  {#if currentProject}
  <div>
    <div>ID: {currentProject.id}</div>
    <div>Title: {currentProject.shortName}</div>
    <div>Description:</div>
    <p>{currentProject.description}</p>
  </div>
  {:else}
  <p>Select a project</p>
  {/if}
</main>
