package xui

import (
	xmldom "bitbucket.org/rj/xmldom-go"
	"gitee.com/ying32/govcl/vcl/types"
)

type TXmlAttrs struct {
	attrs map[string]string
}

func newXmlAttrsMap(node xmldom.Node) (res *TXmlAttrs) {
	if node.NodeType() == 3 {
		return nil
	}
	res = new(TXmlAttrs)
	res.attrs = make(map[string]string, 0)
	var i uint
	for i = 0; i < node.Attributes().Length(); i++ {
		attr := node.Attributes().Item(i)
		if _, ok := res.attrs[attr.NodeName()]; !ok {
			res.attrs[attr.NodeName()] = attr.NodeValue()
		}
	}
	return
}

func (x *TXmlAttrs) HasAttr(name string) bool {
	if _, ok := x.attrs[name]; ok {
		return true
	}
	return false
}

func (x *TXmlAttrs) Get(name string) string {
	return x.GetDef(name, "")
}

func (x *TXmlAttrs) GetDef(name, def string) string {
	if v, ok := x.attrs[name]; ok {
		return v
	}
	return def
}

func (x *TXmlAttrs) ToInt(name string) int32 {
	return int32(atoi(x.Get(name)))
}

func (x *TXmlAttrs) ToIntDef(name string, def int) int32 {
	v := x.GetDef(name, "")
	if v == "" {
		return int32(def)
	}
	return int32(atoi(v))
}

func (x *TXmlAttrs) ToBool(name string) bool {
	return atob(x.Get(name))
}

func (x *TXmlAttrs) ToBoolDef(name string, def bool) bool {
	v := x.GetDef(name, "")
	if v == "" {
		return def
	}
	return atob(v)
}

func (x *TXmlAttrs) Title() string {
	return x.Get("title")
}

func (x *TXmlAttrs) Caption() string {
	return x.Get("caption")
}

func (x *TXmlAttrs) EnabledMax() bool {
	return x.ToBoolDef("enabledmax", true)
}

func (x *TXmlAttrs) EnabledMin() bool {
	return x.ToBoolDef("enabledmin", true)
}

func (x *TXmlAttrs) Text() string {
	return x.Get("text")
}

func (x *TXmlAttrs) Name() string {
	return x.Get("name")
}

func (x *TXmlAttrs) OnClick() string {
	return x.Get("onclick")
}

func (x *TXmlAttrs) OnExecute() string {
	return x.Get("onexecute")
}

func (x *TXmlAttrs) Action() string {
	return x.Get("action")
}

func (x *TXmlAttrs) OnUpdate() string {
	return x.Get("onupdate")
}

func (x *TXmlAttrs) Top() int32 {
	return x.ToInt("top")
}

func (x *TXmlAttrs) Left() int32 {
	return x.ToInt("left")
}

func (x *TXmlAttrs) Width() int32 {
	return x.ToInt("width")
}

func (x *TXmlAttrs) Height() int32 {
	return x.ToInt("height")
}

func (x *TXmlAttrs) Center() bool {
	return x.ToBool("center")
}

func (x *TXmlAttrs) Enabled() bool {
	return x.ToBoolDef("enabled", true)
}

func (x *TXmlAttrs) Visible() bool {
	return x.ToBoolDef("visible", true)
}

func (x *TXmlAttrs) Checked() bool {
	return x.ToBool("checked")
}

func (x *TXmlAttrs) OnChange() string {
	return x.Get("onchange")
}

func (x *TXmlAttrs) ReadOnly() bool {
	return x.ToBoolDef("readonly", false)
}

func (x *TXmlAttrs) ItemIndex() int32 {
	return x.ToIntDef("itemindex", -1)
}

func (x *TXmlAttrs) OnSelected() string {
	return x.Get("onselected")
}

func (x *TXmlAttrs) IntValue() int32 {
	return x.ToInt("value")
}

func (x *TXmlAttrs) Min() int32 {
	return x.ToInt("min")
}

func (x *TXmlAttrs) Max() int32 {
	return x.ToIntDef("max", 100)
}

func (x *TXmlAttrs) Hint() string {
	return x.Get("hint")
}

func (x *TXmlAttrs) ShowHint() bool {
	return x.ToBoolDef("showhint", false)
}

func (x *TXmlAttrs) ActiveIndex() int32 {
	return int32(x.ToIntDef("activeindex", 0))
}

func (x *TXmlAttrs) Align() types.TAlign {
	s := x.Get("align")
	switch s {
	case "alTop":
		return types.AlTop
	case "alBottom":
		return types.AlBottom
	case "alLeft":
		return types.AlLeft
	case "alRight":
		return types.AlRight
	case "alClient":
		return types.AlClient
	case "alCustom":
		return types.AlCustom
	default:
	}
	return types.AlNone
}
