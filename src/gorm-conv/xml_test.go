package main

import (
	"encoding/xml"
	//"io/ioutil"
	"os"
	"testing"
)

func Test_Xml(t *testing.T) {
	t.Logf("testXml")

	f, err := os.Open("./test.xml")
	//bytes, err := ioutil.ReadFile("./test.xml")
	if err != nil {
		t.Logf(err.Error())
	}

	xmlDecoder := xml.NewDecoder(f)
	if xmlDecoder == nil {
		t.Logf("get decoder failed")
	}
	//xml.Unmarshal(bytes)

	var tk xml.Token
	var text bool
	for tk, err = xmlDecoder.Token(); err == nil; tk, err = xmlDecoder.Token() {
		switch token := tk.(type) {
		case xml.StartElement:
			name := token.Name.Local
			if name == "text" {
				text = true
			}
			t.Log(token.Name)
			t.Log(token.Attr)
		case xml.EndElement:
			text = false
		case xml.CharData:
			if text {
				content := string([]byte(token))
				t.Logf(content)
				return
			}
			t.Log(token)
		default:
			return
			//
		}
	}
}
