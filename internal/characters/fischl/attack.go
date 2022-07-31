package fischl

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
)

var attackFrames [][]int
var attackHitmarks = []int{15, 11, 24, 26, 21}

const normalHitNum = 5

func init() {
	attackFrames = make([][]int, normalHitNum)

	// normal cancels (missing Nx -> Aim)
	attackFrames[0] = frames.InitNormalCancelSlice(attackHitmarks[0], 25)
	attackFrames[1] = frames.InitNormalCancelSlice(attackHitmarks[1], 22)
	attackFrames[2] = frames.InitNormalCancelSlice(attackHitmarks[2], 38)
	attackFrames[3] = frames.InitNormalCancelSlice(attackHitmarks[3], 32)
	attackFrames[4] = frames.InitNormalCancelSlice(attackHitmarks[4], 67)
}

func (c *char) Attack(p map[string]int) action.ActionInfo {
	travel, ok := p["travel"]
	if !ok {
		travel = 10
	}

	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       fmt.Sprintf("Normal %v", c.NormalCounter),
		AttackTag:  combat.AttackTagNormal,
		ICDTag:     combat.ICDTagNone,
		ICDGroup:   combat.ICDGroupDefault,
		StrikeType: combat.StrikeTypePierce,
		Element:    attributes.Physical,
		Durability: 25,
		Mult:       auto[c.NormalCounter][c.TalentLvlAttack()],
	}
	c.Core.QueueAttack(
		ai,
		combat.NewDefSingleTarget(c.Core.Combat.DefaultTarget, combat.TargettableEnemy),
		attackHitmarks[c.NormalCounter],
		attackHitmarks[c.NormalCounter]+travel,
	)

	//check for c1
	if c.Base.Cons >= 1 && c.ozActiveUntil < c.Core.F {
		ai := combat.AttackInfo{
			ActorIndex: c.Index,
			Abil:       "Fischl C1",
			AttackTag:  combat.AttackTagNormal,
			ICDTag:     combat.ICDTagNone,
			ICDGroup:   combat.ICDGroupDefault,
			StrikeType: combat.StrikeTypePierce,
			Element:    attributes.Physical,
			Durability: 100,
			Mult:       0.22,
		}
		c.Core.QueueAttack(
			ai,
			combat.NewDefSingleTarget(c.Core.Combat.DefaultTarget, combat.TargettableEnemy),
			attackHitmarks[c.NormalCounter],
			attackHitmarks[c.NormalCounter]+travel,
		)
	}

	defer c.AdvanceNormalIndex()

	return action.ActionInfo{
		Frames:          frames.NewAttackFunc(c.Character, attackFrames),
		AnimationLength: attackFrames[c.NormalCounter][action.InvalidAction],
		CanQueueAfter:   attackHitmarks[c.NormalCounter],
		State:           action.NormalAttackState,
	}
}