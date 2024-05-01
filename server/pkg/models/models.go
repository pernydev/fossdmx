package models

type Showfile struct {
	Name     string              `json:"name"`
	Filename string              `json:"filename"`
	Authors  []string            `json:"authors"`
	Fixtures map[string]*Fixture `json:"fixtures"`

	Sequences map[string]*Sequence `json:"sequences"`
	Scenes    map[string]*Scene    `json:"scenes"`
}

type FixtureType struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Vendor string `json:"vendor"`
	Prefix string `json:"prefix"`

	Modes    []FixtureMode    `json:"modes"`
	Features []FixtureFeature `json:"features"`
}

type FixtureFeature string

const (
	FeatureDimmer FixtureFeature = "dimmer"
	FeatureStrobe FixtureFeature = "strobe"
	FeatureRGB    FixtureFeature = "rgb"
	FeatureW      FixtureFeature = "white"
	FeatureA      FixtureFeature = "amber"
	FeaturePan    FixtureFeature = "pan"
	FeatureTilt   FixtureFeature = "tilt"
)

type FixtureMode struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ChannelCount int    `json:"channel_count"`

	Channels map[uint8]FixtureChannel `json:"channels"`
}

func (m *FixtureMode) Update(fixture *Fixture) {
}

type FixtureChannel struct {
	Type FixtureChannelType `json:"type"`
}

type FixtureChannelType string

const (
	ChannelTypeCustom FixtureChannelType = "custom"

	ChannelTypeRed   FixtureChannelType = "rgb_red"
	ChannelTypeGreen FixtureChannelType = "rgb_green"
	ChannelTypeBlue  FixtureChannelType = "rgb_blue"
	ChannelTypeWhite FixtureChannelType = "rgb_white"
	ChannelTypeAmber FixtureChannelType = "rgb_amber"

	ChannelTypeDim    FixtureChannelType = "dimmer"
	ChannelTypeStrobe FixtureChannelType = "strobe"

	ChannelTypePan  FixtureChannelType = "pan"
	ChannelTypeTilt FixtureChannelType = "tilt"
)

type Channel struct {
	Channel  uint16 `json:"channel"`  // DMX channel
	Universe uint8  `json:"universe"` // DMX universe, I think 256 universes is good enough :D
}

type Color struct {
	Red   uint8 `json:"red"`
	Green uint8 `json:"green"`
	Blue  uint8 `json:"blue"`
	White uint8 `json:"white"`
	Amber uint8 `json:"amber"`
}

type Position struct {
	Pan  uint8 `json:"pan"`
	Tilt uint8 `json:"tilt"`
}

type Fixture struct {
	ID      string  `json:"id"`
	Channel Channel `json:"channel"`

	TypeID string `json:"type_id"`
	ModeID string `json:"mode_id"`

	VisualizorPosition XYZ `json:"visualizor_position"`

	Overwrite FixtureLayer            `json:"preset"`
	Layers    map[string]FixtureLayer `json:"layers"`
}

type XYZ struct {
	X int16 `json:"x"`
	Y int16 `json:"y"`
	Z int16 `json:"z"`
}

type FixtureOverwrite struct {
	Dimmer   uint8    `json:"dimmer"`
	Strobe   uint8    `json:"strobe"`
	Color    Color    `json:"color"`
	Position Position `json:"position"`
}

type FixtureLayer struct {
	Overwrites FixtureOverwrite `json:"overwrites"`
	Weight     uint8            `json:"weight"`
}

type Scene struct {
	ID         string                      `json:"id"`
	Name       string                      `json:"name"`
	Overwrites map[string]FixtureOverwrite `json:"overwrites"` // FixtureID -> FixtureOverwrite
}

type Sequence struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	SceneIDs []string `json:"scene_ids"`
	Delay    uint16   `json:"delay"`
	Fadetime uint16   `json:"fadetime"`
	Loop     bool     `json:"loop"`
}
