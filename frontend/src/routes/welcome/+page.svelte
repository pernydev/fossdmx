<script lang="ts">
	import { goto } from "$app/navigation";
	import { err } from "$lib/error.svelte";
	import { onMount } from "svelte";

    async function loadShowfile(filename: string) {
        const response = await fetch('/api/showfile', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                filename: filename
            })
        });
        if (response.ok) {
            goto('/');
        } else {
            const json = await response.json();
            err(json.error, "LOAD_SHOWFILE_FAILED");
        }
    }
</script>

<article>
	<h1>Welcome to FOSS DMX</h1>
	<p>A fully accessable, free and open-source software for controlling DMX lighting fixtures.</p>

	<h2>Showfiles</h2>
	<ul aria-label="showfiles">
		<li>
			<h3>Demo Show</h3>
            <button on:click={() => loadShowfile('demo.json')}>Load</button>
		</li>
	</ul>
</article>

<style>
	article {
		max-width: 1200px;
		margin: 0 auto;

        padding-inline: 1rem;
	}

	h1 {
		margin-bottom: 0.4em;
	}

    ul {
        list-style-type: none;
        padding: 0;

        display: grid;
        grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
        gap: 1em;
    }

    li {
        padding: 1em;
        background-color: var(--background-900);
        border-radius: 0.5em;
        position: relative;
    }

    h3 {
        margin-top: 0;
        margin-bottom: 0.4em;
    }

    li button::after {
        position: absolute;
        inset: 0;
        background: transparent;
    }
</style>
