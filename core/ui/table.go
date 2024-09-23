package ui

type TableColumn struct {
	Prop  string         `json:"prop"`
	Label string         `json:"label"`
	Width string         `json:"width"`
	Props map[string]any `json:"props"`
}

type Table struct {
	Widget
	Columns []TableColumn `json:"columns"`
	Forms   []string      `json:"forms"`
}

var _ IWidget = (*Table)(nil)

func NewTable() *Table {
	return &Table{
		Widget: Widget{
			Name: "table",
		},
	}
}

func (t *Table) WithColumns(columns []TableColumn) *Table {
	t.Columns = columns
	return t
}
