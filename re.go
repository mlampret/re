package re

import (
	"regexp"
)

func New() *Re {
	return &Re{}
}

func Bytes(bytes []byte) *Re {
	return &Re{text: string(bytes)}
}

func String(str string) *Re {
	return &Re{text: str}
}

func Pattern(pat string) *Re {
	return &Re{pattern: pat}
}

func Matches(s, p string) bool {
	return String(s).Pattern(p).Matches()
}

func Submatch(s, p string, id int) string {
	return String(s).Pattern(p).Submatch(id)
}

func Replace(s, p, r string) string {
	return String(s).Pattern(p).Replace(r)
}

type Re struct {
	text       string
	pattern    string
	goRegexp   *regexp.Regexp
	submatches []string
}

func (re *Re) Bytes(bytes []byte) *Re {
	re.text = string(bytes)
	return re
}

func (re *Re) String(str string) *Re {
	re.text = str
	return re
}

func (re *Re) Pattern(pat string) *Re {
	re.pattern = pat
	re.goRegexp = regexp.MustCompile(re.pattern)
	return re
}

func (re *Re) Matches() bool {
	re.goRegexp = regexp.MustCompile(re.pattern)
	return re.goRegexp.MatchString(re.text)
}

func (re *Re) Submatch(id int) string {
	if len(re.submatches) < 1 {
		re.submatches = re.goRegexp.FindStringSubmatch(re.text)
	}
	if len(re.submatches) > id && id >= 0 {
		return re.submatches[id]
	} else {
		return ""
	}
}
func (re *Re) Replace(repl string) string {
	goRe := regexp.MustCompile(re.pattern)
	return goRe.ReplaceAllString(re.text, repl)
}
