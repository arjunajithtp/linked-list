package singly

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Singly Linked List")
}

var _ = Describe("Singly Linked List", func() {

	Context("when singly linked list is used", func() {

		It("should correctly insert and return the data with Push() and Read() respectively", func() {
			var l List
			l.Push(1, 4, 5, "test", 4)
			actualData := l.Read()
			var testData []interface{}
			// Our linked list is in LIFO order
			testData = append(testData, 4, "test", 5, 4, 1)
			Expect(actualData).To(Equal(testData))
		})

		It("should correctly reverse when Reverse() is used", func() {
			var l List
			l.Push(1, 4, 5, "test", 4)
			l.Reverse()
			actualData := l.Read()
			var testData []interface{}
			// Our linked list is in LIFO order
			testData = append(testData, 1, 4, 5, "test", 4)
			Expect(actualData).To(Equal(testData))
		})

		It("should correctly pop the recent entry when Pop() is used", func() {
			var l List
			l.Push(1, 4, 5, "test", 4)
			popped := l.Pop()
			actualData := l.Read()
			var testData []interface{}
			// Our linked list is in LIFO order
			testData = append(testData, "test", 5, 4, 1)
			Expect(popped).To(Equal(4))
			Expect(actualData).To(Equal(testData))
		})
	})
})
