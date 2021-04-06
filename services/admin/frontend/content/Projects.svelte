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
  <ul class="projects-list">
    {#each projectsList.projects as project}
      <button class="project" on:click="{() => handleProjectClick(project.id)}">
        {project.shortName}
      </button>
    {/each}
  </ul>
  {:else}
  <p>No Projects Found</p>
  {/if}

  {#if currentProject && projectsList.projects}
  <div>
    <div>ID: {currentProject.id}</div>
    <div>Title: {currentProject.shortName}</div>
    <div>Description: {currentProject.description}</div>
  </div>
  {:else}
  <p>Select a project</p>
  {/if}
</main>

<style>

  main {
    width: 400px;
  }

  .projects-list {
    margin: 0px;
  }

  .project {
    border: none;
    display: block;
    width: 100%;
    font-size: 20px;
  }

  .project:nth-child(odd) {
    background-color: rgba(192, 192, 192, 0.803);
  }

  .project:nth-child(odd):hover {
    background-color: silver;
  }

  .project:nth-child(even) {
    background-color: powderblue;
  }

  .project:nth-child(even):hover {
    background-color: rgb(105, 217, 232);
  }
</style>
