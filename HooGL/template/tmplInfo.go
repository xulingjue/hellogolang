package template

/*
 *界面辅助类
 */
type TmplInfo struct {
	BaseUrl    string
	CurrentNav string

	ExtraJs  []string
	Js       []string
	ExtraCss []string
	Css      []string

	Title       string
	Description string
	Data        map[string]interface{}
	BackUrl     string
}

func (tmpl *TmplInfo) AddData(k string, v interface{}) {
	tmpl.new()
	tmpl.Data[k] = v
}

func (tmpl *TmplInfo) new() {
	if tmpl.Data == nil {
		tmpl.Data = make(map[string]interface{})
	}
}
