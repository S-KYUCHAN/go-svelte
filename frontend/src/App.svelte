<script>
	// export let name;
	import axios from 'axios'
	import { onMount } from 'svelte';

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

<main>
	<h1>Hello kyuchan!</h1>
	<h2>Let's start SVELTE!</h2>
	<ul class="articleList">
		{#each articles as article}
			<div class="box">
				<div>{article.title}</div>
				<div>{article.content}</div>	
			</div>
		{/each}
	</ul>
	<!-- <p>{data}</p> -->
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em; 
		font-weight: 100;
	}

	h2 {
		color: orange;
		font-size: 2rem;
		font-weight: 100;
	}

	.box {
		/* position: relative; */
		display: inline-block;
		width: 50%;
		border: 1px solid #aaa;
		border-radius: 2px;
		box-shadow: 2px 2px 8px rgba(0,0,0,0.1);
		padding: 1em;
		margin: 2rem 0 2em 0;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>