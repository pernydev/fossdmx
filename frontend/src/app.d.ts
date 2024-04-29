// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}

		type Showfile = {
			name: string;
			filename: string;
			authors: string[];
			fixtures: { [key: string]: Fixture };
			sequences: { [key: string]: Sequence };
			scenes: { [key: string]: Scene };
		};

		type FixtureType = {
			id: string;
			name: string;
			vendor: string;
			modes: FixtureMode[];
			features: FixtureFeature[];
		};

		type FixtureFeature = 'dimmer' | 'strobe' | 'rgb' | 'white' | 'amber' | 'pan' | 'tilt';

		type FixtureMode = {
			id: string;
			name: string;
			channel_count: number;
			channels: { [key: number]: FixtureChannel };
		};

		type FixtureChannel = {
			type: FixtureChannelType;
		};

		type FixtureChannelType =
			| 'custom'
			| 'rgb_red'
			| 'rgb_green'
			| 'rgb_blue'
			| 'rgb_white'
			| 'rgb_amber'
			| 'dimmer'
			| 'strobe'
			| 'pan'
			| 'tilt';

		type Channel = {
			channel: number; // DMX channel
			universe: number; // DMX universe, I think 256 universes is good enough :D
		};

		type Color = {
			red: number;
			green: number;
			blue: number;
			white: number;
			amber: number;
		};

		type Position = {
			pan: number;
			tilt: number;
		};

		type Fixture = {
			id: number;
			channel: Channel;
			type_id: string;
			mode_id: string;
			visualizor_position: XYZ;
			overwrite: FixtureLayer;
		};

		type XYZ = {
			x: number;
			y: number;
			z: number;
		};

		type FixtureOverwrite = {
			dimmer: number;
			strobe: number;
			color: Color;
			position: Position;
		};

		type FixtureLayer = {
			overwrites: FixtureOverwrite;
			weight: number;
		};

		type Scene = {
			id: string;
			overwrites: FixtureOverwrite;
			fixture_ids: string[];
		};

		type Sequence = {
			scene_ids: string[];
			delay: number;
			fadetime: number;
		};
	}
}

export {};
