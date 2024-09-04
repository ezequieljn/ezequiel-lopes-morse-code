package morse_test

import (
	"testing"

	"github.com/ezequieljn/ezequiel-lopes-morse-code/pkg/morse"
	"github.com/stretchr/testify/assert"
)

func TestDecodeMorseTwo(t *testing.T) {
	space := "   "
	type ExpectedTest struct {
		textValue     string
		expectedValue string
	}

	t.Run("should return error with empty text version two", func(t *testing.T) {
		assert := assert.New(t)
		value := ""
		morseCode := morse.NewMorseTwo(space)
		textDecode, err := morseCode.Decode(value)
		assert.EqualError(err, "text is empty")
		assert.Equal("", textDecode)
	})

	t.Run("should return error with invalid text version two", func(t *testing.T) {
		assert := assert.New(t)
		value := "*********"
		morseCode := morse.NewMorseTwo(space)
		textDecode, err := morseCode.Decode(value)
		assert.EqualError(err, "text invalid")
		assert.Equal("", textDecode)
	})

	t.Run("should be able to decode morse version two", func(t *testing.T) {
		expect := []ExpectedTest{
			{".... . -.--   .--- ..- -.. .", "HEY JUDE"},
			{"       --. .-.. --- -... ---    ", "GLOBO"},
			{". --.. . --.- ..- .. . .-..", "EZEQUIEL"},
			{"    .--. .- .-. .- -... . -. ...", "PARABENS"},
			{".-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..", "A GENTE SE VE POR AQUI"},
			{"...", "S"},
		}
		for _, e := range expect {
			m := morse.NewMorseTwo(space)
			textDecode, err := m.Decode(e.textValue)
			assert.Nil(t, err)
			assert.Equal(t, e.expectedValue, textDecode)
		}
	})
	t.Run("should error space", func(t *testing.T) {
		assert := assert.New(t)
		spaces := []string{"******", "", "  "}
		for _, space := range spaces {
			m := morse.NewMorseTwo(space)
			textDecode, err := m.Decode("... --- ...")
			assert.EqualError(err, "space needs to be '   ' or ' / '")
			assert.Equal("", textDecode)
		}
	})
	t.Run("should not return error in space", func(t *testing.T) {
		assert := assert.New(t)
		spaces := []string{"   ", " / "}
		for _, space := range spaces {
			m := morse.NewMorseTwo(space)
			textDecode, err := m.Decode("... --- ...")
			assert.Nil(err)
			assert.Equal("SOS", textDecode)
		}
	})

}
