package main

type IF interface {
	F()
}

type S struct{ IF }

type S2 struct{}

func (s S2) F() { print("S2") }

type S3 struct{}

func (s S3) F() { print("S3") }

func ifT() {
	s := S{}
	s.F()

	s.IF = S2{}
	s.F()

	s.IF = S3{}
	s.F()
}
