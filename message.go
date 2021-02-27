package go_msteams

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	MessageCard = "MessageCard"
	DefaultContext = "http://schema.org/extensions"
)

type Fact struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type Section struct {
	Title string `json:"activityTitle"`
	Subtitle string `json:"activitySubtitle"`
	Image string `json:"activityImage"`
	Facts []Fact `json:"facts"`
}

func (s *Section) AddFact(name string, value string) {
	var fact = Fact{
		Name: name,
		Value: value,
	}
	if s.Facts == nil {
		s.Facts = []Fact{fact}
		return
	}
	s.Facts = append(s.Facts, fact)

}

type Message struct {
	MessageType string `json:"@type"`
	Context string `json:"@context"`
	Summary string `json:"summary"`
	Text string `json:"text"`
	Sections []Section `json:"sections"`
	Color string `json:"themeColor"`

}

func (m *Message) SetColor(color string) {
	m.Color = color
}

func (m *Message) AddSection(title string, subtitle string, image string) *Section{
	var section = Section{
		Title: title,
		Subtitle: subtitle,
		Image: image,
	}
	if m.Sections == nil {
		m.Sections = []Section{}
	}
	m.Sections = append(m.Sections, section)

	var index = len(m.Sections)
	return &m.Sections[index - 1]

}

type Sender struct {
	webhookUrl string
}


func GetDemoMessage() Message {
	return Message{
		MessageType: MessageCard,
		Context: "http://schema.org/extensions",
		Summary: "Larry Bryant created a new task",
		Text: "Demo text",
		Sections: []Section{
			{
				Title: "Demo section",
				Subtitle: "Demo subtitle",
				Image: "https://images.genius.com/4226c55edbafb61f52f4d6be22a9c65c.500x500x1.jpg",
				Facts: []Fact{
					{
						Name: "Demo Fact",
						Value: "Demo Value",
					},
				},
			},
		},

	}
}

func NewMessage(summary string) Message{
	return Message{
		MessageType: MessageCard,
		Context:     DefaultContext,
		Summary:     summary,
		Text:        "",
		Sections:    nil,
	}
}

func NewSection(title string, subtitle string, image string) Section {
	return Section{
		Title:   title,
		Subtitle: subtitle,
		Image:    image,
		Facts:    []Fact{},
	}
}


func (s Sender) SendMessage(message Message) error{
	var data, err = json.Marshal(message)
	if err != nil {
		return err
	}
	resp, err := http.Post(s.webhookUrl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	body, _ := io.ReadAll(resp.Body)
	responseString := string(body)
	if responseString == "1" {
		return nil
	}
	return errors.New(responseString)

}