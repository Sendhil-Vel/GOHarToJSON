package main

type MainObj struct {
	Main FullObj `json:main`
}

type FullObj struct {
	Info OnfoObj   `json:"info"`
	Item []ParaCol `json:"item"`
}

type OnfoObj struct {
	Name   string `json:"name"`
	Schema string `json:"schema"`
}
type ParaCol struct {
	Name    string     `json:"name"`
	Request RequestCol `json:"request"`
	Event   EventCol   `json:"event"`
}

// Event  EventCol    `json:"event"`
type RequestCol struct {
	Method string      `json:"method"`
	Url    URLCol      `json:"url"`
	Header []HeaderCol `json:"header"`
	Body   BodyCol     `json:"body"`
}

type URLCol struct {
	Raw      string   `json:"raw"`
	ProtoCol string   `json:"protocol"`
	Host     []string `json:"host"`
	Port     string   `json:"port"`
	Path     []string `json:"path"`
}

type HeaderCol struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BodyCol struct {
	Mode    string     `json:"mode"`
	Raw     string     `json:"raw"`
	Options OptionsCol `json:"options"`
}

type OptionsCol struct {
	Raw RawCol `json:"raw"`
}

type RawCol struct {
	Language string `json:"language"`
}

type EventCol struct {
	Listen string    `json:"listen"`
	Script ScriptCol `json:"script"`
}

type ScriptCol struct {
	Exec []string `json:"exec"`
	Type string   `json:"type"`
}
