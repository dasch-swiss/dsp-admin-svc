<script>
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";
  import { currentProjectMetadata } from "../stores";

  export let dataset;

  let isAbstractExpanded;
  let abstractLinesNumber;

  const toggleExpand = () => {
    isAbstractExpanded = !isAbstractExpanded;
  };

  const findObjectById = (id) => {
    return $currentProjectMetadata?.metadata.find(obj => obj.id === id);
  };

  onMount(() => {
    const el = document.getElementById('abstract');
    const lineHeight = parseInt(window.getComputedStyle(el).getPropertyValue('line-height'));
    const divHeight = el.scrollHeight;
    abstractLinesNumber = divHeight / lineHeight;
    isAbstractExpanded = abstractLinesNumber > 6 ? false : true;
  });

  const copyToClipboard = () => {
    let text = document.createRange();
    text.selectNode(document.getElementById('how-to-cite'));
    window.getSelection().removeAllRanges();
    window.getSelection().addRange(text);
    document.execCommand('copy');
    window.getSelection().removeAllRanges();
  };

  const truncateString = (s) => {
    const browserWidth = window.innerWidth;
    if (browserWidth < 992 && s.length > ((browserWidth - 100) / 8)) {
      return `${s.substring(0, (browserWidth - 100) / 8)}...`;
    } else if (browserWidth >= 992 && s.length > (browserWidth / 17)) {
      return `${s.substring(0, (browserWidth / 17))}...`;
    } else return s;
  };

  let mergedAttributions = [];
  const attributions = JSON.parse(JSON.stringify(dataset?.content.qualifiedAttribution));
  for(let a of attributions) {
    if(!mergedAttributions.length) {
      mergedAttributions.push(a);
    } else {
      mergedAttributions.push(a);
      for(let b of mergedAttributions) {
        if(a.agent[0].id === b.agent[0].id && a.role !== b.role){
          b.role.push(a.role[0]);
          mergedAttributions.splice(mergedAttributions.indexOf(a) ,1);
        }
      }
    }
  }

  console.log('loaded dataset', dataset)
</script>

