package main

// DO NOT EDIT: This file was generated by vugu. Please regenerate instead of editing or add additional code in a separate file.

import "fmt"
import "reflect"
import "encoding/json"
import "github.com/vugu/vugu"
import js "github.com/vugu/vugu/js"

import "log"

type Root struct {
	Something	int
	Success		bool

	EmailValue	string
	UseSuffix	bool
}

func (c *Root) OnClickRun(event *vugu.DOMEvent, n int) {
	c.Success = !c.Success
	//log.Printf("HEY, GOT HERE!")
}

func (c *Root) OnUseSuffixChange(event *vugu.DOMEvent) {
	es := event.EventSummary()
	t, _ := es["target"].(map[string]interface{})
	log.Printf("OnUseSuffixChange: %#v", t["checked"])
	//c.EmailValue, _ = t["value"].(string)
}

func (c *Root) EmailChanged(event *vugu.DOMEvent) {

	es := event.EventSummary()
	t, _ := es["target"].(map[string]interface{})
	c.EmailValue, _ = t["value"].(string)

	b, _ := json.Marshal(event.EventSummary())
	log.Printf("event.EventSummary: %s", b)
}

func (c *Root) Build(vgin *vugu.BuildIn) (vgout *vugu.BuildOut, vgreterr error) {

	vgout = &vugu.BuildOut{}

	var vgn *vugu.VGNode
	vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "html", Attr: []vugu.VGAttribute(nil)}
	vgout.Out = append(vgout.Out, vgn)	// root for output
	{
		vgparent := vgn
		_ = vgparent
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "head", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "title", Attr: []vugu.VGAttribute(nil)}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "Test page"}
				vgparent.AppendChild(vgn)
			}
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "link", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "rel", Val: "stylesheet"}, vugu.VGAttribute{Namespace: "", Key: "href", Val: "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.CSS = append(vgout.CSS, vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "script", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "src", Val: "https://code.jquery.com/jquery-3.3.1.slim.min.js"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.JS = append(vgout.JS, vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "script", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "src", Val: "https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.JS = append(vgout.JS, vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
			vgparent.AppendChild(vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "script", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "src", Val: "https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js"}, vugu.VGAttribute{Namespace: "", Key: "integrity", Val: "sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM"}, vugu.VGAttribute{Namespace: "", Key: "crossorigin", Val: "anonymous"}}}
			vgout.JS = append(vgout.JS, vgn)
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n    "}
			vgparent.AppendChild(vgn)
		}
		vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "body", Attr: []vugu.VGAttribute(nil)}
		vgparent.AppendChild(vgn)
		{
			vgparent := vgn
			_ = vgparent
			vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "test-div"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "test_div_id"}}}
			vgparent.AppendChild(vgn)
			{
				vgparent := vgn
				_ = vgparent
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            This is a test.\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "run1"}}}
				vgparent.AppendChild(vgn)
				vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "data-whatever", Val: fmt.Sprint(c.Something)})
				vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
					EventType:	"click",
					Func:		func(event *vugu.DOMEvent) { c.OnClickRun(event, 7) },
					// TODO: implement capture, etc. mostly need to decide syntax
				})
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "run1"}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "button", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "run2"}}}
				vgparent.AppendChild(vgn)
				vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "data-whatever", Val: fmt.Sprint(c.Something)})
				vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
					EventType:	"click",
					Func:		func(event *vugu.DOMEvent) { c.Success = !c.Success },
					// TODO: implement capture, etc. mostly need to decide syntax
				})
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "run2"}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				if c.Success {
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "success"}}}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "success"}
						vgparent.AppendChild(vgn)
					}
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
				vgparent.AppendChild(vgn)
				{
					vghtml := "Some <strong>content</strong> here"
					vgn.InnerHTML = &vghtml
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
				vgparent.AppendChild(vgn)
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "container"}}}
				vgparent.AppendChild(vgn)
				{
					vgparent := vgn
					_ = vgparent
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "form", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-group"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "for", Val: "exampleInputEmail1"}}}
							vgparent.AppendChild(vgn)
							{
								vghtml := "Email address"
								vgn.InnerHTML = &vghtml
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "input", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "email"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-control"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "exampleInputEmail1"}, vugu.VGAttribute{Namespace: "", Key: "aria-describedby", Val: "emailHelp"}, vugu.VGAttribute{Namespace: "", Key: "placeholder", Val: "Enter email"}}}
							vgparent.AppendChild(vgn)
							vgn.Attr = append(vgn.Attr, vugu.VGAttribute{Key: "value", Val: fmt.Sprint(c.EmailValue)})
							vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
								EventType:	"change",
								Func:		func(event *vugu.DOMEvent) { c.EmailChanged(event) },
								// TODO: implement capture, etc. mostly need to decide syntax
							})
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "small", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "id", Val: "emailHelp"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-text text-muted"}}}
							vgparent.AppendChild(vgn)
							{
								vghtml := "We\x26#39;ll never share your email with anyone else."
								vgn.InnerHTML = &vghtml
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
						vgparent.AppendChild(vgn)
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-check"}}}
						vgparent.AppendChild(vgn)
						{
							vgparent := vgn
							_ = vgparent
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "input", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "type", Val: "checkbox"}, vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-check-input"}, vugu.VGAttribute{Namespace: "", Key: "id", Val: "exampleCheck1"}}}
							vgparent.AppendChild(vgn)
							{
								b, err := json.Marshal(true)
								if err != nil {
									return nil, err
								}
								vgn.Prop = append(vgn.Prop, vugu.VGProperty{Key: "checked", Val: json.RawMessage(b)})
							}
							vgn.DOMEventHandlerSpecList = append(vgn.DOMEventHandlerSpecList, vugu.DOMEventHandlerSpec{
								EventType:	"change",
								Func:		func(event *vugu.DOMEvent) { c.OnUseSuffixChange(event) },
								// TODO: implement capture, etc. mostly need to decide syntax
							})
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                        "}
							vgparent.AppendChild(vgn)
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "label", Attr: []vugu.VGAttribute{vugu.VGAttribute{Namespace: "", Key: "class", Val: "form-check-label"}, vugu.VGAttribute{Namespace: "", Key: "for", Val: "exampleCheck1"}}}
							vgparent.AppendChild(vgn)
							{
								vghtml := "Assume @gmail.com"
								vgn.InnerHTML = &vghtml
							}
							vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                    "}
							vgparent.AppendChild(vgn)
						}
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
						vgparent.AppendChild(vgn)
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n                "}
					vgparent.AppendChild(vgn)
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "div", Attr: []vugu.VGAttribute(nil)}
					vgparent.AppendChild(vgn)
					{
						vgparent := vgn
						_ = vgparent
						vgn = &vugu.VGNode{Type: vugu.VGNodeType(3), Data: "em", Attr: []vugu.VGAttribute(nil)}
						vgparent.AppendChild(vgn)
						{
							vghtml := c.EmailValue
							vgn.InnerHTML = &vghtml
						}
					}
					vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n            "}
					vgparent.AppendChild(vgn)
				}
				vgn = &vugu.VGNode{Type: vugu.VGNodeType(1), Data: "\n        "}
				vgparent.AppendChild(vgn)
			}
		}
	}
	return vgout, nil
}

// 'fix' unused imports
var _ fmt.Stringer
var _ reflect.Type
var _ json.RawMessage
var _ js.Value
