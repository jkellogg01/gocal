package main

import (
	"fmt"
	"log"
	"strings"
)

type Component struct {
	Props    map[string]*Property  `json:"props"`
	Children map[string]*Component `json:"children"`
}

type Property struct {
	Value  string              `json:"value"`
	Params map[string][]string `json:"params"`
}

func Compile(data string) (*Component, error) {
	tokens := make([]*ContentLine, 0)
	for i, line := range strings.Split(data, "\r\n") {
		if line == "" {
			log.Printf("Line %d was empty", i)
			continue
		}
		token, err := Scan(line)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	result := new(Stack[Component])
	for _, token := range tokens {
		switch token.Name {
		case "BEGIN":
			result.Push(Component{
				Props:    make(map[string]*Property),
				Children: make(map[string]*Component),
			})
		case "END":
			if result.Length <= 1 {
				break
			}
			comp := result.Pop()
			result.Head.Val.Children[token.Value] = comp
		default:
			props := &result.Peek().Props
			(*props)[token.Name] = &Property{
				Value:  token.Value,
				Params: token.Params,
			}
		}
	}
	return result.Pop(), nil
}

func (c Component) String() string {
	result := new(strings.Builder)

	for key, prop := range c.Props {
		result.WriteString(fmt.Sprintf(
			"%s: %s\n",
			key,
			prop.Value,
		))
		for key, vals := range prop.Params {
			result.WriteString(fmt.Sprintf(
				"\t\t%s=%v\n",
				key,
				vals,
			))
		}
	}

	for _, comp := range c.Children {
		result.WriteString("\t" +
			strings.ReplaceAll(
				comp.String(),
				"\n",
				"\n\t",
			),
		)
	}

	return result.String()
}
