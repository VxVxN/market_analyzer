package smartlabparser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParserParse(t *testing.T) {
	parser := Init("fixp_test_data.csv")
	data, err := parser.Parse()
	require.NoError(t, err, "failed to parse file")

	expectedQuarter := []string{
		"",
		"2017/4",
		"2018/4",
		"2019/4",
		"2020/1",
		"2020/2",
		"2020/3",
		"2020/4",
		"2021/1",
		"2021/2",
		"2021/3",
	}
	assert.NotEqual(t, 0, len(data.Headers), "headers not should be empty")

	for i, quarter := range data.Headers {
		if i == 0 { // first record should be empty
			continue
		}
		expectedValue := expectedQuarter[i]
		assert.Equal(t, expectedValue, quarter, "year/quarter not equal expected")
	}

	expectedRow := []string{
		"sales",
		"",
		"",
		"",
		"40000000000",
		"83000000000",
		"49100000000",
		"",
		"",
		"106000000000",
		"57900000000",
	}
	assert.NotEqual(t, 0, len(data.Rows), "data not should be empty")

	var checkRow []string
	for _, row := range data.Rows {
		if row[0] != expectedRow[0] {
			continue
		}
		checkRow = row
	}
	require.Equal(t, len(expectedRow), len(checkRow), "row len not equal expected")

	for i, expectedValues := range expectedRow {
		assert.Equal(t, expectedValues, checkRow[i], "record not equal expected")
	}
}
