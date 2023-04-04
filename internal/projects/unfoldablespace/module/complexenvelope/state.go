package complexenvelope

type state int

const (
	stateIdle = state(iota)
	stateAttack
	stateDecay
	stateSustain
	stateRelease
)
