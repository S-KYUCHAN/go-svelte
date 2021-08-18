<script>
    import axios from 'axios'
    import { onMount } from 'svelte';
    import { useParams } from "svelte-navigator";

    const params = useParams();

    $: articles = [];
    
    async function getAPI() {
        return await axios.get(`http://localhost:8000/article/${ $params.id }`)
        .then(response => response.data)
    }
    
    async function deleteAPI() {
        console.log("Delete");
        return await axios.delete(`http://localhost:8000/article/${ $params.id }`)
        .then(response => console.log(response))
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

