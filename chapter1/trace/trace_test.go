package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buff bytes.Buffer
	tracer := NewTracer(&buff)
	if tracer == nil {
		t.Errorf("no tracer")
	} else {
		tracer.Trace("こんにちは、traceパッケージ")
		if buff.String() != "こんにちは、traceパッケージ\n" {
			t.Errorf("%sという誤った文字列が出力されました", buff.String())
		}
	}

}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("何かのデータ")
}
