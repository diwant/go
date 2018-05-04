package coinchanger

import "testing"
import "github.com/stretchr/testify/assert"

var test = []struct {
  description string
  input int
  expected map[string]int // ["name_of_coin"] => num_coins
} {
  {
    description: "Get 2 quarters if input is 50",
    input: 50,
    expected: map[string]int{"quarter":2},
  },
}

func TestEachCase(t *testing.T) {
  for _, tc := range tests {
    actual := MakeChange(tc.input)
    assert.Equal(t, tc.expected, actual) 
  }
}
