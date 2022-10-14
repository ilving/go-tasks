package main

type someMethod func(a, b, c string) string

func castT() {
	f := func(p interface{}) {
		if v, ok := p.(interface{ F() }); ok {
			v.F()
		}

		if fun, ok := p.(someMethod); ok {
			var r string
			r = fun("a", "b", "c")
			print(r)
		}
	}

	f(func(d, e, f string) string {
		return ""
	})

}
