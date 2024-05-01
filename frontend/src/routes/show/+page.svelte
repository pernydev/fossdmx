<script lang="ts">
	import AttributeModifier from '$lib/component/AttributeModifier.svelte';
	import { scenes } from '$lib/main.svelte';

	type Selection = {
		type: 'scene' | 'sequence' | 'fixture';
		id: string;
	};
	let selected: Selection | null = $state(null);
	let attributes: FixtureOverwrite = $state({});
	let features: FixtureFeature[] = $state([]);

</script>

<div class="row">
	<div aria-label="library" class="padded" style="min-width: 15rem;">
		<div style="display: flex; justify-content: space-between;">
			<p>Library</p>
			<button>NEW</button>
		</div>
		<ul class="library">
			{#each Object.values(scenes) as i}
				<li>
					<button class="unstyled" onclick={() => (selected = { type: 'scene', id: i.id })}>
						{i.name}
					</button>
				</li>
			{/each}
		</ul>
	</div>
	<div class="center f1" aria-label="visualizer, non accessable">
		<span>Visualizer not implemented</span>
	</div>
</div>
<div class={`padded ${selected ? '' : 'center'}`} style="min-height: 10rem;">
	{#if selected}
		{#if selected.type === 'scene'}
			<div>
				<div style="display: flex; flex-direction: column; gap: 0.5rem; max-width: 15rem;">
					<label for="scene-name">Name</label>
					<input type="text" value={scenes[selected.id].name} id="scene-name" />
					<span>{''}</span><span>{''}</span>
				</div>
				<AttributeModifier bind:attributes bind:features />
			</div>
		{:else if selected.type === 'sequence'}
			{selected.id}
		{:else if selected.type === 'fixture'}
			{selected.id}
		{/if}
	{:else}
		<span>Select a scene, sequence, or fixture to view</span>
	{/if}
</div>

<style>
	.row {
		flex: 1;
		display: flex;
		flex-direction: row;
	}

	.row > div {
		border: 0.5px solid var(--primary-900);
	}

	.row > div.f1 {
		flex: 1;
	}

	.padded {
		padding: 1rem;
	}

	.center {
		display: flex;
		justify-content: center;
		align-items: center;
	}

	.library {
		list-style-type: none;
		padding: 0;
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
	}

	.library li button {
		display: block;
		width: calc(100% - 1rem);
		background-color: var(--background-800);
		border-radius: 5px;
		padding: 0.5rem;
	}
</style>
