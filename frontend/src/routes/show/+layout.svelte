<script lang="ts">
	import { goto } from '$app/navigation';
	import { fixture_modes, fixture_types, fixtures, scenes, sequences } from '$lib/main.svelte';

	async function fetchShowFile() {
		const sf = await fetch('/api/showfile');
		if (!sf.ok) {
			goto('/welcome');
		}

		const showfile = (await sf.json()) as Showfile;

		Object.keys(fixtures).forEach((key) => delete fixtures[key]);
		Object.assign(fixtures, showfile.fixtures);

		Object.keys(scenes).forEach((key) => delete scenes[key]);
		Object.assign(scenes, showfile.scenes);

		Object.keys(sequences).forEach((key) => delete sequences[key]);
		Object.assign(sequences, showfile.sequences);

		const fixture_types_resp = await fetch('/api/registry/fixture_types');
		const fixture_types_json = (await fixture_types_resp.json()) as Record<string, FixtureType>;
		console.log(fixture_types_json);
		Object.assign(fixture_types, fixture_types_json);
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
