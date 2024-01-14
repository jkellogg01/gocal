package main

import (
	"fmt"
	"regexp"
	"strings"
)

// since almost all iCalendar information units are separated by line breaks,
// the most logical approach seems to be to parse each line as its own "token"
type ContentLine struct {
	Name   string
	Value  string
	Params []*Param
}

type Param struct {
	Name  string
	Value string
}

var ErrInvalidContentLine error = fmt.Errorf("failed to parse content line")
var ErrInvalidParam error = fmt.Errorf("failed to parse parameter")

func Scan(line string) (*ContentLine, error) {
	contentLine, err := regexp.Compile(`^(.*?)((?:;\w*=.*?)*):(.*)$`)
	if err != nil {
		return nil, err
	}

	lineData := contentLine.FindStringSubmatch(line)
	if len(lineData) <= 1 {
		return nil, ErrInvalidContentLine
	}

	name := lineData[1]
	value := lineData[len(lineData)-1]
	params := make([]*Param, 0)
	for _, rawParam := range strings.Split(lineData[2], ";")[1:] {
		name, value, found := strings.Cut(rawParam, "=")
		if !found {
			return nil, ErrInvalidParam
		}
		params = append(params, &Param{
			Name:  name,
			Value: value,
		})
	}
	return &ContentLine{
		Name:   name,
		Value:  value,
		Params: params,
	}, nil
}

func (cl *ContentLine) ToString() string {
	result := new(strings.Builder)
	result.WriteString(fmt.Sprintf("Name:     %s\n", cl.Name))
	result.WriteString(fmt.Sprintf("Value:    %s\n", cl.Value))
	for i, param := range cl.Params {
		result.WriteString(fmt.Sprintf(
			"Param %2d: %s=%s\n",
			i, param.Name, param.Value,
		))
	}
	return result.String()
}
