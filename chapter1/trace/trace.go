package trace

import (
	"fmt"
	"io"
)

type tracer struct {
	out io.Writer //io.Writerはインターフェース。つまり、outはWrite()メソッドを持つオブジェクト出ないといけない
}

type Tracer interface {
	Trace(...interface{})
}

// interface Tracerを返す
// 引数に渡すのは、bytes.Bufferでも何でもOK(write()メソッドを持っていれば)
func NewTracer(a io.Writer) Tracer {
	return &tracer{out: a}
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...interface{}) {}

func Off() Tracer {
	return &nilTracer{}
}
