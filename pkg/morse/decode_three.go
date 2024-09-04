package morse

import (
	"bytes"
	"errors"
	"strings"
	"sync"
)

type MorseTree struct {
	text       string
	finalWords []string
	space      string
}

func NewMorseThree(space string) *MorseTree {
	return &MorseTree{
		space: space,
	}
}

func (m *MorseTree) Decode(text string) (string, error) {
	m.text = strings.TrimSpace(text)
	if err := m.validate(); err != nil {
		return "", err
	}
	word := strings.Split(m.text, m.space)
	var errors []error
	wg := sync.WaitGroup{}
	wg.Add(len(word))
	m.finalWords = make([]string, len(word))
	for i, w := range word {
		go func() {
			defer wg.Done()
			letters := strings.Split(w, " ")
			for _, l := range letters {
				if err := m.convertLatter(i, l); err != nil {
					errors = append(errors, err)
				}
			}
			m.isSpace(len(word), i)
		}()
	}
	wg.Wait()
	var buffer bytes.Buffer
	for _, value := range m.finalWords {
		buffer.WriteString(value)
	}
	return buffer.String(), nil
}

func (m *MorseTree) isSpace(lenWord int, index int) {
	if index+indexLatter < lenWord {
		m.finalWords[index] += " "
	}
}

func (m *MorseTree) convertLatter(i int, latter string) error {
	v, ok := MorseMap[latter]
	if !ok {
		return errors.New("text invalid")
	}
	m.finalWords[i] += v
	return nil
}

func (m *MorseTree) validate() error {
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
