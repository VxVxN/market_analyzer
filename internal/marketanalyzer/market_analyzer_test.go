package marketanalyzer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		Earnings: {
			big.NewInt(-1000000000),
			big.NewInt(-2000000000),
			big.NewInt(-1000000000),
			big.NewInt(1000000000),
			big.NewInt(2000000000),
			big.NewInt(1000000000),
			big.NewInt(-1000000000),
		},
		Sales: {
			big.NewInt(1000000000),
			big.NewInt(2000000000),
			big.NewInt(3000000000),
			big.NewInt(4000000000),
			big.NewInt(5000000000),
			big.NewInt(6000000000),
			big.NewInt(7000000000),
		},
		MarketCap: {
			big.NewInt(1000000000),
			big.NewInt(1000000000),
			big.NewInt(1000000000),
			big.NewInt(1000000000),
			big.NewInt(1000000000),
			big.NewInt(1000000000),
			big.NewInt(1000000000),
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
				big.NewInt(7000000000),
				big.NewInt(18000000000),
			},
		},
	}
	for _, testCase := range testCases {
		marketAnalyzer.SetPeriodMode(testCase.periodMode)
		data, err := marketAnalyzer.Calculate()
		require.NoError(t, err)

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
		rowName                RowName
		expectedQuarters       []YearQuarter
		expectedPercentageData []*big.Float
	}{
		// case 1
		{
			name:    "check percentage changes in NormalMode",
			mode:    NormalMode,
			rowName: Sales,
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
				big.NewFloat(1),
				big.NewFloat(0.5),
				big.NewFloat(0.3333333333),
				big.NewFloat(0.25),
				big.NewFloat(0.2),
				big.NewFloat(0.1666666667),
			},
		},
		// case 2
		{
			name:    "check percentage changes in FirstQuarterMode",
			mode:    FirstQuarterMode,
			rowName: Sales,
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
				big.NewFloat(0.5),
				big.NewFloat(0.6666666667),
			},
		},
		// case 3
		{
			name:    "check percentage changes in YearMode",
			mode:    YearMode,
			rowName: Sales,
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
				big.NewFloat(1),
				big.NewFloat(2.5),
				big.NewFloat(1.571428571),
			},
		},
		// case 4
		{
			name:    "check percentage changes with negative and positive numbers",
			mode:    NormalMode,
			rowName: Earnings,
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
				big.NewFloat(-1000000000),
				big.NewFloat(-1),
				big.NewFloat(0.5),
				big.NewFloat(2),
				big.NewFloat(1),
				big.NewFloat(-0.5),
				big.NewFloat(-2),
			},
		},
	}
	for _, testCase := range testCases {
		marketAnalyzer.SetPeriodMode(testCase.mode)
		data, err := marketAnalyzer.Calculate()
		require.NoError(t, err)

		for _, quarter := range data.Quarters {
			var isFind bool
			for _, expectedQuarter := range testCase.expectedQuarters {
				if quarter.Year == expectedQuarter.Year && quarter.Quarter == expectedQuarter.Quarter {
					isFind = true
				}
			}
			assert.Equalf(t, true, isFind, "[%s] year with quarter not found %d %d", testCase.name, quarter.Year, quarter.Quarter)
		}

		percentageData, _ := data.PercentageChanges[testCase.rowName]
		assert.NotNilf(t, percentageData, "[%s] sales not found in percentage data", testCase.name)

		for i, record := range percentageData {
			expectedData := testCase.expectedPercentageData[i]
			assert.Equalf(t, expectedData.String(), record.String(), "[%s] percentage data not equal expected", testCase.name)
		}
	}
}
