package main

import (
	"encoding/json"
	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

// подсказки
//https://github.com/mailru/easyjson
//https://github.com/sadlil/experiments/blob/master/go/easyjson/main.go#L12
//https://stackoverflow.com/questions/40587860/using-easyjson-with-golang

import (
	//"encoding/json"
	"fmt"
	//"github.com/mailru/easyjson"
	//"github.com/mailru/easyjson/jlexer"
	//"github.com/mailru/easyjson/jwriter"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

//easyjson:json
type User struct {
	Browser []string `json:"browsers"`
	Company string `json:"company"`
	Country string `json:"country"`
	Email string `json:"email"`
	Job string `json:"job"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	//SlowSearch(out)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileContents), "\n")

	var users []User
	for _, line := range lines {
		user := &User{}
		user.UnmarshalJSON([]byte(line))
		users = append(users, *user)
		//fmt.Println(users)
	}

	r := regexp.MustCompile("@")
	seenBrowsers := []string{}
	uniqueBrowsers := 0
	foundUsers := ""

	for i, user := range users {

		isAndroid := false
		isMSIE := false

		browsers := user.Browser
		for _, browserRaw := range browsers {
			browser := browserRaw
			if strings.Contains(browser, "Android") {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
			//if ok, err := regexp.MatchString("Android", browser); ok && err == nil {
			//	isAndroid = true
			//	notSeenBefore := true
			//	for _, item := range seenBrowsers {
			//		if item == browser {
			//			notSeenBefore = false
			//		}
			//	}
			//	if notSeenBefore {
			//		// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
			//		seenBrowsers = append(seenBrowsers, browser)
			//		uniqueBrowsers++
			//	}
			//}
		}

		for _, browserRaw := range browsers {
			browser := browserRaw
			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
			//if ok, err := regexp.MatchString("MSIE", browser); ok && err == nil {
			//	isMSIE = true
			//	notSeenBefore := true
			//	for _, item := range seenBrowsers {
			//		if item == browser {
			//			notSeenBefore = false
			//		}
			//	}
			//	if notSeenBefore {
			//		// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
			//		seenBrowsers = append(seenBrowsers, browser)
			//		uniqueBrowsers++
			//	}
			//}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := r.ReplaceAllString(user.Email, " [at] ")
		foundUsers += fmt.Sprintf("[%d] %s <%s>\n", i, user.Name, email)
	}

	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))
}


func easyjson3486653aDecodeHw3BenchPack2(in *jlexer.Lexer, out *User) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "browsers":
			if in.IsNull() {
				in.Skip()
				out.Browser = nil
			} else {
				in.Delim('[')
				if out.Browser == nil {
					if !in.IsDelim(']') {
						out.Browser = make([]string, 0, 4)
					} else {
						out.Browser = []string{}
					}
				} else {
					out.Browser = (out.Browser)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Browser = append(out.Browser, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "company":
			out.Company = string(in.String())
		case "country":
			out.Country = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "phone":
			out.Phone = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3486653aEncodeHw3BenchPack2(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"browsers\":"
		out.RawString(prefix[1:])
		if in.Browser == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Browser {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"company\":"
		out.RawString(prefix)
		out.String(string(in.Company))
	}
	{
		const prefix string = ",\"country\":"
		out.RawString(prefix)
		out.String(string(in.Country))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"phone\":"
		out.RawString(prefix)
		out.String(string(in.Phone))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3486653aEncodeHw3BenchPack2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3486653aEncodeHw3BenchPack2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3486653aDecodeHw3BenchPack2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3486653aDecodeHw3BenchPack2(l, v)
}
