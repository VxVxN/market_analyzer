package humanizer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

func TestHumanizerHumanize(t *testing.T) {
	humanizer := Init(&marketanalyzer.MarketData{
		Quarters: []marketanalyzer.YearQuarter{
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
		PercentageChanges: map[marketanalyzer.RowName][]*big.Float{
			marketanalyzer.Earnings: {
				big.NewFloat(0),
				big.NewFloat(-1000000000),
				big.NewFloat(-1),
				big.NewFloat(-0.5),
				big.NewFloat(-0.333333333),
				big.NewFloat(-0.25),
				big.NewFloat(0.2),
			},
			marketanalyzer.Sales: {
				big.NewFloat(1000000000),
				big.NewFloat(1),
				big.NewFloat(0.5),
				big.NewFloat(0.333333333),
				big.NewFloat(0.25),
				big.NewFloat(0.2),
				big.NewFloat(0.166666667),
			},
			marketanalyzer.MarketCap: {
				big.NewFloat(0),
				big.NewFloat(1000000000),
				big.NewFloat(0),
				big.NewFloat(0),
				big.NewFloat(0),
				big.NewFloat(0),
				big.NewFloat(0),
			},
		},
		RawData: map[marketanalyzer.RowName][]*big.Int{
			marketanalyzer.Earnings: {
				big.NewInt(0),
				big.NewInt(-1000000000),
				big.NewInt(-2000000000),
				big.NewInt(-3000000000),
				big.NewInt(-4000000000),
				big.NewInt(-5000000000),
				big.NewInt(-4000000000),
			},
			marketanalyzer.Sales: {
				big.NewInt(1000000000),
				big.NewInt(2000000000),
				big.NewInt(3000000000),
				big.NewInt(4000000000),
				big.NewInt(5000000000),
				big.NewInt(6000000000),
				big.NewInt(7000000000),
			},
			marketanalyzer.MarketCap: {
				big.NewInt(0),
				big.NewInt(1000000000),
				big.NewInt(1000000000),
				big.NewInt(1000000000),
				big.NewInt(1000000000),
				big.NewInt(1000000000),
				big.NewInt(1000000000),
			},
		},
	})
	humanizer.SetPrecision(1)
	humanizer.SetFieldsForDisplay([]marketanalyzer.RowName{marketanalyzer.Earnings, marketanalyzer.Sales, marketanalyzer.MarketCap})
	humanizer.SetNumbersMode(NumbersWithPercentagesMode)
	data := humanizer.Humanize()

	testCases := []struct {
		name            string
		expectedHeaders []string
		expectedRows    [][]string
	}{
		// case 1
		{
			name: "check NumbersWithPercentagesMode",
			expectedHeaders: []string{
				"#",
				"2016/2",
				"2017/1",
				"2018/1",
				"2018/2",
				"2019/1",
				"2019/2",
				"2019/3",
			},
			expectedRows: [][]string{
				{
					"market_cap",
					"-",
					"1.000.000.000",
					"1.000.000.000(+0.0%)",
					"1.000.000.000(+0.0%)",
					"1.000.000.000(+0.0%)",
					"1.000.000.000(+0.0%)",
					"1.000.000.000(+0.0%)",
				},
				{
					"sales",
					"1.000.000.000",
					"2.000.000.000(+100.0%)",
					"3.000.000.000(+50.0%)",
					"4.000.000.000(+33.3%)",
					"5.000.000.000(+25.0%)",
					"6.000.000.000(+20.0%)",
					"7.000.000.000(+16.7%)",
				},
				{
					"earnings",
					"-",
					"-1.000.000.000",
					"-2.000.000.000(-100.0%)",
					"-3.000.000.000(-50.0%)",
					"-4.000.000.000(-33.3%)",
					"-5.000.000.000(-25.0%)",
					"-4.000.000.000(+20.0%)",
				},
			},
		},
	}
	for _, testCase := range testCases {
		for i, header := range data.Headers {
			expectedHeader := testCase.expectedHeaders[i]
			assert.Equalf(t, expectedHeader, header, "[%s] header not equal expected", testCase.name)
		}

		assert.Equalf(t, len(testCase.expectedRows), len(data.Rows), "[%s] expected numbers of row", testCase.name)

		for i, record := range data.Rows {
			expectedRecord := testCase.expectedRows[i]
			assert.Equalf(t, expectedRecord, record, "[%s] record not equal expected", testCase.name)
		}
	}
}
