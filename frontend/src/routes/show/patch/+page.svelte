<script lang="ts">
	import { err } from '$lib/error.svelte';
	import { fixture_modes, fixtures } from '$lib/main.svelte';
	import { calculatePositions, getUniverse } from '$lib/patch.svelte';
	import { onMount } from 'svelte';
	import AddFixture from './AddFixture.svelte';

	let selectedFixture: string | null = $state(null);
	let selectedTouchFixture: string | null = $state(null);
	let offset = $state(0);
	let whereweat = $state(0);
	let mousedown = $state(false);
	let dragging = $state(false);
	let locations = $state(calculatePositions());
	let universeLocations = $state(getUniverse(locations));
	
	let dialogOpen = $state(false);

	$effect(() => {
		if (!mousedown) return;
		if (!dragging) {
			selectedFixture = universeLocations[whereweat] || '';
			offset = whereweat - locations[selectedFixture].start;
			dragging = true;
		}

		if (selectedFixture) {
			const newStart = whereweat - offset;
			locations[selectedFixture].start = newStart;
			locations[selectedFixture].end = newStart + fixture_modes[fixtures[selectedFixture].mode_id].channel_count;

			universeLocations = getUniverse(locations);
		}
	});

	function mouseDown() {
		mousedown = true;
	}

	function mouseUp() {
		mousedown = false;
		if (dragging) {
			sendPatch();
		}
		dragging = false;
	}

	async function sendPatch() {
		if (!selectedFixture) return;
		console.log(selectedFixture);
		const response = await fetch('/api/patch', {
			method: 'PATCH',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				fixture_id: selectedFixture,
				channel: {
					channel: locations[selectedFixture].start,
					universe: fixtures[selectedFixture].channel.universe
				}
			})
		});

		if (!response.ok) {
			err('Patching failed!', '/api/patch NOT OK');
		}
	}

	function touch(event: TouchEvent) {
		const elm = document.elementFromPoint(
			event.changedTouches[0].clientX,
			event.changedTouches[0].clientY
		);
		if (elm instanceof HTMLElement) {
			if (!elm.dataset.channel) return;
			if (selectedFixture) {
				locations[selectedTouchFixture || ''].start = parseInt(elm.dataset.channel);
				locations[selectedTouchFixture || ''].end = parseInt(elm.dataset.channel) + 6;
				universeLocations = getUniverse(locations);
				mouseUp();
				return;
			}
			selectedTouchFixture = universeLocations[parseInt(elm.dataset.channel)] || '';
		}
	}

	onMount(() => {
		window.addEventListener('mousedown', mouseDown);
		window.addEventListener('mouseup', mouseUp);
		window.addEventListener('touchend', touch);

		return () => {
			window.removeEventListener('mousedown', mouseDown);
			window.removeEventListener('mouseup', mouseUp);
			window.removeEventListener('touchend', touch);
		};
	});
</script>

<button onclick={() => (dialogOpen = true)}>Add Fixture</button>
<AddFixture bind:open={dialogOpen} />

<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	class="patch"
	aria-details="This element is sadly not accessible. Please enable accessibility mode to manage fixtures."
	onmouseleave={() => (whereweat = 0)}
>
	{#each Array.from({ length: 512 }, (_, i) => i + 1) as i}
		{#if universeLocations[i]}
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<!-- svelte-ignore a11y_mouse_events_have_key_events -->
			<div
				class="fixture"
				data-position={locations[universeLocations[i] || '']?.start === i
					? 'start'
					: locations[universeLocations[i] || '']?.end === i
						? 'end'
						: 'middle'}
				onmouseover={() => (whereweat = i)}
				data-channel={i}
			>
				<span><small>{i}</small></span>
				{#if locations[universeLocations[i] || '']?.start === i}
					<p>{universeLocations[i]}</p>
				{/if}
			</div>
		{:else}
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<!-- svelte-ignore a11y_mouse_events_have_key_events -->
			<div onmouseover={() => (whereweat = i)} data-channel={i}>
				<span><small>{i}</small></span>
			</div>
		{/if}
	{/each}
</div>

<style>
	.patch {
		display: flex;
		flex-wrap: wrap;
		gap: 0.2rem;
		padding: 1rem;
	}

	.patch div {
		display: flex;
		flex-direction: column;
		justify-content: start;
		align-items: start;
		background-color: var(--background-800);
		border-radius: 5px;
		padding-inline: 0.5rem;
		width: 3rem;
		height: 3rem;
	}

	.patch div span {
		margin: 0;
	}

	.patch div p {
		margin: 0;
		font-family: monospace;
	}

	.patch div * {
		user-select: none;
	}

	div.fixture {
		background-color: var(--primary-600);
	}

	div.fixture[data-position='middle'] {
		border-radius: 0;
		/* a little hack but hey, we are using CSS :). Box shadow 0.25 rem on both sides */
		box-shadow:
			0.25rem 0 0 var(--primary-600),
			-0.25rem 0 0 var(--primary-600);
	}

	div.fixture[data-position='end'] {
		border-radius: 0 5px 5px 0;
	}

	div.fixture[data-position='start'] {
		border-radius: 5px 0 0 5px;
	}
</style>
