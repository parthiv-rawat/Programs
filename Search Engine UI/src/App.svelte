<script>
	import HomePage from "./HomePage.svelte";
	import Results from "./Results.svelte";

	let resultPage = false;
	let searchResults = [];
	let field = "";
	let signal = false;
	let searchQuery = "";

	async function handleSubmit(event) {
		event.preventDefault();
		searchQuery = event.target.searchInput.value;
		// console.log(searchQuery);
		field = "";
		signal = true;

		try {
			const response1 = await fetch(`http://localhost:8081/search?query=${searchQuery}`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(searchQuery),
			});
			// searchResults = await response.json();
		} catch (error) {
			console.error("Error sending data to Golang:", error);
		}

		try {
		  const response = await fetch(`http://localhost:8081/search?query=${searchQuery}`)
		  searchResults = await response.json();
		} catch (error) {
		  console.error('Error fetching data:', error);
		  searchResults = [];
		}
		resultPage = true;
		console.log(searchResults);
	}

</script>

{#if !resultPage}
	<HomePage  handleSubmit={handleSubmit} />
{:else}
	<Results handleSubmit={handleSubmit} {searchResults} />
{/if}
