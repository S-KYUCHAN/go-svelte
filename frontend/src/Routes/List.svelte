<script>
    import axios from 'axios'
    import { onMount } from 'svelte';
    import { useNavigate, useLocation, Link } from "svelte-navigator";

    const navigate = useNavigate();
    const location = useLocation();

    $: articles = [];

	async function getAPI() {
		return await axios.get(`http://localhost:8000/articles/`)
			.then(response => response.data)
	}

	onMount(async () => {
		const res = await getAPI();
		articles = res;
		console.log(articles);
	});
</script>


<ul class="articleList">
    {#each articles as article}
        <div class="box">
            <Link to="/{article.id}"> {article.title}</Link>
            <div>{article.content}</div>	
        </div>
    {/each}
</ul>

<style>
    .box {
		/* position: relative; */
		display: inline-block;
		width: 50%;
		border: 1px solid #aaa;
		border-radius: 2px;
		box-shadow: 2px 2px 8px rgba(0,0,0,0.1);
		padding: 1em;
		margin: 1rem 0 2em 0;
	}
</style>