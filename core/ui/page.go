package ui

type IPage interface {
	GetTitle() string
	GetWidgets() []IWidget
	GetComponentName() string
}

type Page struct {
	Title     string    `json:"title"`
	Component string    `json:"component"`
	Widgets   []IWidget `json:"widgets"`
}

var _ IPage = (*Page)(nil)

func (p *Page) GetTitle() string {
	return p.Title
}

func (p *Page) GetWidgets() []IWidget {
	return p.Widgets
}

func (p *Page) GetComponentName() string {
	return p.Component
}
