package instrmtleggrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentleg"
)

func New() *InstrmtLegGrp {
	var m InstrmtLegGrp
	return &m
}

//NoLegs is a repeating group in InstrmtLegGrp
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//InstrmtLegGrp is a fix50sp1 Component
type InstrmtLegGrp struct {
	//NoLegs is a non-required field for InstrmtLegGrp.
	NoLegs []NoLegs `fix:"555,omitempty"`
}

func (m *InstrmtLegGrp) SetNoLegs(v []NoLegs) { m.NoLegs = v }
