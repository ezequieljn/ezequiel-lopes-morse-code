package morse_test

import (
	"testing"

	"github.com/ezequieljn/morse-code/pkg/morse"
	"github.com/stretchr/testify/assert"
)

func TestDecodeMorseOne(t *testing.T) {
	space := "   "
	type ExpectedTest struct {
		value    string
		expected string
	}

	t.Run("should return error with empty text version one", func(t *testing.T) {
		assert := assert.New(t)
		value := ""
		morseCode := morse.NewMorseOne(space)
		textDecode, err := morseCode.Decode(value)
		assert.EqualError(err, "text is empty")
		assert.Equal("", textDecode)
	})

	t.Run("should return error with invalid text version two", func(t *testing.T) {
		assert := assert.New(t)
		value := "*********"
		morseCode := morse.NewMorseOne(space)
		textDecode, err := morseCode.Decode(value)
		assert.EqualError(err, "text invalid")
		assert.Equal("", textDecode)
	})

	t.Run("should be able to decode morse version one", func(t *testing.T) {
		assert := assert.New(t)
		expect := []ExpectedTest{
			{".... . -.--   .--- ..- -.. .", "HEY JUDE"},
			{"       --. .-.. --- -... ---    ", "GLOBO"},
			{". --.. . --.- ..- .. . .-..", "EZEQUIEL"},
			{"    .--. .- .-. .- -... . -. ...", "PARABENS"},
			{".-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..", "A GENTE SE VE POR AQUI"},
			{"...", "S"},
		}
		for _, e := range expect {
			m := morse.NewMorseOne(space)
			textDecode, err := m.Decode(e.value)
			assert.Nil(err)
			assert.Equal(e.expected, textDecode)
		}
	})

	t.Run("should error space", func(t *testing.T) {
		assert := assert.New(t)
		spaces := []string{"******", "", "  "}
		for _, space := range spaces {
			m := morse.NewMorseOne(space)
			textDecode, err := m.Decode("... --- ...")
			assert.EqualError(err, "space needs to be '   ' or ' / '")
			assert.Equal("", textDecode)
		}
	})
	t.Run("should not return error in space", func(t *testing.T) {
		assert := assert.New(t)
		spaces := []string{"   ", " / "}
		for _, space := range spaces {
			m := morse.NewMorseOne(space)
			textDecode, err := m.Decode("... --- ...")
			assert.Nil(err)
			assert.Equal("SOS", textDecode)
		}
	})

}
