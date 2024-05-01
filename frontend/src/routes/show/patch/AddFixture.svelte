<script lang="ts">
	import Dialog from '$lib/component/Dialog.svelte';
	import { err } from '$lib/error.svelte';
	import { fixture_modes, fixture_types, fixtures } from '$lib/main.svelte';
	import { calculatePositions, nextAvailableChannel } from '$lib/patch.svelte';
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

	async function createFixture(type_id: string, mode_id: string) {
		const fixture: Fixture = {
			id: fixture_types[type_id].prefix + "_" + Math.random().toString(36).substr(2, 9),
			channel: {
				universe: 1,
				channel: nextAvailableChannel(fixture_modes[mode_id].channel_count, calculatePositions())
			},
			mode_id,
			type_id,
			overwrite: {
				dimmer: 0,
				color: {
					red: 0,
					green: 0,
					blue: 0,
					white: 0,
					amber: 0
				},
				strobe: 0,
				position: {
					pan: 0,
					tilt: 0
				}
			},
			layers: {}
		};
		const response = await fetch('/api/patch', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ fixture })
		});

		if (!response.ok) {
			err('Fixture creation failed!', '/api/patch NOT OK');
			return;
		}

		const data = await response.json();
		fixtures[fixture.id] = data;
		open = false;
	}
</script>

<Dialog bind:dialog>
	<h2>Add Fixture</h2>
	<div class="lights">
		<ul>
			{#each vendors as v}
				<!-- svelte-ignore a11y_role_supports_aria_props_implicit -->
				<li data-selected={vendor === v}>
					<button onclick={() => (vendor = v)}>{v}</button>
				</li>
			{/each}
		</ul>
		{#if vendor !== ''}
			<ul>
				{#each Object.values(fixture_types).filter((ft) => ft.vendor === vendor) as ft}
					<li>
						<button onclick={() => createFixture(ft.id, ft.modes[0].id)}>{ft.name}</button>
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
