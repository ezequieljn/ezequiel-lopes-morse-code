package morse_test

import (
	"testing"

	"github.com/ezequieljn/morse-code/pkg/morse"
)

func BenchmarkCodeMorse(b *testing.B) {
	space := "   "
	textForBenchmark := []string{
		".... . -.--   .--- ..- -.. .",
		"       --. .-.. --- -... ---    ",
		". --.. . --.- ..- .. . .-..",
		"    .--. .- .-. .- -... . -. ...",
		".-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..",
		".-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..   .-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..   .-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..   .-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..   .-   --. . -. - .   ... .   ...- .   .--. --- .-.   .- --.- ..- ..",
	}

	b.Run("code morse version: one", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, text := range textForBenchmark {
				codeMorse := morse.NewMorseTwo(space)
				codeMorse.Decode(text)
			}
		}
	})

	b.Run("code morse version: two", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, text := range textForBenchmark {
				codeMorse := morse.NewMorseTwo(space)
				codeMorse.Decode(text)
			}
		}
	})

	b.Run("code morse version: three", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for _, text := range textForBenchmark {
				codeMorse := morse.NewMorseThree(space)
				codeMorse.Decode(text)
			}
		}
	})
}
