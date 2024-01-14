package main

type ContentLine struct {
	Name   string
	Params []Param
	Value  string
}

type Param struct {
	Name  string
	Value string
}
