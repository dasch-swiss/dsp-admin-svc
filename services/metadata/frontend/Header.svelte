<script lang="ts">
  import Category from "./content/Category.svelte";
  import { getProjects } from "./content/stores";

  let showSearchbar = false;
  let showFilters = false;
  let showMenu = false;
  let enteredString = '';

  function toggleSearchbar() {
    showSearchbar = !showSearchbar;
    showFilters = false;
    showMenu = false;
  }

  function toggleFilters() {
    showFilters = !showFilters;
    showSearchbar = false;
    showMenu = false;
  }

  function toggleMenu() {
    showMenu = !showMenu;
    showFilters = false;
    showSearchbar = false;
  }

  let search = (e: Event) => {
    const q = (e.target as HTMLInputElement).value;
    getProjects(1, q);
    enteredString = '';
  }
</script>

<header>
  <div class="header-container">
    <a href="/" class="header-left">
      <img class="logo s-inline-block" src="assets/logo/DaSCH_Logo_sw.svg" alt="DaSCH logo" />
      <img class="icon-logo s-hidden" src="assets/icon/dasch-icon-black.svg" alt="DaSCH logo" />
      <h1 class="title">Repository Explorer</h1>
    </a>
    <div class="header-right">
      <input on:change={search} bind:value={enteredString} class="searchbar-in-header xs-inline-block" type="text" name="searchbar" placeholder="search..." />
      <!-- searchbar button -->
      <button class="xs-hidden" on:click="{toggleSearchbar}">
        <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
        </svg>
      </button>
      <!-- filter button -->
      <button class="m-hidden" on:click="{toggleFilters}">
        <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4" />
        </svg>
      </button>
      <!-- menu button -->
      <button on:click="{toggleMenu}">
        <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16" />
        </svg>
      </button>
    </div>
  </div>
  <div class="searchbar-container xs-hidden" class:hidden={!showSearchbar}>
    <input on:change={search} bind:value={enteredString} class="searchbar" name="searchbar" placeholder="search..." />
  </div>
  <div class="filter-container m-hidden" class:hidden={!showFilters}>
    <Category/>
  </div>
  <div class="menu" class:hidden={!showMenu}>
    <a class="menu-item" href="/">dasch.swiss</a>
    <a class="menu-item" href="/">app.dasch.swiss</a>
    <a class="menu-item" href="/">admin.dasch.swiss</a>
  </div>
</header>

<style>
  header {
    background-color: var(--cl-background);
    position: sticky;
    top: 0px;
  }
  .header-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  .header-left {
    padding: 0px;
  }
  .header-right {
    padding: 0.5rem 0rem;
    margin-right: 0.25rem;
  }
  .logo {
    height: 5rem;
    vertical-align: middle;
    cursor: pointer;
    display: none;
  }
  .icon-logo {
    height: 3rem;
    padding: 5px;
    vertical-align: middle;
    cursor: pointer;
  }
  .title {
    color: var(--dasch-text);
    padding: 12px 0;
    white-space: nowrap;
    font-size: 0.8rem;
    display: inline-block;
  }
  .searchbar-in-header {
    display: none;
    width: 35vw;
    max-width: 400px;
    padding: 0.25rem;
    vertical-align: middle;
    line-height: 1.5rem;
    box-sizing: border-box;
  }
  .searchbar {
    display: block;
    line-height: 1.5rem;
    width: 100%;
    box-sizing: border-box;
  }
  .searchbar-container, .filter-container {
    background-color: var(--cl-background-light);
    padding: 12px;
  }
  button {
    display: inline-block;
    vertical-align: middle;
    border-radius: 0.25rem;
    background-color: inherit;
    border: none;
    padding: 0px;
  }
  button:hover {
    color: var(--dasch-violet);
    background-color: var(--cl-transparent-dark);
  }
  .icon {
    width: 1.5rem;
    height: 1.5rem;
    margin: 0.25rem;
  }
  .menu {
    background-color: var(--cl-background-light);
  }
  .menu-item {
    display: block;
    padding: 1rem 1.5rem;
  }
  .menu-item:hover {
    background-color: var(--cl-transparent-light);
  }
  @media screen and (min-width: 576px) {
    .menu {
      width: 20rem;
      height: 90vh;
      position: absolute;
      right: 0px;
    }
    .title {
      font-size: 1rem;
      padding: 12px;
    }
  }
</style>
