// Code generated by "pipeline"; DO NOT EDIT.
package chevreuse

import (
	_ "embed"

	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	1: {"hold"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Chevreuse, ValidateParamKeys)
}

func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		attack_3,
		{attack_4},
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.531299,
		0.574545,
		0.61779,
		0.679569,
		0.722814,
		0.772237,
		0.840194,
		0.908151,
		0.976108,
		1.050243,
		1.124378,
		1.198513,
		1.272647,
		1.346782,
		1.420917,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.493107,
		0.533243,
		0.57338,
		0.630718,
		0.670855,
		0.716725,
		0.779797,
		0.842869,
		0.90594,
		0.974746,
		1.043552,
		1.112357,
		1.181163,
		1.249968,
		1.318774,
	}
	// attack: attack_3 = [2 3]
	attack_3 = [][]float64{
		{
			0.276449,
			0.298951,
			0.321453,
			0.353598,
			0.3761,
			0.401816,
			0.437176,
			0.472535,
			0.507895,
			0.546469,
			0.585044,
			0.623618,
			0.662192,
			0.700767,
			0.739341,
		},
		{
			0.324527,
			0.350942,
			0.377357,
			0.415093,
			0.441508,
			0.471697,
			0.513206,
			0.554715,
			0.596225,
			0.641508,
			0.68679,
			0.732073,
			0.777356,
			0.822639,
			0.867922,
		},
	}
	// attack: attack_4 = [4]
	attack_4 = []float64{
		0.772615,
		0.835503,
		0.89839,
		0.988229,
		1.051116,
		1.122988,
		1.22181,
		1.320633,
		1.419456,
		1.527263,
		1.63507,
		1.742877,
		1.850683,
		1.95849,
		2.066297,
	}
	// attack: charge = [5]
	charge = []float64{
		1.2169,
		1.31595,
		1.415,
		1.5565,
		1.65555,
		1.76875,
		1.9244,
		2.08005,
		2.2357,
		2.4055,
		2.5753,
		2.7451,
		2.9149,
		3.0847,
		3.2545,
	}
	// skill: arkhe = [6]
	arkhe = []float64{
		0.288,
		0.3096,
		0.3312,
		0.36,
		0.3816,
		0.4032,
		0.432,
		0.4608,
		0.4896,
		0.5184,
		0.5472,
		0.576,
		0.612,
		0.648,
		0.684,
	}
	// skill: skillHold = [1]
	skillHold = []float64{
		1.728,
		1.8576,
		1.9872,
		2.16,
		2.2896,
		2.4192,
		2.592,
		2.7648,
		2.9376,
		3.1104,
		3.2832,
		3.456,
		3.672,
		3.888,
		4.104,
	}
	// skill: skillHpFlat = [5]
	skillHpFlat = []float64{
		256.79184,
		282.47458,
		310.29758,
		340.26077,
		372.36423,
		406.60788,
		442.9918,
		481.51593,
		522.1803,
		564.98486,
		609.9297,
		657.0147,
		706.24,
		757.6055,
		811.11127,
	}
	// skill: skillHpRegen = [4]
	skillHpRegen = []float64{
		0.026667,
		0.028667,
		0.030667,
		0.033333,
		0.035333,
		0.037333,
		0.04,
		0.042667,
		0.045333,
		0.048,
		0.050667,
		0.053333,
		0.056667,
		0.06,
		0.063333,
	}
	// skill: skillOvercharged = [2]
	skillOvercharged = []float64{
		2.824,
		3.0358,
		3.2476,
		3.53,
		3.7418,
		3.9536,
		4.236,
		4.5184,
		4.8008,
		5.0832,
		5.3656,
		5.648,
		6.001,
		6.354,
		6.707,
	}
	// skill: skillPress = [0]
	skillPress = []float64{
		1.152,
		1.2384,
		1.3248,
		1.44,
		1.5264,
		1.6128,
		1.728,
		1.8432,
		1.9584,
		2.0736,
		2.1888,
		2.304,
		2.448,
		2.592,
		2.736,
	}
	// burst: burst = [0]
	burst = []float64{
		3.6816,
		3.95772,
		4.23384,
		4.602,
		4.87812,
		5.15424,
		5.5224,
		5.89056,
		6.25872,
		6.62688,
		6.99504,
		7.3632,
		7.8234,
		8.2836,
		8.7438,
	}
	// burst: burstSecondary = [1]
	burstSecondary = []float64{
		0.49088,
		0.527696,
		0.564512,
		0.6136,
		0.650416,
		0.687232,
		0.73632,
		0.785408,
		0.834496,
		0.883584,
		0.932672,
		0.98176,
		1.04312,
		1.10448,
		1.16584,
	}
)
