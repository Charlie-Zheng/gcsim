// Code generated by "pipeline"; DO NOT EDIT.
package itto

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
	1: {"travel", "ushihit"},
	2: {"prestack"},
	5: {"collision"},
	6: {"collision"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Itto, ValidateParamKeys)
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
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
)

var (
	// attack: akCombo = [5]
	akCombo = []float64{
		0.9115999937057495,
		0.98580002784729,
		1.059999942779541,
		1.1660000085830688,
		1.2402000427246094,
		1.3250000476837158,
		1.44159996509552,
		1.5582000017166138,
		1.6748000383377075,
		1.8020000457763672,
		1.9477499723434448,
		2.119152069091797,
		2.2905540466308594,
		2.461956024169922,
		2.648940086364746,
	}
	// attack: akFinal = [6]
	akFinal = []float64{
		1.9091999530792236,
		2.0645999908447266,
		2.2200000286102295,
		2.441999912261963,
		2.597399950027466,
		2.7750000953674316,
		3.019200086593628,
		3.263400077819824,
		3.5076000690460205,
		3.7739999294281006,
		4.079249858856201,
		4.438223838806152,
		4.7971978187561035,
		5.156171798706055,
		5.5477800369262695,
	}
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.7923179864883423,
		0.8568090200424194,
		0.9212999939918518,
		1.0134299993515015,
		1.0779210329055786,
		1.1516250371932983,
		1.2529679536819458,
		1.3543109893798828,
		1.4556540250778198,
		1.5662100315093994,
		1.6928889751434326,
		1.8418630361557007,
		1.9908369779586792,
		2.1398110389709473,
		2.3023290634155273,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.7636799812316895,
		0.8258399963378906,
		0.8880000114440918,
		0.9768000245094299,
		1.0389599800109863,
		1.1100000143051147,
		1.2076799869537354,
		1.305359959602356,
		1.4030400514602661,
		1.509600043296814,
		1.631700038909912,
		1.7752900123596191,
		1.9188790321350098,
		2.062469005584717,
		2.219111919403076,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.9164159893989563,
		0.9910079836845398,
		1.065600037574768,
		1.172160029411316,
		1.2467520236968994,
		1.3320000171661377,
		1.4492160081863403,
		1.566431999206543,
		1.6836479902267456,
		1.811519980430603,
		1.9580399990081787,
		2.130347967147827,
		2.302654981613159,
		2.4749629497528076,
		2.6629340648651123,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		1.1722489595413208,
		1.2676639556884766,
		1.3630800247192383,
		1.4993879795074463,
		1.594804048538208,
		1.7038500308990479,
		1.85378897190094,
		2.003727912902832,
		2.1536660194396973,
		2.3172359466552734,
		2.5046589374542236,
		2.725069999694824,
		2.9454801082611084,
		3.1658899784088135,
		3.406337022781372,
	}
	// attack: collision = [8]
	collision = []float64{
		0.8183349967002869,
		0.8849430084228516,
		0.9515519738197327,
		1.046707034111023,
		1.1133160591125488,
		1.1894400119781494,
		1.2941110134124756,
		1.3987809419631958,
		1.503451943397522,
		1.6176379919052124,
		1.7318249940872192,
		1.8460110425949097,
		1.9601969718933105,
		2.074383020401001,
		2.188570022583008,
	}
	// attack: highPlunge = [10]
	highPlunge = []float64{
		2.0438549518585205,
		2.2102160453796387,
		2.3765759468078613,
		2.61423397064209,
		2.7805941104888916,
		2.970720052719116,
		3.232142925262451,
		3.4935669898986816,
		3.7549901008605957,
		4.0401787757873535,
		4.3253679275512695,
		4.6105570793151855,
		4.895747184753418,
		5.180935859680176,
		5.466125011444092,
	}
	// attack: lowPlunge = [9]
	lowPlunge = []float64{
		1.6363229751586914,
		1.7695120573043823,
		1.9027010202407837,
		2.092971086502075,
		2.2261600494384766,
		2.378376007080078,
		2.5876729488372803,
		2.7969698905944824,
		3.0062670707702637,
		3.234591007232666,
		3.4629149436950684,
		3.691240072250366,
		3.9195640087127686,
		4.14788818359375,
		4.376212120056152,
	}
	// attack: saichiSlash = [4]
	saichiSlash = []float64{
		0.9047200083732605,
		0.9783599972724915,
		1.0520000457763672,
		1.1571999788284302,
		1.2308399677276611,
		1.315000057220459,
		1.4307199716567993,
		1.5464400053024292,
		1.662160038948059,
		1.7884000539779663,
		1.9330500364303589,
		2.1031579971313477,
		2.2732670307159424,
		2.4433751106262207,
		2.6289479732513428,
	}
	// skill: skill = [0]
	skill = []float64{
		3.072000026702881,
		3.3024001121520996,
		3.5327999591827393,
		3.8399999141693115,
		4.070400238037109,
		4.30079984664917,
		4.607999801635742,
		4.915200233459473,
		5.222400188446045,
		5.529600143432617,
		5.8368000984191895,
		6.144000053405762,
		6.5279998779296875,
		6.9120001792907715,
		7.296000003814697,
	}
	// burst: defconv = [1]
	defconv = []float64{
		0.5759999752044678,
		0.6191999912261963,
		0.6624000072479248,
		0.7200000286102295,
		0.7631999850273132,
		0.8064000010490417,
		0.8640000224113464,
		0.9215999841690063,
		0.979200005531311,
		1.0368000268936157,
		1.0944000482559204,
		1.1519999504089355,
		1.2239999771118164,
		1.2960000038146973,
		1.3680000305175781,
	}
)
