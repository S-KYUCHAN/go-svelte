<script>
	// export let name;
	import axios from 'axios'


	$: data = [];
	$: items = axios.get(`http://localhost:8000/articles/`).then(response => data = response.data)
	console.log(data)
</script>

<main>
	<!-- <h1>Hello {name}!</h1> -->
	<h1>Hello kyuchan!</h1>
	<p>Let's start SVELTE!</p>
	<p>{data}</p>
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

	p {
		color: orange;
		font-size: 2rem;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>