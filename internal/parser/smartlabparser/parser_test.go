package smartlabparser

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/VxVxN/market_analyzer/internal/marketanalyzer"
)

func TestParserParse(t *testing.T) {
	parser := Init("fixp_test_data.csv")
	data, err := parser.Parse()
	require.NoError(t, err, "failed to parse file")

	expectedQuarter := []marketanalyzer.YearQuarter{
		{
			Year:    2017,
			Quarter: 4,
		},
		{
			Year:    2018,
			Quarter: 4,
		},
		{
			Year:    2019,
			Quarter: 4,
		},
		{
			Year:    2020,
			Quarter: 1,
		},
		{
			Year:    2020,
			Quarter: 2,
		},
		{
			Year:    2020,
			Quarter: 3,
		},
		{
			Year:    2020,
			Quarter: 4,
		},
		{
			Year:    2021,
			Quarter: 1,
		},
		{
			Year:    2021,
			Quarter: 2,
		},
		{
			Year:    2021,
			Quarter: 3,
		},
	}
	assert.NotEqual(t, 0, len(data.YearQuarters), "headers not should be empty")

	for i, quarter := range data.YearQuarters {
		if i == 0 { // first record should be empty
			continue
		}
		expectedValue := expectedQuarter[i]
		assert.Equal(t, expectedValue, quarter, "year/quarter not equal expected")
	}

	expectedData := map[marketanalyzer.RowName][]*big.Int{
		marketanalyzer.Sales: {
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(40000000000),
			big.NewInt(83000000000),
			big.NewInt(49100000000),
			big.NewInt(0),
			big.NewInt(0),
			big.NewInt(106000000000),
			big.NewInt(57900000000),
		},
	}
	assert.NotEqual(t, 0, len(data.Data), "data not should be empty")

	for name, expectedValues := range expectedData {
		records, ok := data.Data[name]
		assert.Equalf(t, true, ok, "row %s not found", name)

		for i, record := range records {
			expectedValue := expectedValues[i]
			assert.Equal(t, expectedValue, record, "record not equal expected")
		}
	}
}
