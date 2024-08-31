package hashtable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable_Add(t *testing.T) {
	type args struct {
		key []string
	}
	tests := []struct {
		name string
		ht   *HashTable
		args args
	}{
		{
			name: "test hash table",
			ht:   initHashTable(),
			args: args{key: []string{"ERIK", "sams", "erik", "listen", "cric", "trio"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, key := range tt.args.key {
				tt.ht.Add(key)
			}

			assert.True(t, tt.ht.Search("erik"))

			tt.ht.Delete("listen")

			assert.False(t, tt.ht.Search("listen"))
		})
	}
}
