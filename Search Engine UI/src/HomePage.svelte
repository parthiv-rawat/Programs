<script>
    import { fade,fly,scale } from 'svelte/transition'
    import { elasticInOut } from "svelte/easing";
    import { tweened } from "svelte/motion";

    let field = "";
    let isVisible = true;
    export let handleSubmit;

    function toggleVisibility() {
    isVisible = !isVisible;
  }
</script>

<header>
    {#if isVisible}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <img src="images/image.png" on:click={toggleVisibility} transition:scale alt="Logo" class="logo" />
    {/if}
    <h1>FlyHigh Edutech Search Engine</h1>
</header>

<main>
    <form on:submit={handleSubmit}>
        <input
            type="text"
            id="search-input"
            name="searchInput"
            placeholder="Enter your search query..."
            bind:value={field}
        />
        <button type="submit" on:click
            ><img src="images/search.png" alt="" /></button
        >
    </form>
</main>

<footer>
    <p>© 2023 Search Engine. All rights reserved.</p>
</footer>

<style>
    img {
        display: inline;
        height: 100px;
    }
    header {
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: #e9ab39;
        color: #fff;
        padding: 5px;
        text-align: center;
    }

    h1 {
        display: inline;
        margin: 5px;
    }

    main {
        display: flex;
        flex-direction: column;
        height: 69vh;

        padding: 20px;
        text-align: center;
    }

    footer {
        margin-bottom: 0;
    }
    footer {
        display: flex;
        justify-content: center;

        align-items: center;
        width: 98.5%;
        height: 20px;
        background-color: #e9ab39;
        color: #fff;
        padding: 10px;
        text-align: center;
    }

    /* --------------------------------- */

    form {
        display: flex;
        justify-content: center;
        margin-top: 30vh;
    }

    input[type="text"] {
        padding: 10px;
        width: 1000px;
        border-radius: 50px;
        font-size: 16px;
        border: 1px solid #ccc;
        box-shadow: 3px 3px 4px rgba(29, 28, 28, 0.1);
    }

    button[type="submit"] {
        margin-left: 20px;
        padding: 10px 25px;
        font-size: 16px;
        border-radius: 40px;
        background-color: #e9ab39;
        color: #fff;
        border: none;
        cursor: pointer;
    }

    button img {
        height: 30px;
        padding: 2px;
    }
</style>
