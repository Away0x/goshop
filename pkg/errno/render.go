package errno

const (
	// RenderDetailTypeJSON json 渲染类型
	RenderDetailTypeJSON = "json"
	// RenderDetailTypeHTML html 渲染类型
	RenderDetailTypeHTML = "html"

	// DefaultRenderDetailType 默认渲染类型
	DefaultRenderDetailType = RenderDetailTypeJSON

	// RenderDetailHTMLTemplate error HTML 的模板地址
	RenderDetailHTMLTemplate = "errors/error.html"
)

// RenderDetail -
type RenderDetail struct {
	Type     string                 `json:"-"`                 // render 类型 json/html
	Content  map[string]interface{} `json:"content,omitempty"` // 具体错误信息
	Template string                 `json:"-"`                 // 模板 name
}

// NewRenderDetail -
func NewRenderDetail(tp string, content map[string]interface{}) *RenderDetail {
	tpl := ""

	switch tp {
	case RenderDetailTypeHTML:
		tpl = RenderDetailHTMLTemplate
	default:
	}

	if content == nil {
		content = map[string]interface{}{}
	}

	return &RenderDetail{
		Type:     tp,
		Content:  content,
		Template: tpl,
	}
}

// DefaultRenderDetail -
func DefaultRenderDetail(content map[string]interface{}) *RenderDetail {
	return NewRenderDetail(DefaultRenderDetailType, content)
}

// WrapRenderContent -
func WrapRenderContent(detail *RenderDetail, tp string, otherContent ...map[string]interface{}) *RenderDetail {
	var newDetail *RenderDetail

	if detail == nil {
		newDetail = NewRenderDetail(tp, map[string]interface{}{})
	} else {
		newDetail = NewRenderDetail(tp, detail.Content)
	}

	for _, v1 := range otherContent {
		for k2, v2 := range v1 {
			newDetail.Content[k2] = v2
		}
	}

	return newDetail
}
