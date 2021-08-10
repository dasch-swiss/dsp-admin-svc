<script>
    import {getProject, currentProject, currentUser, getProjects} from "./store";
    import {onMount} from 'svelte';
    import Content from "./Modal/Content.svelte";
    import Modal from 'svelte-simple-modal';

    export let token;
    const projectID = window.location.pathname.split("/")[2];
    const requiredGroup = "Group:" + projectID + ":ProjectAdmin"

    onMount(async () => {
        currentUser.subscribe(async userInfo => {
            if ($currentUser.groups && $currentUser.groups.length !== 0 &&
                ($currentUser.groups.includes(requiredGroup) || $currentUser.groups.includes("Group:SystemAdmin"))) {
                await getProject(userInfo.token, projectID);
            }
        });
    });

</script>

<div class="projects">
    <div>
        <h1>Project Info</h1>
    </div>
    {#if $currentProject.longName !== undefined}
        <div class="info">
            <p>Short Code: {$currentProject.shortCode}</p>
            <p>Short Name: {$currentProject.shortName}</p>
            <p>Long Name: {$currentProject.longName}</p>
            <p>Description: {$currentProject.description}</p>
        </div>
        <!--    Modal for editing a project-->
        <Modal>
            <Content modalType="edit" token="{$currentUser.token}"/>
        </Modal>
    {:else}
        {#if $currentUser.token}
            <div>
                <p>You are not a member of this project.</p>
            </div>
        {:else }
            <div>
                <p>You must be logged in to access this project.</p>
            </div>
        {/if}
    {/if}
</div>
