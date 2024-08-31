package heap

import (
	"fmt"
	"testing"
)

func TestHeap_Push(t *testing.T) {

	tests := []struct {
		name string
		h    *Heap
	}{
		{
			name: "heap push test",
			h:    &Heap{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Push(10)
			fmt.Println(tt.h)
			tt.h.Push(5)
			fmt.Println(tt.h)
			tt.h.Push(1)
			fmt.Println(tt.h)
			tt.h.Push(20)
			fmt.Println(tt.h)
			tt.h.Push(30)
			fmt.Println(tt.h)

			tt.h.Pop()
			fmt.Println(tt.h)

		})
	}
}
