package marketanalyzer

import "fmt"

type PeriodMode int

const (
	FirstQuarterMode PeriodMode = iota
	SecondQuarterMode
	ThirdQuarterMode
	FourthQuarterMode
	YearMode
	NormalMode
)

func (mode PeriodMode) String() string {
	switch mode {
	case FirstQuarterMode:
		return "first"
	case SecondQuarterMode:
		return "second"
	case ThirdQuarterMode:
		return "third"
	case FourthQuarterMode:
		return "four"
	case YearMode:
		return "year"
	case NormalMode:
		return "normal"
	}
	return ""
}

func PeriodModeFromString(str string) (PeriodMode, error) {
	switch str {
	case "first":
		return FirstQuarterMode, nil
	case "second":
		return SecondQuarterMode, nil
	case "third":
		return ThirdQuarterMode, nil
	case "four":
		return FourthQuarterMode, nil
	case "year":
		return YearMode, nil
	case "normal":
		return NormalMode, nil
	}
	return NormalMode, fmt.Errorf("invalid period mode: %v", str)
}
