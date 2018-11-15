package trace
// Tracer is the interface that describes and object capable of
// tracing events throughout the code.
import (
  "io"
  "fmt"
)

type nilTracer struct{}
func (t *nilTracer) Trace(a ...interface{}) {}

// Off creates a Tracer that will ignore calls to Trace.
func Off() Tracer {
  return &nilTracer{}
}

type Tracer interface {
  Trace(...interface{})
}

type tracer struct {
  out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
  t.out.Write([]byte(fmt.Sprint(a...)))
  t.out.Write([]byte("\n"))
}

func New(w io.Writer) Tracer {
  return &tracer{out: w}
}
