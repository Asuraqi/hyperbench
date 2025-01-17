package recorder

import (
	"bytes"
	"github.com/meshplus/hyperbench/common"
	"strings"
	"sync"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	times := 100
	b.Run("bytes", func(b *testing.B) {
		var pool sync.Pool
		pool.New = func() interface{} {
			return bytes.NewBuffer(nil)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := pool.Get().(*bytes.Buffer)
			buf.Reset()
			for j := 0; j < times; j++ {
				buf.WriteString("1234567890")
			}
			_ = buf.String()
			pool.Put(buf)
		}
	})

	b.Run("bytes-no-str", func(b *testing.B) {
		var pool sync.Pool
		pool.New = func() interface{} {
			return bytes.NewBuffer(nil)
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := pool.Get().(*bytes.Buffer)
			buf.Reset()
			for j := 0; j < times; j++ {
				buf.WriteString("1234567890")
			}
			pool.Put(buf)
		}
	})

	b.Run("builder", func(b *testing.B) {
		var pool sync.Pool
		for i := 0; i < 4; i++ {
			pool.Put(&strings.Builder{})
		}
		pool.New = func() interface{} {
			return &strings.Builder{}
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			buf := pool.Get().(*strings.Builder)
			buf.Reset()
			for j := 0; j < times; j++ {
				buf.WriteString("1234567890")
			}
			_ = buf.String()
			pool.Put(buf)
		}
	})

}

func TestLogTile(t *testing.T) {
	newLogProcessor(common.GetLogger("test")).logTitle()
}
