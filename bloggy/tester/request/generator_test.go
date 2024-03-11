package request

import (
	"testing"
)

func BenchmarkSignUp(b *testing.B) {

	for key := 0; key < b.N; key++ {
		SignUp()
	}

}
