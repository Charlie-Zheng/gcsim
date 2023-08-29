package combat

import (
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

func (c *Handler) AbsorbCheck(p AttackPattern, prio ...attributes.Element) attributes.Element {
	// check targets for collision first
	for _, e := range prio {
		for _, x := range c.enemies {
			t, ok := x.(TargetWithAura)
			if !ok {
				continue
			}
			if collision, _ := t.AttackWillLand(p); collision && t.AuraContains(e) {
				c.Log.NewEvent(
					"infusion check picked up "+e.String(),
					glog.LogElementEvent,
					-1,
				).
					Write("source", "enemy").
					Write("key", t.Key())
				return e
			}
		}
		for _, x := range c.gadgets {
			t, ok := x.(TargetWithAura)
			if !ok {
				continue
			}
			if collision, _ := t.AttackWillLand(p); collision && t.AuraContains(e) {
				c.Log.NewEvent(
					"infusion check picked up "+e.String(),
					glog.LogElementEvent,
					-1,
				).
					Write("source", "gadget").
					Write("key", t.Key())
				return e
			}
		}
		if t, ok := c.player.(TargetWithAura); ok {
			if collision, _ := t.AttackWillLand(p); collision && t.AuraContains(e) {
				c.Log.NewEvent(
					"infusion check picked up "+e.String(),
					glog.LogElementEvent,
					-1,
				).
					Write("source", "player")
				return e
			}
		}

	}
	return attributes.NoElement
}