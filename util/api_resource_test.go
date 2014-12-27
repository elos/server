package util_test

import (
	. "github.com/elos/server/util"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApiResource", func() {
	It("WriteResourceResponse", func() {
		w := httptest.NewRecorder()

		type TestStruct struct {
			This string
			is   int
		}

		testStruct := TestStruct{
			This: "data",
			is:   1,
		}

		WriteResourceResponse(w, 200, testStruct)

		bytesArray, err := ToJSON(testStruct)
		Expect(err).ToNot(HaveOccurred())

		Expect(w.Code).To(Equal(200))
		Expect(w.HeaderMap["Content-Type"][0]).To(Equal("application/json; charset=utf-8"))
		Expect(w.Body.String()).To(Equal(string(bytesArray)))
	})

})
