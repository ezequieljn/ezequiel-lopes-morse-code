package morse

import (
	"bytes"
	"errors"
	"strings"
)

type MorseTwo struct {
	Text        string
	TextDecode  string
	index       int
	word        rune
	wordCurrent string
	space       string
	textBuffer  bytes.Buffer
}

func NewMorseTwo(space string) *MorseTwo {
	return &MorseTwo{
		space: space,
	}
}

func (m *MorseTwo) Decode(text string) (string, error) {
	m.Text = strings.TrimSpace(text)
	if err := m.validate(); err != nil {
		return "", err
	}
	for m.index = 0; m.index < len(m.Text); m.index++ {
		m.word = rune(m.Text[m.index])
		if m.addTheSymbol() {
			continue
		}
		ok, err := m.isSpace()
		if err != nil {
			return "", err
		}
		if ok {
			continue
		}
		if err := m.convertLatter(); err != nil {
			return "", err
		}

	}
	m.convertLastLatter()
	return m.textBuffer.String(), nil
}

func (m *MorseTwo) validate() error {
	if m.space != "   " && m.space != " / " {
		return errors.New("space needs to be '   ' or ' / '")
	}
	if m.Text == "" {
		return errors.New("text is empty")
	}
	allowedChars := ".- /"
	for _, v := range m.Text {
		if m.space != " / " && v == '/' {
			return errors.New("text invalid")
		}
		if !strings.ContainsRune(allowedChars, v) {
			return errors.New("text invalid")
		}
	}
	return nil
}

func (m *MorseTwo) convertLatter() error {
	v, ok := MorseMap[m.wordCurrent]
	if !ok {
		return errors.New("text invalid")
	}
	m.wordCurrent = ""
	m.textBuffer.WriteString(v)
	return nil
}

const indexLetters = 4

func (m *MorseTwo) isSpace() (bool, error) {
	if m.index+indexLetters > len(m.Text) {
		return false, nil
	}
	if ok := m.nextCharactersAreSpaces(); ok {
		v, err := ConvertMorseToLetter(m.wordCurrent)
		if err != nil {
			return false, err
		}
		m.textBuffer.WriteString(v)
		m.textBuffer.WriteString(" ")
		m.wordCurrent = ""
		m.index += 2
		return true, nil
	}
	return false, nil
}

func (m *MorseTwo) nextCharactersAreSpaces() bool {
	if string(m.word) == string(m.space[0]) && string(m.Text[m.index+1]) == string(m.space[1]) && string(m.Text[m.index+2]) == string(m.space[2]) {
		return true
	}
	return false
}

func (m *MorseTwo) addTheSymbol() bool {
	if string(m.word) != " " {
		m.wordCurrent += string(m.word)
		return true
	}
	return false
}

func (m *MorseTwo) convertLastLatter() (string, error) {
	if m.wordCurrent != "" {
		v, err := ConvertMorseToLetter(m.wordCurrent)
		if err != nil {
			return "", err
		}
		m.textBuffer.WriteString(v)
	}
	return "", errors.New("text invalid")
}

func ConvertMorseToLetter(letter string) (string, error) {
	value, ok := MorseMap[letter]
	if !ok {
		return "", errors.New("text invalid")
	}
	return value, nil
}
