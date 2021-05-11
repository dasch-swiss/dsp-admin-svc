<script lang="ts">
  import Header from "./Header.svelte";
  import Footer from "./Footer.svelte";
  import Router from "svelte-spa-router";
  import routes from "./routes";
  import { previousRoute } from "./stores";

  // preventing go back button to get back out of domain
  const routeEvent = (event: Event) => {
    previousRoute.set(window.location.href);
  }
</script>

<Header/>

<div class=wrapper>
  <div class=content-container>
    <Router {routes} on:routeEvent={routeEvent} />
  </div>
</div>

<Footer/>

<style>
  .content-container {
    flex-direction: row;
  }
@media screen and (min-width: 992px) {
  .wrapper {
    display: flex;
    flex-direction: row;
    justify-content: center;
    flex: 1 0 auto;
  }
  .content-container {
    display: flex;
    flex-direction: row;
    max-width: 1920px;
  }
}
</style>
