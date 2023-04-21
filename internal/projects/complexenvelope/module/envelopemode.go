package module

type EnvelopeMode int

const (
	EnvelopeModeAD = EnvelopeMode(iota)
)

func (e EnvelopeMode) String() string {
	switch e {
	case EnvelopeModeAD:
		return "AD"
	default:
		return ""
	}
}
