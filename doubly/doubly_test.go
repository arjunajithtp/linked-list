package doubly

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Doubly Linked List")
}

var _ = Describe("Doubly Linked List", func() {

	Context("when Doubly linked list is used", func() {

		It("should correctly insert and return the data with Push() and ReadLIFO()/ReadFIFO() respectively", func() {
			var l List
			l.Push(1, 4, 5, "test", 4)
			actualLIFOData := l.ReadLIFO()
			actualFIFOData := l.ReadFIFO()
			var testLIFOData, testFIFOData []interface{}
			testLIFOData = append(testLIFOData, 4, "test", 5, 4, 1)
			testFIFOData = append(testFIFOData, 1, 4, 5, "test", 4)
			Expect(actualLIFOData).To(Equal(testLIFOData))
			Expect(actualFIFOData).To(Equal(testFIFOData))
		})

		It("should correctly pop the recent entry when PopHead() is used", func() {
			var l List
			l.Push(1, 4, 5, "test", 4)
			popped := l.PopHead()
			actualLIFOData := l.ReadLIFO()
			actualFIFOData := l.ReadFIFO()
			var testLIFOData, testFIFOData []interface{}
			testLIFOData = append(testLIFOData, "test", 5, 4, 1)
			testFIFOData = append(testFIFOData, 1, 4, 5, "test")
			Expect(popped).To(Equal(4))
			Expect(actualLIFOData).To(Equal(testLIFOData))
			Expect(actualFIFOData).To(Equal(testFIFOData))
		})

		It("should correctly pop the first entry when PopTail() is used", func() {
			var l List
			l.Push(1, 4, 5, "test", 4)
			popped := l.PopTail()
			actualLIFOData := l.ReadLIFO()
			actualFIFOData := l.ReadFIFO()
			var testLIFOData, testFIFOData []interface{}
			testLIFOData = append(testLIFOData, 4, "test", 5, 4)
			testFIFOData = append(testFIFOData, 4, 5, "test", 4)
			Expect(popped).To(Equal(1))
			Expect(actualLIFOData).To(Equal(testLIFOData))
			Expect(actualFIFOData).To(Equal(testFIFOData))
		})
	})
})
