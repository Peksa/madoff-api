package tests

import "github.com/robfig/revel"

type UsersTest struct {
	revel.TestSuite
}

func (t UsersTest) Before() {
	println("Set up")
}

func (t UsersTest) TestThatUserListWorks() {
	t.Get("/users")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t UsersTest) TestThatUserShowWorks() {
	t.Get("/users/3")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t UsersTest) After() {
	println("Tear down")
}
