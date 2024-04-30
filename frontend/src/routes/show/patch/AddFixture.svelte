<script lang="ts">
	import Dialog from '$lib/component/Dialog.svelte';
	import { fixture_types } from '$lib/main.svelte';
	import { onMount } from 'svelte';
    let { open = $bindable() } = $props();

	let dialog: HTMLDialogElement;
    let vendor = $state('');
    let vendors: string[] = $derived(getVendors(fixture_types));

    $effect(() => {
        if (open) {
            dialog.showModal();
            return;
        }
        dialog.close();
    });

    function getVendors(ftypes: Record<string, FixtureType>): string[] {
        console.log(fixture_types);
        const vendors: string[] = [];
        for (let fixturetype of Object.values(ftypes)) {
            console.log(fixturetype);
            if (!vendors.includes(fixturetype.vendor)) {
                vendors.push(fixturetype.vendor);
            }
        }   
        return vendors;
    }
</script>

<Dialog bind:dialog>
	<h2>Add Fixture</h2>
	<div class="lights">
		<ul>
			{#each vendors as v}
				<!-- svelte-ignore a11y_role_supports_aria_props_implicit -->
				<li data-selected={vendor === v}>
					<button onclick={() => vendor = v}>{v}</button>
				</li>
			{/each}
		</ul>
		{#if vendor !== ''}
			<ul>
                {#each Object.values(fixture_types).filter(ft => ft.vendor === vendor) as ft}
                    <li>
                        <button>{ft.name}</button>
                    </li>
                {/each}
			</ul>
		{:else}
			<div class="select">
				<p>Select a vendor</p>
			</div>
		{/if}
	</div>
</Dialog>7

<style>
	h2 {
		margin-top: 0;
	}

	.lights {
		display: flex;
		justify-content: space-between;

		max-height: 300px;

		gap: 1rem;
	}

	.lights ul {
		list-style: none;
		padding: 0;
		margin: 0;

		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.lights li {
		display: block;
		overflow: hidden;
	}

	.lights > * {
		flex: 1;

		min-width: 200px;
	}

	.lights button {
		display: block;
		width: calc(100% - 1rem);

		padding: 0.5rem;
		background-color: var(--background-800);
		border-radius: 5px;
	}

	.select {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 100%;
		text-align: center;
	}
</style>
