package go_msteams

type ActionCard struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	Actions []Action `json:"actions"`
	Inputs []Input `json:"inputs"`
}

type OpenUriTarget struct {
	Uri string `json:"uri"`
	Os string `json:"os"`
}

type OpenUri struct {
	Type string             `json:"@type"`
	Name string             `json:"name"`
	Targets []OpenUriTarget `json:"targets"`
}

type Action interface {
	isAction()
}

type Input interface {
	isInput()
}

type Header struct {
	Name string `json:"name"`
	Value string `json:"value"`
}

type HttpPostAction struct {
	Type string `json:"@type"`
	Name string `json:"name"`
	Target string `json:"target"`
	Body string `json:"body"`
	Headers []Header `json:"headers"`
}



type TextInput struct {
	Type string `json:"@type"`
	Id string `json:"id"`
	Title string `json:"title"`
	IsMultiline bool `json:"isMultiline"`
	MaxLength int `json:"maxLength"`
}
type DateInput struct {
	Type string `json:"@type"`
	Id string `json:"id"`
	Title string `json:"title"`
	IncludeTime bool `json:"includeTime"`
}

type MultilineChoice struct {
	Display string `json:"display"`
	Value string `json:"value"`
}

type MultilineInput struct {
	Type string `json:"@type"`
	Id string `json:"id"`
	Title string `json:"title"`
	Choices []MultilineChoice `json:"choices"`
	IsMultiSelect bool `json:"isMultiSelect"`
	Style string `json:"style"`
}

type Button interface {
	isButton()
}

func (h HttpPostAction) isAction(){}
func (h HttpPostAction) isButton(){}
func (h *HttpPostAction) AddHeader(name string, value string){
	h.Headers = append(h.Headers, Header{Value: value, Name: name})
}
func (tp TextInput) isInput(){}
func (tp DateInput) isInput(){}
func (tp MultilineInput) isInput(){}
func (tp *MultilineInput) addChoice(display string, value string){
	tp.Choices = append(tp.Choices, MultilineChoice{
		Display: display,
		Value: value,
	})
}

func (a ActionCard) isButton() {}
func (o OpenUri) isButton() {}
func (o OpenUri) isAction() {}


func NewHTTPPostAction(name string,target string, body string) HttpPostAction {
	return HttpPostAction{
		Type:   "HttpPOST",
		Name:   name,
		Target: target,
		Body:   body,
	}
}


func NewTextInput(id string, title string, isMultiline bool, maxLength int) TextInput{
	return TextInput{
		Type:        "TextInput",
		Id:          id,
		Title:       title,
		IsMultiline: isMultiline,
		MaxLength:   maxLength,
	}
}

func NewDateInput(id string, title string, includeTime bool ) DateInput {
	return DateInput{
		Type:        "DateInput",
		Id:          id,
		Title:       title,
		IncludeTime: includeTime,
	}
}

func NewMultilineInput(id string, title string, isMultiSelect bool) MultilineInput {
	return MultilineInput{
		Type:          "MultichoiceInput",
		Id:            id,
		Title:         title,
		Choices:       []MultilineChoice{},
		IsMultiSelect: isMultiSelect,
		Style:         "",
	}
}

func NewActionCard(name string) ActionCard {
	return ActionCard{
		Type:    "ActionCard",
		Name:    name,
		Actions: []Action{},
		Inputs:  []Input{},
	}
}
func (b *ActionCard) AddAction(action Action) {
	b.Actions = append(b.Actions, action)
}
func (b *ActionCard) AddInput(input Input) {
	b.Inputs = append(b.Inputs, input)
}

func (o *OpenUri) AddTarget(target OpenUriTarget) {
	o.Targets = append(o.Targets, target)
}

func NewOpenUri(name string) OpenUri {
	return OpenUri{
		Type:    "OpenUri",
		Name:    name,
		Targets: []OpenUriTarget{},
	}
}

func NewOpenUriTarget(uri string, os string) OpenUriTarget {
	return OpenUriTarget{
		Uri: uri,
		Os:  os,
	}
}