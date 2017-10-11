package xui

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"

	xmldom "bitbucket.org/rj/xmldom-go"
	"gitee.com/ying32/govcl/vcl"
)

const (
	methodTypeClicked = iota + 0
	methodTypeChanged
	methodTypeSelected
	methodTypeToggled
)

// TEvents 之后使用都要继承此类型
type TEvents struct {
	// Form 默认，不可改变，会由内部构建时自动填充值
	Form *vcl.TForm
}

// TXMLForm
type TXMLForm struct {
	Form  *vcl.TForm
	event interface{}
}

func NewFromFile(xmlfile string, event interface{}) (*TXMLForm, error) {
	bs, err := ioutil.ReadFile(xmlfile)
	if err != nil {
		return nil, err
	}
	return NewFormBytes(bs, event)
}

func NewFormBytes(xmlstr []byte, event interface{}) (*TXMLForm, error) {
	doc, err := xmldom.ParseXml(bytes.NewReader(xmlstr))
	if err != nil {
		return nil, err
	}
	w := new(TXMLForm)
	w.event = event

	root := doc.DocumentElement()
	fmt.Println(root.NodeName())
	if root != nil && root.NodeName() == "Form" {
		w.Form = w.buildWindow(root)
		if w.Form.IsValid() {
			w.buildControls(root, w.Form, nil)
		} else {
			panic("Window创建失败!")
		}
	} else {
		panic("xml不符合要求！")
	}
	return w, nil
}

func (x *TXMLForm) setBounds(control vcl.IControl, attrs *TXmlAttrs) {
	if attrs.HasAttr("left") {
		control.SetLeft(attrs.Left())
	}
	if attrs.HasAttr("top") {
		control.SetTop(attrs.Top())
	}
	if attrs.HasAttr("width") {
		control.SetWidth(attrs.Width())
	}
	if attrs.HasAttr("height") {
		control.SetHeight(attrs.Height())
	}
}

func (x *TXMLForm) buildWindow(node xmldom.Node) *vcl.TForm {
	attrs := newXmlAttrsMap(node)
	if attrs == nil {
		return nil
	}
	w := vcl.Application.CreateForm()
	if w.IsValid() {
		x.setFiledVal("Form", w)
		x.setBounds(w, attrs)
		w.SetCaption(attrs.Caption())
		w.EnabledMaximize(attrs.EnabledMax())
		w.EnabledMinimize(attrs.EnabledMin())
		w.SetName(attrs.Name())
		if attrs.Center() {
			w.ScreenCenter()
		}
	}
	return w
}

func (x *TXMLForm) getMethod(name string) (reflect.Method, bool) {
	if name == "" {
		var m reflect.Method
		return m, false
	}
	return reflect.TypeOf(x.event).MethodByName(name)
}

// callMethod 动态call方法
func (x *TXMLForm) callMethod(m reflect.Method, param ...interface{}) {
	ps := make([]reflect.Value, len(param)+1)
	ps[0] = reflect.ValueOf(x.event)
	for i := 1; i <= len(param); i++ {
		ps[i] = reflect.ValueOf(param[i-1])
	}
	m.Func.Call(ps)
}

// setFiledVal 设置字段的值
func (x *TXMLForm) setFiledVal(name string, v interface{}) {
	vx := reflect.ValueOf(x.event).Elem().FieldByName(name)
	if vx.IsValid() {
		vx.Set(reflect.ValueOf(v))
	}
}

