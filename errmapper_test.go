package errormapper

import (
	"testing"
)

func TestLangAccepted(t *testing.T) {
	m := New()
	m.languages = []string{"FR", "ENG", "ESP"}
	tests := []struct {
		lang string
		ok   bool
		ind  int
	}{
		{"fr", true, 0},
		{"EnG", true, 1},
		{"eSP", true, 2},
		{"ESPP", false, -1},
	}

	for i, test := range tests {
		accepted, j := m.Lang_accepted(test.lang)
		if accepted != test.ok || test.ind != j {
			t.Errorf("test %d: expected %b received %b", i, test.ok, accepted)
		}
	}
}

func TestAddLang(t *testing.T) {
	m := New()
	//m.Add_language([]string{"FR", "Eng"})
	tests := []struct {
		lang  []string
		toAdd bool
	}{
		{[]string{"FR"}, true}, {[]string{"ENG"}, true}, {[]string{"ESP", "ITA"}, false},
	}
	n := 0
	for i, test := range tests {
		if test.toAdd == true {
			n = n + m.Add_languages(test.lang)
		}
		//t.Errorf("%v", m.languages)
		//accepted, _ := m.Lang_accepted(test.lang)
		if len(m.languages) != n {
			t.Errorf("test %d: expected %b received %b", i, n, len(m.languages))
		}
	}
}

func TestDeleteLanguages(t *testing.T) {
	m := ErrMap{languages: []string{"FR", "ENG", "ESP"}}
	m.Delete_languages([]string{"FR", "ENG"})
	if len(m.languages) != 1 {
		t.Errorf("test %d: expected %d remaining language, but counted %d", 0, 1, len(m.languages))
	}
	if m.languages[0] != "ESP" {
		t.Errorf("test %d: expected '%s' value, but found '%s'", 0, "ESP", m.languages[0])
	}
}

func TestAddMessage(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
	}

	for k, test := range tests {
		m.Add_msg(test.lang, test.code, test.message)
		if m.def[test.code][test.lang] != test.message {
			t.Errorf("test %d: expected message '%s' received '%s'", k, test.message, m.def[test.code][test.lang])
		}
	}
}

func TestDeleteMsg(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
	}

	for _, test := range tests {
		m.Add_msg(test.lang, test.code, test.message)
	}
	L := m.Count()
	//t.Errorf("%v", m.def)
	for k, test := range tests {
		m.Delete_msg(test.code, test.lang)
		L = L - 1
		//t.Errorf("%v", m.def)
		if m.Count() != L {
			t.Errorf("test %d: expected %d message(s) left,counted %d", k, L, m.Count())
		}
	}
	m.Add_msg("FR", 12, "le message ultime après purge")
	//t.Errorf("%v", m.def)
}

func TestDeleteError(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
	}

	for _, test := range tests {
		m.Add_msg(test.lang, test.code, test.message)
	}
	L := m.Count()
	//t.Errorf("%v", m.def)
	for k, test := range tests {
		nsuppr := m.Delete_error(test.code)
		L = L - nsuppr
		//t.Errorf("%v", m.def)
		if m.Count() != L {
			t.Errorf("test %d: expected %d message(s) left,counted %d", k, L, m.Count())
		}
	}
	m.Add_msg("FR", 12, "le message ultime après purge")
	//t.Errorf("%v", m.def)
}

func TestDeleteMsgsIn(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
	}

	for _, test := range tests {
		m.Add_msg(test.lang, test.code, test.message)
	}
	L := m.Count()
	//t.Errorf("%v", m.def)
	nres := m.Delete_msgs_in("FR")
	//t.Errorf("%v", m.def)
	if m.Count() != L-nres {
		t.Errorf("expected %d message(s) left,counted %d", L-nres, m.Count())
	}
}

func TestCount(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
	}

	n := 0
	for _, test := range tests {
		ok := m.Add_msg(test.lang, test.code, test.message)
		n = n + ok
		if m.Count() != n {
			t.Errorf("expected %d message(s),counted %d", n, m.Count())
		}
	}
}

func TestAddError(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		mess     Message
		code     int
		expected int
	}{
		{Message{"FR": "le premier message", "ENG": "the first message"}, 1, 2},
		{Message{"FR": "le second message"}, 2, 1},
		{Message{"ESP": "el secundo"}, 2, 0},
	}
	//t.Errorf("%v", m.def)
	for k, test := range tests {
		res := m.Add_error(test.code, test.mess)
		//t.Errorf("%v", m.def)
		if res != test.expected {
			t.Errorf("test %d: expected message '%s' added,received '%s'", k, test.expected, res)
		}
	}
}

func TestGetMsg(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
		{"FR", 1, "le premier message màj"},
	}

	for k, test := range tests {
		m.Add_msg(test.lang, test.code, test.message)
		s, n := m.Get_message(test.code, test.lang)
		if n == 0 || s != test.message {
			t.Errorf("test %d: expected message '%s' received '%s'", k, test.message, s)
		}
	}
}

func TestGetMsgs(t *testing.T) {
	m := New()
	m.Add_languages([]string{"FR", "ENG"})
	tests := []struct {
		lang    string
		code    int
		message string
	}{
		{"FR", 1, "le premier message"},
		{"ENG", 1, "the first message"},
		{"FR", 2, "le second message"},
		{"FR", 1, "le premier message màj"},
	}

	for _, test := range tests {
		m.Add_msg(test.lang, test.code, test.message)
	}
	s, n := m.Get_messages([]int{1, 2}, "FR")
	if n == 0 || s != "le premier message màj, le second message" {
		t.Errorf("expected message '%s' received '%s'", "le premier message màj, le second message", s)
	}
}
