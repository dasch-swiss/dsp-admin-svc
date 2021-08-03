<script>
    import {onMount} from 'svelte';
    import {Router, Link} from "svelte-routing";

    onMount(() => {

    });


</script>

<div class="home">
    <div>
        <h1>Admin Services</h1>
    </div>
    <div class="services">
        <Router>
            <Link to={`/projects`} class="regular-link">
                <div class="box">Manage projects</div>
            </Link>
        </Router>
        <Router>
            <Link to={`/`} class="regular-link">
                <div class="box">Manage users</div>
            </Link>
        </Router>
        <Router>
            <Link to={`/`} class="regular-link">
                <div class="box">Manage permissions</div>
            </Link>
        </Router>
    </div>
</div>

<style>
    .home {
        width: 96%;
        padding-left: 2%;
        padding-right: 2%;
    }

    .services {
        width: 100%;
        text-align: center;
    }

    .services .box{
        width: 30%;
        margin: 0% 0.3%;
        border: 2px solid #418dff;
        border-radius: 5px;
        display: inline-block;
        min-height: 100px;
        line-height: 100px;
        font-size: 24px;
        transition: all ease-in-out 0.3s;
    }

    .services .box:hover {
        background-color: #5799fe;
        color: #fff;
        letter-spacing: 5px;
        transition: all ease-in-out 0.3s;
    }
</style>