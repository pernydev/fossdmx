<script lang="ts">
	import ErrorDialog from '$lib/component/ErrorDialog.svelte';
import '../app.css';
</script>

<c>
	<nav aria-label="toolbar">
        <div class="options">
            <button>FILE</button>
            <button>EDIT</button>
            <button>VIEW</button>
            <button>HELP</button>
        </div>
        <div class="project">
            <p aria-label="project">My Lightshow</p>
        </div>
	</nav>

	<slot></slot>
</c>

<ErrorDialog />

<style>
    c {
        display: flex;
        flex-direction: column;
        height: 100vh;
    }
    
	nav {
        width: 100%;
		display: flex;
        justify-content: space-between;
        background-color: var(--background-900);
	}

    .options {
        display: flex;
    }
    
    nav button {
        all: unset;
        cursor: pointer;
        padding: .7em 1.2em;
        background-color: var(--background-900);
        font-weight: bold;
        text-transform: capitalize;
        
    }

    nav button:hover {
        background-color: var(--background-800);
        color: white;
    }

    .project {
        display: flex;
        align-items: center;
        padding: 0 1em;
    }
</style>
