package xiangling

import (
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/glog"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/reactable"
	"github.com/genshinsim/gcsim/pkg/target"
)

type panda struct {
	*target.Target
	*reactable.Reactable
}

func newGuoba(c *core.Core) *panda {
	p := &panda{}
	p.Target = target.New(c, 0, 0, 0.5)
	p.Reactable = &reactable.Reactable{}
	p.Reactable.Init(p, c)

	p.Target.HPCurrent = 1
	p.Target.HPMax = 1

	return p
}

func (p *panda) Type() combat.TargettableType { return combat.TargettableObject }

func (p *panda) Attack(atk *combat.AttackEvent, evt glog.Event) (float64, bool) {
	//don't take damage, trigger swirl reaction only on sucrose E
	if p.Core.Player.Chars()[atk.Info.ActorIndex].Base.Key != keys.Sucrose {
		return 0, false
	}
	if atk.Info.AttackTag != combat.AttackTagElementalArt {
		return 0, false
	}
	//check pyro window
	if p.Durability[attributes.Pyro] <= 0 {
		return 0, false
	}

	//cheat a bit, set the durability just enough to match incoming sucrose E gauge
	oldDur := p.Durability[attributes.Pyro]
	p.Durability[attributes.Pyro] = infuseDurability
	p.React(atk)
	// restore the durability after
	p.Durability[attributes.Pyro] = oldDur

	return 0, false
}