<div id=dataset in:fade="{{duration: 200}}">
  {#if dataset}
    {#if dataset?.content.alternativeTitle}
    <div>
      <span class=label>Alternative Title</span>
      <span class=data>{dataset?.content.alternativeTitle}</span>
    </div>
    {/if}
  <div class="grid-wrapper">
    <div>
      <span class=label>Access</span>
      <span class=data>{dataset?.content.conditionsOfAccess}</span>
    </div>
    <div>
      <span class=label>Status</span>
      <span class=data>{dataset?.content.status}</span>
    </div>
    {#if dataset.content.dateCreated}
    <div>
      <span class=label>Date Created</span>
      <span class=data>{dataset?.content.dateCreated}</span>
    </div>
    {/if}
    {#if dataset.content.dateModified}
    <div>
      <span class=label>Date Modified</span>
      <span class=data>{dataset?.content.dateModified}</span>
    </div>
    {/if}
    <div>
      <span class=label>License</span>
      {#if Array.isArray(dataset?.content.license)}
        {#each dataset?.content.license as l}
        <a href={l.url} class=data target=_>CC {(`${l.url.split("/")[4]} ${l.url.split("/")[5]}`).toUpperCase()}</a>
        {/each}
      {/if}
    </div>
    <div>
      <span class=label>Type of Data</span>
      <span class=data>{dataset?.content.typeOfData.join(', ')}</span>
    </div>
    {#if dataset?.content.documentation}
    <div style="grid-column-start: 1;grid-column-end: 3;">
      <span class=label>Additional documentation</span>
      {#if Array.isArray(dataset?.content.documentation)}
        {#each dataset?.content.documentation as d}
          {#if d.url}
          <a class=data href={d.url} target=_>{truncateString(d.name)}</a>
          {:else if d.match("http")}
          <a class=data href={d} target=_>{truncateString(d)}</a>
          {:else}
          <span class=data>{d}</span>
          {/if}
        {/each}
      {/if}
    </div>
    {/if}
  </div>
  <div class="grid-wrapper" style="grid-template-columns: repeat(1, 1fr)">
    <div>
      <span class=label>Languages</span>
      <span class=data>{dataset?.content.language.join(', ')}</span>
    </div>
  </div>

  <div class="property-row">
    <span class=label style="display:inline">
      How To Cite
      <button on:click={copyToClipboard} title="copy text to clpboard">
         <svg class="icon" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"></path></svg>
      </button>
    </span>
    <span id=how-to-cite class=data>{dataset?.content.howToCite}</span>
  </div>

  <div>
    <span class=label>Abstract</span>
    {#if Array.isArray(dataset?.content.abstract)}
    <div id=abstract class="data {isAbstractExpanded ? '' : 'abstract-short'}">
      {#each dataset?.content.abstract as a}
        {#if a.url}
        <div><a class=data href={a.url} target=_>{truncateString(a.name)}</a></div>
        {:else}
        <div>{a}</div>
        {/if}
      {/each}
    </div>
    {/if}
  </div>

  {#if abstractLinesNumber > 6}
  <div on:click={toggleExpand} class=expand-button>show {isAbstractExpanded ? "less" : "more"}</div>
  {/if}

  <span class=label>Attributions</span>
  <div class="grid-wrapper">
    {#if Array.isArray(mergedAttributions)}
      {#each mergedAttributions as a}
        <div class="attributions data">
          <div class=role>{a.role.join(", ")}</div>
          {#if findObjectById(a.agent[0].id).type === "http://ns.dasch.swiss/repository#Person"}
            {#if findObjectById(a.agent[0].id)?.sameAs}
              <a href={findObjectById(a.agent[0].id)?.sameAs[0].url} target=_>{findObjectById(a.agent[0].id)?.givenName.split(";").join(" ")} {findObjectById(a.agent[0].id)?.familyName.split(";").join(" ")}</a>
            {:else}
              <div>{findObjectById(a.agent[0].id)?.givenName.split(";").join(" ")} {findObjectById(a.agent[0].id)?.familyName.split(";").join(" ")}</div>
            {/if}
            {#if Array.isArray(findObjectById(a.agent[0].id)?.memberOf)}
              {#each findObjectById(a.agent[0].id)?.memberOf as o}
                <div>{findObjectById(o.id).name}</div>
              {/each}
            {/if}
            <div>{findObjectById(a.agent[0].id)?.jobTitle[0]}</div>
          {:else}
            <div>{findObjectById(a.agent[0].id)?.name}</div>
          {/if}
          {#if findObjectById(a.agent[0].id)?.email && Array.isArray(findObjectById(a.agent[0].id)?.email)}
            <a class=email href="mailto:{findObjectById(a.agent[0].id)?.email[0]}">{findObjectById(a.agent[0].id)?.email[0]}</a>
          {/if}
        </div>
      {/each}
    {/if}
  </div>

  {/if}
</div>

<style>
  a {color: var(--lead-colour);}
  button {
    border: none;
    background-color: inherit;
    padding: 0;
    position: relative;
    top: 10px;
    color: var(--lead-colour);
    z-index: 0;
  }
  .icon {
    margin: -1rem 0 0.25rem;
  }
  .role {
    color: var(--secondary-colour);
  }
  .abstract-short {
    display: -webkit-box;
    -webkit-line-clamp: 6;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }
  .attributions {
    padding: 10px 10px 0 0;
    line-height: 1.5;
  }
  .grid-wrapper {
    display: grid;
    grid-template-columns: repeat(1, 1fr);
  }
  @media screen and (min-width: 576px) {
    .grid-wrapper {
      grid-template-columns: repeat(2, 1fr);
    }
  }
  @media screen and (min-width: 1200px) {
    .grid-wrapper {
      grid-template-columns: repeat(3, 1fr);
    }
  }
</style>
