package logging_test

import (
	. "github.com/elos/server/util/logging"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Logger", func() {
	Describe("SetLog", func() {
		It("Should set NullLog)", func() {
			SetLog(NullLog)
			Expect(Log).To(Equal(NullLog))
		})

		It("Should set StdOutLog)", func() {
			SetLog(StdOutLog)
			Expect(Log).To(Equal(StdOutLog))
		})

		It("Should set FileLog)", func() {
			SetLog(FileLog)
			Expect(Log).To(Equal(FileLog))
		})
	})

	Describe("FormatService", func() {
		It("Capitalizes the word given to it", func() {
			s := "service"
			Expect(FormatService(s)).To(BeEquivalentTo("SERVIC"))
			s = "SERVICE"
			Expect(FormatService(s)).To(BeEquivalentTo("SERVIC"))
			s = "Ser    "
			Expect(FormatService(s)).To(BeEquivalentTo("SER   "))
		})
	})

	Describe("FormateLogMessage", func() {
		It("Formats [SERVIC]: message", func() {
			s := "service"
			m := "message"
			Expect(FormatLogMessage(s, m)).To(BeEquivalentTo("[SERVIC]: message"))
			s = "s"
			Expect(FormatLogMessage(s, m)).To(BeEquivalentTo("[S     ]: message"))
		})
	})
})
