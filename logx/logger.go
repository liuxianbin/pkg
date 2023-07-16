package logx

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"runtime"
	"time"
)

type Fields map[string]any

type Loggerx struct {
	*log.Logger
	fields       Fields   // 公共字段
	callers      []string // 调用栈信息
	callersLevel int      // 调用栈层级
}

func New(w io.Writer, prefix string, flag int) *Loggerx {
	l := log.New(w, prefix, flag)
	return &Loggerx{
		Logger:       l,
		callersLevel: 20,
	}
}

func (l *Loggerx) SetCallersLevel(callersLevel int) *Loggerx {
	l.callersLevel = callersLevel
	return l
}

func (l *Loggerx) Copy() *Loggerx {
	_l := *l
	return &_l
}

func (l *Loggerx) WithFields(f Fields) *Loggerx {
	_l := l.Copy()
	if _l.fields == nil {
		_l.fields = make(Fields)
	}
	for k, v := range f {
		_l.fields[k] = v
	}
	return _l
}

func (l *Loggerx) WithCaller(skip int) *Loggerx {
	_l := l.Copy()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		_l.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return _l
}

func (l *Loggerx) WithCallers() *Loggerx {
	_l := l.Copy()
	pcs := make([]uintptr, l.callersLevel)
	size := runtime.Callers(1, pcs)
	fs := runtime.CallersFrames(pcs[:size])
	callers := make([]string, size)[:0]
	for {
		f, ok := fs.Next()
		callers = append(callers, fmt.Sprintf("%s: %d %s", f.File, f.Line, f.Function))
		if !ok {
			break
		}
	}
	_l.callers = callers
	return _l
}

func (l *Loggerx) Jsonf(format string, v ...any) {
	l.Json(fmt.Sprintf(format, v...))
}

func (l *Loggerx) Json(v ...any) {
	data := make(Fields, len(l.fields)+3)
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = fmt.Sprint(v...)
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}
	body, _ := json.Marshal(data)
	l.Println(string(body))
}
