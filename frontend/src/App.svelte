<script>
	import { Router, Route, Link } from "svelte-navigator";
	import List from "./Routes/List.svelte";	
	import Detail from "./Routes/Detail.svelte"
	import Create from "./Routes/Create.svelte";
</script>

<Router>
	<header>
		<h1>Go-SVELTE</h1>
		<nav>
			<Link to="/">Home</Link>
		</nav>
	</header>

	<main>
		<Route path="/">
			<h3>Article Lists</h3>
			<Link to="/create">
				<button>New article</button>
			</Link>
			<List />
		</Route>
		<Route path="/:id">
			<Detail />
		</Route>
		<Route path="/create">
			<Create />
		</Route>
	</main>
</Router>

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

	h3 {
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