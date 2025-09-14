package apaas

type ApaasModeType uint8

const (
	EngineMode ApaasModeType = iota // gorm is default engine mode
	ViewMode
	DirectMode
)

func (m ApaasModeType) String() string {
	switch m {
	case EngineMode:
		return "EngineMode"
	case ViewMode:
		return "ViewMode"
	case DirectMode:
		return "DirectMode"
	}
	return "DirectMode"
}
