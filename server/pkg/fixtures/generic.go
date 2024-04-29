package fixtures

import "github.com/pernydev/fossdmx/pkg/models"

var GenericPAR = models.FixtureType{
	ID:     "generic.par",
	Name:   "PAR",
	Vendor: "Generic",
	Features: []models.FixtureFeature{
		models.FeatureDimmer,
		models.FeatureRGB,
		models.FeatureStrobe,
	},

	Modes: []models.FixtureMode{
		{
			ID:           "generic.par.3ch",
			Name:         "4 Channel",
			ChannelCount: 4,
			Channels: map[uint8]models.FixtureChannel{
				0: {Type: models.ChannelTypeDim},
				1: {Type: models.ChannelTypeRed},
				2: {Type: models.ChannelTypeGreen},
				3: {Type: models.ChannelTypeBlue},
			},
		},
		{
			ID:           "generic.par.5ch",
			Name:         "5 Channel",
			ChannelCount: 5,
			Channels: map[uint8]models.FixtureChannel{
				0: {Type: models.ChannelTypeDim},
				1: {Type: models.ChannelTypeRed},
				2: {Type: models.ChannelTypeGreen},
				3: {Type: models.ChannelTypeBlue},
				4: {Type: models.ChannelTypeStrobe},
			},
		},
	},
}

var GenericMovingHead = models.FixtureType{
	ID:     "generic.movinghead",
	Name:   "Moving Head",
	Vendor: "Generic",
	Features: []models.FixtureFeature{
		models.FeatureDimmer,
		models.FeatureRGB,
		models.FeaturePan,
		models.FeatureTilt,
	},

	Modes: []models.FixtureMode{
		{
			ID:           "generic.movinghead.9ch",
			Name:         "7 Channel",
			ChannelCount: 7,
			Channels: map[uint8]models.FixtureChannel{
				0: {Type: models.ChannelTypeDim},
				1: {Type: models.ChannelTypeRed},
				2: {Type: models.ChannelTypeGreen},
				3: {Type: models.ChannelTypeBlue},
				6: {Type: models.ChannelTypePan},
				7: {Type: models.ChannelTypeTilt},
			},
		},
	},
}
