//package error mapper maps an error code to its description depending on the language, and generates the
//appropriate error msg to be returned. [](code+lang) -> msg
package errormapper

import (
	"fmt"
	"strings"
)

type Message map[string]string
type errdef map[int32]Message

//ErrMap defines the correspondance for translation between error code (parameter) and msg (output) depending on the language (parameter)
type ErrMap struct {
	languages []string
	def       map[int32]Message
}

func (e *ErrMap) Get_message(code int32, lang string) (string, int32) {
	n := int32(0)
	if ok, _ := e.Lang_accepted(lang); ok == true {
		if s, ok := e.def[code][lang]; ok {
			n++
			return s, n
		} else {
			return "", n
		}
	} else {
		return "", n
	}
}

func (e *ErrMap) Get_messages(codes []int32, lang string) (string, int32) {
	m := ""
	n := int32(0)
	if ok, _ := e.Lang_accepted(lang); ok == true {
		for k, i := range codes {
			if k > 0 {
				m = fmt.Sprintf("%s, ", m)
			}
			if s, ok := e.def[i][lang]; ok {
				n++
				m = fmt.Sprintf("%s%s", m, s)
			}
		}
	}
	return m, n
}

func (e *ErrMap) Lang_accepted(lang string) (bool, int32) {
	//l := strings.ToUpper(lang)
	for k, l := range e.languages {
		if strings.ToUpper(lang) == strings.ToUpper(l) {
			return true, int32(k)
		}
	}
	return false, -1
}

func (e *ErrMap) Add_languages(langs []string) int32 {
	if len(langs) == 0 {
		return 0
	}
	/*if len(e.languages) == 0 {
		return 0
	}*/
	n3 := int32(0)
	for _, lang := range langs {
		if ok, _ := e.Lang_accepted(lang); ok == false {
			e.languages = append(e.languages, strings.ToUpper(lang))
			n3++
		}
	}
	return n3
}

func (e *ErrMap) Delete_languages(langs []string) int32 {
	if len(langs) == 0 {
		return 0
	}
	if len(e.languages) == 0 {
		return 0
	}
	n3 := int32(0)
	for _, lang := range langs {
		l := strings.ToUpper(lang)
		if ok, ind := e.Lang_accepted(l); ok == true {
			n3++
			e.languages = append(e.languages[:ind], e.languages[ind+1:]...)
		}
	}
	return n3
}

func (e *ErrMap) Add_msg(lang string, code int32, msg string) int32 {
	l := strings.ToUpper(lang)
	if ok, _ := e.Lang_accepted(l); ok == true {
		if len(e.def) > 0 {
			if m1, ok := e.def[code]; ok {
				m1[lang] = msg
				e.def[code] = m1
			} else {
				m1bis := make(Message)
				m1bis[lang] = msg
				e.def[code] = m1bis
			}
		} else {
			m2 := make(errdef)
			m1 := make(Message)
			m1[lang] = msg
			m2[code] = m1
			e.def = m2
		}
		return 1
	} else {
		return 0
	}
}
func (e *ErrMap) Add_error(code int32, mess Message) int32 {
	if len(mess) == 0 {
		return 0
	}
	n := int32(0)
	m1 := make(Message)
	m0 := make(map[int32]Message)
	if len(e.def) > 0 {
		m0 = e.def
	}
	//ok := false
	if m2, ok := m0[code]; ok {
		for key1, val1 := range m2 {
			if ok, _ := e.Lang_accepted(key1); ok == true {
				m1[key1] = val1
				//n++
			}
		}
	}
	for key2, val2 := range mess {
		if ok, _ := e.Lang_accepted(key2); ok == true {
			m1[key2] = val2
			n++
		}
	}
	m0[code] = m1
	e.def = m0
	return n
}

func (e *ErrMap) Delete_msg(code int32, lang string) int32 {
	l := strings.ToUpper(lang)
	if ok, _ := e.Lang_accepted(l); ok == true {
		//e.def[code][l] = strings.Title(msg)
		delete(e.def[code], l)
	} else {
		return 0
	}

	if len(e.def) > 0 {
		res := make(map[int32]Message)
		for key1, val1 := range e.def {
			for key2, val2 := range val1 {
				if m1, ok := res[key1]; ok {
					m1[key2] = val2
					res[key1] = m1
				} else {
					m1bis := make(Message)
					m1bis[key2] = val2
					res[key1] = m1bis
				}
			}
		}
		e.def = res

	} else {
		e.def = nil
	}
	return 1

}

func (e *ErrMap) Delete_error(code int32) int32 {
	n := int32(0)
	if _, ok := e.def[code]; ok == true {
		n = int32(len(e.def[code]))
		delete(e.def, code)
	} else {
		return n
	}

	if len(e.def) > 0 {
		res := make(map[int32]Message)
		for key1, val1 := range e.def {
			for key2, val2 := range val1 {
				if m1, ok := res[key1]; ok {
					m1[key2] = val2
					res[key1] = m1
				} else {
					m1bis := make(Message)
					m1bis[key2] = val2
					res[key1] = m1bis
				}
			}
		}
		e.def = res

	} else {
		e.def = nil
	}
	return n
}
func (e *ErrMap) Delete_msgs_in(lang string) int32 {
	n := int32(0)
	l := strings.ToUpper(lang)
	for k, _ := range (*e).def {
		e.Delete_msg(k, l)
		n++
	}
	return n
}

func (e *ErrMap) Purge_msgs() int32 {
	n := e.Count()
	e.def = nil
	return n
}

func (e *ErrMap) Count() int32 {
	length := int32(0)
	for _, v := range e.def {
		for _, _ = range v {
			//k = ""
			length++
		}
	}
	return length
}

func New() *ErrMap {
	return new(ErrMap)
}
