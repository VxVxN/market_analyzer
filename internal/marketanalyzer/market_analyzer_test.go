package marketanalyzer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

var marketAnalyzer = Init(&RawMarketData{
	YearQuarters: []YearQuarter{
		{
			Year:    2016,
			Quarter: 2,
		},
		{
			Year:    2017,
			Quarter: 1,
		},
		{
			Year:    2018,
			Quarter: 1,
		},
		{
			Year:    2018,
			Quarter: 2,
		},
		{
			Year:    2019,
			Quarter: 1,
		},
		{
			Year:    2019,
			Quarter: 2,
		},
		{
			Year:    2019,
			Quarter: 3,
		},
	},
	Data: map[RowName][]*big.Int{
		Sales: {
			big.NewInt(1000000000),
			big.NewInt(2000000000),
			big.NewInt(3000000000),
			big.NewInt(4000000000),
			big.NewInt(5000000000),
			big.NewInt(6000000000),
			big.NewInt(7000000000),
		},
	},
})

func TestMarketAnalyzerCalculatePeriodMode(t *testing.T) {
	testCases := []struct {
		name             string
		periodMode       PeriodMode
		expectedQuarters []YearQuarter
		expectedRawData  []*big.Int
	}{
		// case 1
		{
			name:       "check SecondQuarterMode",
			periodMode: SecondQuarterMode,
			expectedQuarters: []YearQuarter{
				{
					Year:    2016,
					Quarter: 2,
				},
				{
					Year:    2018,
					Quarter: 2,
				},
				{
					Year:    2019,
					Quarter: 2,
				},
			},
			expectedRawData: []*big.Int{
				big.NewInt(1000000000),
				big.NewInt(4000000000),
				big.NewInt(6000000000),
			},
		},
		// case 2
		{
			name:       "check YearMode",
			periodMode: YearMode,
			expectedQuarters: []YearQuarter{
				{
					Year: 2016,
				},
				{
					Year: 2017,
				},
				{
					Year: 2018,
				},
				{
					Year: 2019,
				},
			},
			expectedRawData: []*big.Int{
				big.NewInt(1000000000),
				big.NewInt(2000000000),
				big.NewInt(3500000000),
				big.NewInt(6000000000),
			},
		},
	}
	for _, testCase := range testCases {
		marketAnalyzer.SetPeriodMode(testCase.periodMode)
		data := marketAnalyzer.Calculate()

		for _, quarter := range data.Quarters {
			var isFind bool
			for _, expectedQuarter := range testCase.expectedQuarters {
				if quarter.Year == expectedQuarter.Year && quarter.Quarter == expectedQuarter.Quarter {
					isFind = true
				}
			}
			assert.Equalf(t, true, isFind, "[%s] year with quarter not found %d %d", testCase.name, quarter.Year, quarter.Quarter)
		}

		rawData, _ := data.RawData[Sales]
		assert.NotNilf(t, rawData, "[%s] sales not found in raw data", testCase.name)

		for i, record := range rawData {
			expectedData := testCase.expectedRawData[i]
			assert.Equalf(t, expectedData, record, "[%s] raw data not equal expected", testCase.name)
		}
	}
}

func TestMarketAnalyzerCalculatePercentageChanges(t *testing.T) {
	testCases := []struct {
		name                   string
		mode                   PeriodMode
		expectedQuarters       []YearQuarter
		expectedPercentageData []*big.Float
	}{
		// case 1
		{
			name: "check percentage changes in NormalMode",
			mode: NormalMode,
			expectedQuarters: []YearQuarter{
				{
					Year:    2016,
					Quarter: 2,
				},
				{
					Year:    2017,
					Quarter: 1,
				},
				{
					Year:    2018,
					Quarter: 1,
				},
				{
					Year:    2018,
					Quarter: 2,
				},
				{
					Year:    2019,
					Quarter: 1,
				},
				{
					Year:    2019,
					Quarter: 2,
				},
				{
					Year:    2019,
					Quarter: 3,
				},
			},
			expectedPercentageData: []*big.Float{
				big.NewFloat(1000000000),
				big.NewFloat(2),
				big.NewFloat(1.5),
				big.NewFloat(1.333333333),
				big.NewFloat(1.25),
				big.NewFloat(1.2),
				big.NewFloat(1.166666667),
			},
		},
		// case 2
		{
			name: "check percentage changes in FirstQuarterMode",
			mode: FirstQuarterMode,
			expectedQuarters: []YearQuarter{
				{
					Year:    2017,
					Quarter: 1,
				},
				{
					Year:    2018,
					Quarter: 1,
				},
				{
					Year:    2019,
					Quarter: 1,
				},
			},
			expectedPercentageData: []*big.Float{
				big.NewFloat(2000000000),
				big.NewFloat(1.5),
				big.NewFloat(1.666666667),
			},
		},
		// case 3
		{
			name: "check percentage changes in YearMode",
			mode: YearMode,
			expectedQuarters: []YearQuarter{
				{
					Year: 2016,
				},
				{
					Year: 2017,
				},
				{
					Year: 2018,
				},
				{
					Year: 2019,
				},
			},
			expectedPercentageData: []*big.Float{
				big.NewFloat(1000000000),
				big.NewFloat(2),
				big.NewFloat(1.75),
				big.NewFloat(1.714285714),
			},
		},
	}
	for _, testCase := range testCases {
		marketAnalyzer.SetPeriodMode(testCase.mode)
		data := marketAnalyzer.Calculate()

		for _, quarter := range data.Quarters {
			var isFind bool
			for _, expectedQuarter := range testCase.expectedQuarters {
				if quarter.Year == expectedQuarter.Year && quarter.Quarter == expectedQuarter.Quarter {
					isFind = true
				}
			}
			assert.Equalf(t, true, isFind, "[%s] year with quarter not found %d %d", testCase.name, quarter.Year, quarter.Quarter)
		}

		percentageData, _ := data.PercentageChanges[Sales]
		assert.NotNilf(t, percentageData, "[%s] sales not found in percentage data", testCase.name)

		for i, record := range percentageData {
			expectedData := testCase.expectedPercentageData[i]
			assert.Equalf(t, expectedData.String(), record.String(), "[%s] percentage data not equal expected", testCase.name)
		}
	}
}
