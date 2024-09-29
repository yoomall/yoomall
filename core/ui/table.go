package ui

type TableColumn struct {
	Prop  string         `json:"prop"`
	Label string         `json:"label"`
	Width string         `json:"width"`
	Props map[string]any `json:"props"`
}

type TableActionType string

var (
	TableActionTypeForm = TableActionType("form")
	TableActionTypeApi  = TableActionType("api")
	TableActionTypeLink = TableActionType("link")
)

type Confirm struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type Action struct {
	Label string          `json:"label"`
	Icon  string          `json:"icon"`
	Props map[string]any  `json:"props"`
	Type  TableActionType `json:"type"`

	FormKey string `json:"formKey"`

	ApiKey    string              `json:"apiKey"`
	ParamKeys []map[string]string `json:"paramKeys"`
	Confirm   *Confirm            `json:"confirm"`

	Path string `json:"path"`
}

func NewEditAction() *Action {
	return &Action{
		Label:   "编辑",
		Icon:    "ep:edit",
		Type:    TableActionTypeForm,
		FormKey: "user-form",
	}
}

func NewDeleteAction() *Action {
	return &Action{
		Label:  "删除",
		Icon:   "ep:delete",
		Type:   TableActionTypeApi,
		ApiKey: "user-delete",
		ParamKeys: []map[string]string{
			{
				"key":   "id",
				"label": "ID",
			},
		},
		Props: map[string]any{"color": "red", "type": "danger"},
	}
}

func (a *Action) WithApiKey(key string) *Action {
	a.ApiKey = key
	return a
}

func (a *Action) WithFormKey(key string) *Action {
	a.FormKey = key
	return a
}

func (a *Action) WithConfirm(title, text string) *Action {
	a.Confirm = &Confirm{
		Title: title,
		Text:  text,
	}
	return a
}

func (a *Action) WithPath(path string) *Action {
	a.Path = path
	return a
}

func (a *Action) WithParamKeys(keys ...map[string]string) *Action {
	a.ParamKeys = keys
	return a
}

func (a *Action) WithProps(props map[string]any) *Action {
	a.Props = props
	return a
}

func (a *Action) WithLabel(label string) *Action {
	a.Label = label
	return a
}

type Form struct {
	Key  string      `json:"key"`
	Name string      `json:"name"`
	Rows []*FormItem `json:"rows"`
}

type FormItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

func NewFormItem(label, value string) *FormItem {
	return &FormItem{
		Label: label,
		Value: value,
	}
}

func NewForm(key, name string) *Form {
	return &Form{
		Key:  key,
		Name: name,
	}
}

func (f *Form) WithRows(rows []*FormItem) *Form {
	f.Rows = rows
	return f
}

type Table struct {
	Widget
	Columns []TableColumn    `json:"columns"`
	Forms   map[string]*Form `json:"forms"`
	Actions []*Action        `json:"actions"`
}

var _ IWidget = (*Table)(nil)

func NewTable() *Table {
	return &Table{
		Widget: Widget{
			Name: "table",
		},
	}
}

func (t *Table) WithForms(forms map[string]*Form) *Table {
	t.Forms = forms
	return t
}

func (t *Table) WithColumns(columns []TableColumn) *Table {
	t.Columns = columns
	return t
}

func (t *Table) WithActions(actions []*Action) *Table {
	t.Actions = actions
	return t
}
