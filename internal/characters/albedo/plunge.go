package albedo

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player"
)

var highPlungeFrames []int
var lowPlungeFrames []int

const collisionHitmark = 30
const highPlungeHitmark = 34
const lowPlungeHitmark = 35

func init() {
	// high_plunge -> x
	highPlungeFrames = frames.InitAbilSlice(66)
	highPlungeFrames[action.ActionAttack] = 44
	highPlungeFrames[action.ActionSkill] = 43
	highPlungeFrames[action.ActionBurst] = 42
	highPlungeFrames[action.ActionDash] = 39
	highPlungeFrames[action.ActionJump] = 61
	highPlungeFrames[action.ActionSwap] = 47

	// low_plunge -> x
	lowPlungeFrames = frames.InitAbilSlice(64)
	lowPlungeFrames[action.ActionAttack] = 41
	lowPlungeFrames[action.ActionSkill] = 42
	lowPlungeFrames[action.ActionBurst] = 40
	lowPlungeFrames[action.ActionDash] = 36
	lowPlungeFrames[action.ActionJump] = 59
	lowPlungeFrames[action.ActionSwap] = 46
}

// High Plunge attack damage queue generator
// Use the "collision" optional argument if you want to do a falling hit on the way down
// Default = 0
func (c *char) HighPlungeAttack(p map[string]int) (action.Info, error) {
	switch c.Core.Player.Airborne() {
	case player.AirborneVenti:
		return action.Info{}, fmt.Errorf("%s high_plunge while airborne due to venti is unimplemented due to lack of frame data. Please see https://docs.gcsim.app/mechanics/frames for how to contribute", c.Base.Key.String())
	case player.AirborneXianyun:
		return c.highPlungeXY(p)
	default:
		return action.Info{}, fmt.Errorf("%s high_plunge can only be used while airborne", c.Base.Key.String())
	}
}

func (c *char) highPlungeXY(p map[string]int) (action.Info, error) {
	collision, ok := p["collision"]
	if !ok {
		collision = 0 // Whether or not collision hit
	}

	if collision > 0 {
		c.plungeCollision(collisionHitmark)
	}

	poiseDMG := 150.0
	highPlungeRadius := 5.0

	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "High Plunge",
		AttackTag:  attacks.AttackTagPlunge,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeBlunt,
		PoiseDMG:   poiseDMG,
		Element:    attributes.Physical,
		Durability: 25,
		Mult:       highPlunge[c.TalentLvlAttack()],
	}
	c.Core.QueueAttack(
		ai,
		combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, highPlungeRadius),
		highPlungeHitmark,
		highPlungeHitmark,
	)

	return action.Info{
		Frames:          frames.NewAbilFunc(highPlungeFrames),
		AnimationLength: highPlungeFrames[action.InvalidAction],
		CanQueueAfter:   highPlungeFrames[action.ActionAttack],
		State:           action.PlungeAttackState,
	}, nil
}

// Low Plunge attack damage queue generator
// Use the "collision" optional argument if you want to do a falling hit on the way down
// Default = 0
func (c *char) LowPlungeAttack(p map[string]int) (action.Info, error) {
	switch c.Core.Player.Airborne() {
	case player.AirborneVenti:
		return action.Info{}, fmt.Errorf("%s low_plunge while airborne due to venti hold E is unimplemented due to lack of frame data. Please see https://docs.gcsim.app/mechanics/frames for how to contribute", c.Base.Key.String())
	case player.AirborneXianyun:
		return c.lowPlungeXY(p)
	default:
		return action.Info{}, fmt.Errorf("%s low_plunge can only be used while airborne", c.Base.Key.String())
	}
}

func (c *char) lowPlungeXY(p map[string]int) (action.Info, error) {
	collision, ok := p["collision"]
	if !ok {
		collision = 0 // Whether or not collision hit
	}

	if collision > 0 {
		c.plungeCollision(collisionHitmark)
	}

	poiseDMG := 100.0
	lowPlungeRadius := 3.0

	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Low Plunge",
		AttackTag:  attacks.AttackTagPlunge,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeBlunt,
		PoiseDMG:   poiseDMG,
		Element:    attributes.Physical,
		Durability: 25,
		Mult:       lowPlunge[c.TalentLvlAttack()],
	}
	c.Core.QueueAttack(
		ai,
		combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, lowPlungeRadius),
		lowPlungeHitmark,
		lowPlungeHitmark,
	)

	return action.Info{
		Frames:          frames.NewAbilFunc(lowPlungeFrames),
		AnimationLength: lowPlungeFrames[action.InvalidAction],
		CanQueueAfter:   lowPlungeFrames[action.ActionSkill],
		State:           action.PlungeAttackState,
	}, nil
}

// Plunge normal falling attack damage queue generator
// Standard - Always part of high/low plunge attacks
func (c *char) plungeCollision(delay int) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Plunge Collision",
		AttackTag:  attacks.AttackTagPlunge,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeSlash,
		Element:    attributes.Physical,
		Durability: 0,
		Mult:       collision[c.TalentLvlAttack()],
	}
	c.Core.QueueAttack(ai, combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 1), delay, delay)
}