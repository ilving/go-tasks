package main

type IRepo interface {
	Get()
}

type Service struct {
	r IRepo
}

func (s Service) Get() { print("SVC") }

func svcConstructor(r IRepo) *Service {
	return &Service{
		r: r,
	}
}

type repo struct {
	a int
}

func (r *repo) Get() { print("REPO") }

func t() {
	type r2 struct{ repo }
	r2v := r2{}
	r2v.Get()
	print(r2v.a)

	s := svcConstructor(&r2v)
	s.Get()

	print(s)
}
