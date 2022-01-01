package humanizer

import "fmt"

type NumberMode int

const (
	NumbersWithPercentagesMode NumberMode = iota
	NumbersMode
	PercentagesMode
)

func (mode NumberMode) String() string {
	switch mode {
	case NumbersWithPercentagesMode:
		return "num_with_percent"
	case NumbersMode:
		return "number"
	case PercentagesMode:
		return "percent"
	}
	return ""
}

func NumberModeFromString(str string) (NumberMode, error) {
	switch str {
	case "num_with_percent":
		return NumbersWithPercentagesMode, nil
	case "number":
		return NumbersMode, nil
	case "percent":
		return PercentagesMode, nil
	}
	return NumbersMode, fmt.Errorf("invalid number mode: %v", str)
}
