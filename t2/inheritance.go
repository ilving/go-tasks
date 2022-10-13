package main

type InhA struct{}

func (i InhA) Func() { println("A") }

type InhB struct{ InhA }

func (i InhB) Func() { println("B") }

type InhC struct{ InhB }

func (i InhC) Func() { println("C") }

func inheritanceTest() {
	c := InhC{}
	c.Func()

	c.InhB.InhA.Func()
	c.InhA.Func()
}
