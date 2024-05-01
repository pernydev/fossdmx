import { fixture_modes, fixtures } from './main.svelte';

/*
let dragoffset = $state(0);
let whereweat = $state(0);
let mousedown = $state(false);
let dragging = $state(false);

$effect(() => {
	if (1 > whereweat - dragoffset) {
		whereweat = 0;
	}
    //if (!dragging && (fixtureend < whereweat || whereweat < fixturestart)) return;

	if (!dragging) {
		dragoffset = whereweat - fixturestart;
	}
    */

/*
	if (whereweat > 0 && mousedown) {
		fixturestart = whereweat - dragoffset;
		dragging = true;
	}
});
*/

export function calculatePositions(): Record<string, { start: number; end: number }> {
	const positions: Record<string, { start: number; end: number }> = {};
	for (const [id, fixture] of Object.entries(fixtures)) {
		positions[id] = {
			start: fixture.channel.channel,
			end: fixture_modes[fixture.mode_id].channel_count + fixture.channel.channel
		};
	}
	return positions;
}

export function getFixtureAtPosition(
	position: number,
	calculation: Record<string, { start: number; end: number }>
): string | null {
	for (const [id, { start, end }] of Object.entries(calculation)) {
		if (position >= start && position <= end) {
			return id;
		}
	}
	return null;
}

export function getUniverse(
	calculation: Record<string, { start: number; end: number }>
): Record<number, string | null> {
	const universe: Record<number, string | null> = {};
	for (let i = 0; i < 512; i++) {
		universe[i] = getFixtureAtPosition(i, calculation);
	}
	return universe;
}

/**
 * Find the next available channel for a fixture with a given channel count.
 * @param count The channel count of the fixture.
 * @param calculation The current channel usage.
 * @returns The next available channel.
 */
/**
 * Find the next available channel for a fixture with a given channel count.
 * @param count The channel count of the fixture.
 * @param calculation The current channel usage.
 * @returns The next available channel.
 */
export function nextAvailableChannel(
	count: number,
	calculation: Record<string, { start: number; end: number }>
): number {
	let availableChannel = 1;
	const positions = Object.values(calculation);

	positions.sort((a, b) => a.start - b.start);

	for (const position of positions) {
		if (availableChannel + count <= position.start) {
			return availableChannel;
		} else {
			availableChannel = position.end + 1;
		}
	}

	if (availableChannel + count <= 512) {
		return availableChannel;
	}

	return -1;
}
