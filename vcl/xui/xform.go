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
		control.SetLeft(attrs.Top())
	}
	if attrs.HasAttr("width") {
		control.SetLeft(attrs.Width())
	}
	if attrs.HasAttr("height") {
		control.SetLeft(attrs.Height())
	}
}

func (x *TXMLForm) buildWindow(node xmldom.Node) *vcl.TForm {
	attrs := newXmlAttrsMap(node)
	if attrs == nil {
		return nil
	}
	w := vcl.Application.CreateForm()
	if w.IsValid() {
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
				//				subm.set
				m, ok := x.getMethod(attrs.OnClick())
				if ok {
					subm.SetOnClick(func(sender vcl.IObject) {
						m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
					})
				}
				menu.Add(subm)
				x.buildControls(subnode, nil, subm)
			}

		case "Button":
			btn := vcl.NewButton(x.Form)
			btn.SetParent(parent)
			btn.SetCaption(attrs.Caption())
			pcontrol = btn
			m, ok := x.getMethod(attrs.OnClick())
			if ok {
				btn.SetOnClick(func(sender vcl.IObject) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}

		case "Edit":
			edit := vcl.NewEdit(x.Form)
			edit.SetReadOnly(attrs.ReadOnly())
			edit.SetParent(parent)

			m, ok := x.getMethod(attrs.OnChange())
			if ok {
				edit.SetOnChange(func(sender vcl.IObject) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}

			edit.SetText(attrs.Text())

		case "Memo":
			memo := vcl.NewMemo(x.Form)
			memo.SetReadOnly(attrs.ReadOnly())

			m, ok := x.getMethod(attrs.OnChange())
			if ok {
				memo.SetOnChange(func(sender vcl.IObject) {
					m.Func.Call([]reflect.Value{reflect.ValueOf(x.event), reflect.ValueOf(sender)})
				})
			}
			pcontrol = memo
			memo.SetText(attrs.Text())

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
			combox.SetItemIndex(int32(attrs.Selected()))
			continue

		default:
			continue
		}

		if pcontrol != nil {
			x.setBounds(pcontrol, attrs)
			pcontrol.SetEnabled(attrs.Enabled())
			pcontrol.SetVisible(attrs.Visible())
			pcontrol.SetAlign(attrs.Align())
		}
	}
	return
}
