package humanizer

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

func TestHumanizer_Humanize(t *testing.T) {
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
				big.NewFloat(2000000000),
				big.NewFloat(3),
				big.NewFloat(3.5),
				big.NewFloat(3.333333333),
				big.NewFloat(3.25),
				big.NewFloat(3.2),
				big.NewFloat(3.166666667),
			},
			marketanalyzer.Sales: {
				big.NewFloat(1000000000),
				big.NewFloat(2),
				big.NewFloat(1.5),
				big.NewFloat(1.333333333),
				big.NewFloat(1.25),
				big.NewFloat(1.2),
				big.NewFloat(1.166666667),
			},
			marketanalyzer.Debts: {
				big.NewFloat(3000000000),
				big.NewFloat(4),
				big.NewFloat(4.5),
				big.NewFloat(4.333333333),
				big.NewFloat(4.25),
				big.NewFloat(4.2),
				big.NewFloat(4.166666667),
			},
		},
		RawData: map[marketanalyzer.RowName][]*big.Int{
			marketanalyzer.Sales: {
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
	humanizer.SetPrecision(1)
	humanizer.SetFieldsForDisplay([]marketanalyzer.RowName{marketanalyzer.Sales})
	humanizer.SetNumbersMode(NumbersWithPercentagesMode)
	data := humanizer.Humanize()

	expectedHeaders := []string{
		"#",
		"2016/2",
		"2017/1",
		"2018/1",
		"2018/2",
		"2019/1",
		"2019/2",
		"2019/3",
	}

	for i, header := range data.Headers {
		expectedHeader := expectedHeaders[i]
		assert.Equal(t, expectedHeader, header, "header not equal expected")
	}

	expectedRows := [][]string{
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
	}
	assert.Equal(t, 1, len(data.Rows), "expected one row")

	for i, record := range data.Rows {
		expectedRecord := expectedRows[i]
		assert.Equal(t, expectedRecord, record, "record not equal expected")
	}
}
