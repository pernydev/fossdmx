<script lang="ts">
	import { goto } from '$app/navigation';
	import { fixture_modes, fixtures } from '$lib/main.svelte';

	async function fetchShowFile() {
		const sf = await fetch('/api/showfile');
		if (!sf.ok) {
			goto('/welcome');
		}

		const showfile = (await sf.json()) as Showfile;

		Object.keys(fixtures).forEach((key) => delete fixtures[key]);
		Object.assign(fixtures, showfile.fixtures);

		const fixture_types = await fetch('/api/registry/fixture_types');
		const fixture_types_json = (await fixture_types.json()) as Record<string, FixtureType>;
		for (const val of Object.values(fixture_types_json)) {
			for (const mode of val.modes) {
				fixture_modes[mode.id] = mode;
			}
		}
	}
</script>

{#await fetchShowFile()}
	<p>Loading...</p>
{:then}
	<slot></slot>
{:catch error}
	<p>{error.message}</p>
{/await}
