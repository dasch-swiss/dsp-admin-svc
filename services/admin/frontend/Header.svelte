<script>
    import {currentUser} from "./store";
    import {onMount} from 'svelte';
    import KeyCloak from 'keycloak-js';

    let kc = new KeyCloak("/keycloak.json");

    let logged_in = null;

    onMount(() => {
        kc.init({onLoad: "check-sso", enableLogging: true}).then((auth) => {
            logged_in = auth;
            if (auth) {
                logged_in = true;

                kc.loadUserInfo().then((user) => {
                    user.token = kc.idToken;

                    currentUser.set(user);
                });
            }
        })
    });


</script>

<div class="header">
    <pre>{JSON.stringify($currentUser, null, 2)}</pre>
<!--    <div><p>{JSON.stringify($currentUser.token, null, 2)}</p></div>-->
    <div class="login-logout">
        {#if logged_in && $currentUser.preferred_username}
            <div>
                {$currentUser.preferred_username}
                <button on:click={() => { kc.logout(); }}>
                    Logout
                </button>
            </div>
        {/if}

        {#if logged_in == false}
            <div>
                <button on:click={() => { kc.login(); }}>
                    Login
                </button>
            </div>
        {/if}
    </div>
</div>


<style lang="scss">
  .header {
    width: 96%;
    padding-left: 2%;
    padding-right: 2%;
  }

  .header .login-logout {
    float: right;
  }

  .header .login-logout button{
    color: white;
    background: #1e90ff;
    border: 1px #1e90ff solid;
    border-radius: 5px;
    padding: 5px 10px;
  }

  .header .login-logout button:hover{
    background: #027cf4;
  }

</style>
