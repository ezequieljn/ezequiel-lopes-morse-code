package morse

import (
	"errors"
	"strings"
)

type MorseOne struct {
	text       string
	textDecode string
	space      string
}

func NewMorseOne(space string) *MorseOne {
	return &MorseOne{
		space: space,
	}
}

func (m *MorseOne) Decode(text string) (string, error) {
	m.text = strings.TrimSpace(text)
	if err := m.validate(); err != nil {
		return "", err
	}
	word := strings.Split(m.text, m.space)
	for i, w := range word {
		letters := strings.Split(w, " ")
		for _, l := range letters {
			if err := m.convertLatter(l); err != nil {
				return "", err
			}
		}
		m.isSpace(len(word), i)
	}
	return m.textDecode, nil
}

const indexLatter = 1

func (m *MorseOne) isSpace(lenWord int, index int) {
	if index+indexLatter < lenWord {
		m.textDecode += " "
	}
}

func (m *MorseOne) convertLatter(latter string) error {
	v, ok := MorseMap[latter]
	if !ok {
		return errors.New("text invalid")
	}
	m.textDecode += v
	return nil
}

func (m *MorseOne) validate() error {
	if m.space != "   " && m.space != " / " {
		return errors.New("space needs to be '   ' or ' / '")
	}
	if m.text == "" {
		return errors.New("text is empty")
	}
	allowedChars := ".- /"
	for _, v := range m.text {
		if m.space != " / " && v == '/' {
			return errors.New("text invalid")
		}
		if !strings.ContainsRune(allowedChars, v) {
			return errors.New("text invalid")
		}
	}
	return nil
}
