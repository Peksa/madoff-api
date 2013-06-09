package tests

import "github.com/robfig/revel"

type ReceiptsTest struct {
	revel.TestSuite
}

func (t ReceiptsTest) Before() {
	println("Set up")
}

func (t ReceiptsTest) TestThatReceiptListWorks() {
	t.Get("/receipts")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t ReceiptsTest) TestThatReceiptShowtWorks() {
	t.Get("/receipts/2")
	t.AssertOk()
	t.AssertContentType("application/json")
}

func (t ReceiptsTest) After() {
	println("Tear down")
}