func (x *TXMLForm) buildControls(node xmldom.Node, parent vcl.IControl, menu *vcl.TMenuItem) {
	if !node.HasChildNodes() {
		return
	}
	var i uint
	var pcontrol vcl.IControl
	var attrs *TXmlAttrs

	for i = 0; i < node.ChildNodes().Length(); i++ {
		subnode := node.ChildNodes().Item(i)
		if subnode.NodeType() != 1 {
			continue
		}
		pcontrol = nil
		attrs = newXmlAttrsMap(subnode)

		//		fmt.Println("subnode.NodeName():", subnode.NodeName(), ", Caption:", attrs.Caption())
		switch subnode.NodeName() {

		case "MainMenu":

			mmenu := vcl.NewMainMenu(x.Form)
			mmenu.SetName(attrs.Name())
			pcontrol = nil
			x.buildControls(subnode, nil, mmenu.Items())

		case "MenuItem":

			pcontrol = nil
			if menu != nil {

				subm := vcl.NewMenuItem(x.Form)
				subm.SetCaption(attrs.Caption())
				subm.SetEnabled(attrs.Enabled())
				subm.SetChecked(attrs.Checked())
				subm.SetVisible(attrs.Visible())
				subm.SetName(attrs.Name())

				m, ok := x.getMethod(attrs.OnClick())
				if ok {
					subm.SetOnClick(func(sender vcl.IObject) {
						x.callMethod(m, sender)
					})
				}
				menu.Add(subm)
				x.buildControls(subnode, nil, subm)
			}

		case "Button":

			btn := vcl.NewButton(x.Form)
			btn.SetParent(parent)
			btn.SetCaption(attrs.Caption())
			x.setFiledVal(attrs.Name(), btn)
			pcontrol = btn
			m, ok := x.getMethod(attrs.OnClick())
			if ok {
				btn.SetOnClick(func(sender vcl.IObject) {
					x.callMethod(m, sender)
				})
			}

		case "Edit":
			edit := vcl.NewEdit(x.Form)
			edit.SetReadOnly(attrs.ReadOnly())
			edit.SetParent(parent)

			m, ok := x.getMethod(attrs.OnChange())
			if ok {
				edit.SetOnChange(func(sender vcl.IObject) {
					x.callMethod(m, sender)
				})
			}

			edit.SetText(attrs.Text())

		case "Memo":

			memo := vcl.NewMemo(x.Form)
			memo.SetReadOnly(attrs.ReadOnly())
			memo.SetParent(parent)

			m, ok := x.getMethod(attrs.OnChange())
			if ok {
				memo.SetOnChange(func(sender vcl.IObject) {
					x.callMethod(m, sender)
				})
			}
			pcontrol = memo
			if attrs.HasAttr("text") {
				memo.SetText(attrs.Text())
			}
			x.buildControls(subnode, memo, menu)

		case "Label":
			lbl := vcl.NewLabel(x.Form)
			lbl.SetParent(parent)
			lbl.SetCaption(attrs.Text())
			pcontrol = lbl

		case "Panel":
			pnl := vcl.NewPanel(x.Form)
			pcontrol = pnl
			pnl.SetParent(parent)
			x.buildControls(subnode, pnl, menu)

		case "Checkbox":
			chk := vcl.NewCheckBox(x.Form)
			chk.SetParent(parent)
			chk.SetCaption(attrs.Text())
			chk.SetChecked(attrs.Checked())

			pcontrol = chk

		case "Combobox":
			combox := vcl.NewComboBox(x.Form)
			combox.SetParent(parent)
			pcontrol = combox
			x.buildControls(subnode, combox, menu)
			combox.SetItemIndex(int32(attrs.ItemIndex()))

		case "PageControl":
			pgc := vcl.NewPageControl(x.Form)
			pcontrol = pgc
			pgc.SetParent(parent)
			x.setFiledVal(attrs.Name(), pgc)
			x.buildControls(subnode, pgc, menu)
			pgc.SetActivePageIndex(attrs.ActiveIndex())

		case "TabSheet":

			sheet := vcl.NewTabSheet(x.Form)
			pcontrol = sheet
			sheet.SetCaption(attrs.Caption())
			sheet.SetPageControl(parent)
			x.setFiledVal(attrs.Name(), sheet)
			x.buildControls(subnode, sheet, menu)

		// 伪类名
		case "TextItem":
			if parent.IsValid() {
				switch parent.ClassName() {
				case "TComboBox":
					vcl.ComboBoxFromObj(parent).Items().Add(attrs.Text())
				case "TMemo":
					vcl.MemoFromObj(parent).Lines().Add(attrs.Text())
				default:
				}
			}

		default:
			continue
		}

		if pcontrol != nil {
			x.setBounds(pcontrol, attrs)
			if attrs.HasAttr("name") {
				pcontrol.SetName(attrs.Name())
			}
			if attrs.HasAttr("enabled") {
				pcontrol.SetEnabled(attrs.Enabled())
			}
			if attrs.HasAttr("align") {
				pcontrol.SetAlign(attrs.Align())
			}
			if attrs.HasAttr("visible") {
				pcontrol.SetVisible(attrs.Visible())
			}
		}
	}
	return
}
