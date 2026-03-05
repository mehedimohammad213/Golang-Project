package rag

import (
	"fmt"
	"strings"
)

// FormatVectorForPG converts a float32 slice to pgvector string format "[a,b,c,...]".
func FormatVectorForPG(embedding []float32) string {
	var b strings.Builder
	b.WriteString("[")
	for i, v := range embedding {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(fmt.Sprintf("%g", v))
	}
	b.WriteString("]")
	return b.String()
}
