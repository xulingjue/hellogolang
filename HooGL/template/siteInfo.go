package template

/*
 *界面辅助类
 */
type SiteInfo struct {
	BaseUrl    string
	CurrentNav string

	ExtraJs  []string
	Js       []string
	ExtraCss []string
	Css      []string

	Title       string
	SiteName    string
	Description string

	BackUrl string
}
