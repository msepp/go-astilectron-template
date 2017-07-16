package message

import "testing"

type Foo struct {
	Foo       int    `json:"foo"`
	Bar       string `json:"bar"`
	Xyz       bool   `json:"xyz"`
	TagMapped string `mapstructure:"a_field" json:"a_field"`
}

func TestDataMap(t *testing.T) {
	f := map[string]interface{}{"foo": 123.0, "bar": "xyzzy", "xyz": true}
	m := Message{ID: "id123", Key: "key", Type: "type", Data: f}

	dm, ok := m.DataMap()
	if !ok {
		t.Fatal("Unable to get datamap!")
	}

	if i, ok := dm.GetFloat64("foo"); !ok || i != 123.0 {
		t.Error("Didn't get foo as float64 correctly")
	}

	if i, ok := dm.GetInt("foo"); !ok || i != 123 {
		t.Error("Didn't get foo as int correctly")
	}

	if i, ok := dm.GetString("bar"); !ok || i != "xyzzy" {
		t.Error("Didn't get bar as string correctly")
	}

	if i, ok := dm.GetBool("xyz"); !ok || !i {
		t.Error("Didn't get xyz as bool correctly")
	}
}

func TestInto(t *testing.T) {
	f := map[string]interface{}{
		"foo":     123.0,
		"bar":     "xyzzy",
		"xyz":     true,
		"a_field": "a_value",
	}
	m := Message{ID: "id123", Key: "key", Type: "type", Data: f}

	var tgt Foo
	if err := m.Into(&tgt); err != nil {
		t.Errorf("Error converting into Foo: %s", err)
	}

	if tgt.Foo != 123 {
		t.Errorf("Wrong value for Foo: %d", tgt.Foo)
	}
	if tgt.Bar != "xyzzy" {
		t.Errorf("Wrong value for Bar: %s", tgt.Bar)
	}
	if tgt.Xyz != true {
		t.Errorf("Wrong value for Xyz: %t", tgt.Xyz)
	}
	if tgt.TagMapped != "a_value" {
		t.Errorf("Wrong value for TagMapped: %s", tgt.TagMapped)
	}
}

type afoo struct {
	A int `json:"a"`
	B int `json:"b"`
}
type bbar struct {
	afoo     `mapstructure:",squash"`
	Selected bool `json:"selected"`
}

func TestIntoComposite(t *testing.T) {
	input := map[string]interface{}{
		"a":        123,
		"b":        456,
		"selected": true,
	}

	m := Message{ID: "id123", Key: "key", Type: "type", Data: input}
	var tgt bbar
	if err := m.Into(&tgt); err != nil {
		t.Errorf("Error converting into Foo: %s", err)
	}

	if tgt.A != input["a"] {
		t.Errorf("tgt.A has invalid value: %d (not %d)", tgt.A, input["a"])
	}
	if tgt.B != input["b"] {
		t.Errorf("tgt.A has invalid value: %d (not %d)", tgt.B, input["b"])
	}
	if tgt.Selected != input["selected"] {
		t.Errorf("tgt.Selected has invalid value: %t (not %t)", tgt.Selected, input["selected"])
	}
}

func TestBasicGetters(t *testing.T) {
	var m Message

	m = Message{Data: 123.0}
	if i, ok := m.DataFloat64(); !ok || i != 123.0 {
		t.Error("Unable to read float64 type data correctly")
	}
	m = Message{Data: 123.0}
	if i, ok := m.DataInt(); !ok || i != 123 {
		t.Error("Unable to read int type data correctly from float64")
	}
	m = Message{Data: true}
	if i, ok := m.DataBool(); !ok || i != true {
		t.Error("Unable to read bool type data correctly")
	}
	m = Message{Data: "123"}
	if i, ok := m.DataString(); !ok || i != "123" {
		t.Error("Unable to read string type data correctly")
	}

}
