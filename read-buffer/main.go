package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b = bytes.NewBuffer(make([]byte, 26))
	var texts = []string{
		`Frankness applauded.`,
		`Collected favourite now for for and rapturous repulsive consulted.`,
		`An seems green be wrote again. She add what own only like.`,
		`Tolerably we as extremity exquisite do commanded. Doubtful offended do entrance of landlord moreover is mistress in. Nay was appear entire ladies.`,
	}
	for i := range texts {
		b.Reset()
		b.WriteString(texts[i])
		fmt.Printf("Length:%d\tCapacity:%d\n", b.Len(), b.Cap())
	}
}
