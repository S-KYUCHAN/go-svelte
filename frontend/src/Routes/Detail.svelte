<script>
    import axios from 'axios'
    import { onMount } from 'svelte';
    import { useNavigate, useLocation, useParams, Link } from "svelte-navigator";

    const params = useParams();
    const navigate = useNavigate();
    const location = useLocation();

    $: articles = [];
    
    async function getAPI() {
        return await axios.get(`http://localhost:8000/article/${ $params.id }`)
        .then(response => response.data)
    }
    
    async function deleteAPI() {
        console.log("Delete");
        await axios.delete(`http://localhost:8000/article/${ $params.id }`)
        .then(response => console.log(response))
        const from = ($location.state && $location.state.from) || "/";
        navigate(from, { replace: true });
    }

    onMount(async () => {
        const res = await getAPI();
        articles = res;
        console.log(articles);
    });
</script>

{#each articles as article}
<h3>{article.title}</h3>
<div class="box">
    <div>{article.content}</div>	
</div>

<button on:click={deleteAPI}>Delete</button>

{/each}

