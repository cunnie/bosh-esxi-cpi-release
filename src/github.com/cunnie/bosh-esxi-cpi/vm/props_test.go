package vm_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cunnie/bosh-esxi-cpi/vm"
)

var _ = Describe("VMProps", func() {
	Describe("Unmarshal", func() {
		It("picks up ESXi configuration", func() {
			var props VMProps

			err := json.Unmarshal([]byte(`{ "ports": ["6868"],"MemorySizeMB":1024,"NumCpu":2}`), &props)
			Expect(err).ToNot(HaveOccurred())
			Expect(props.MemorySizeMB).To(Equal(int32(1024)))
			Expect(props.NumCpu).To(Equal(int32(2)))
		})
	})
})
