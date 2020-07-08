<script>
	import router from "page";
	import SignIn from "./SignIn.svelte";
	import Home from "./Home.svelte";
	import Error from "./Error.svelte";
	let page;
	let params;

	router("/", () => page = SignIn)
	router("/home", () => page = Home)
	router("/*", () => page = Error)
	router.start();
	function handleClick() {
		alert('no more alerts')
	}
</script>
<svelte:component this={page} params={params}></svelte:component>

<main>
	<h1>Glance!!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn how to build Svelte apps.</p>
	<button class="ui button" on:click|once={handleClick}>
		Follow
	</button>
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

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>