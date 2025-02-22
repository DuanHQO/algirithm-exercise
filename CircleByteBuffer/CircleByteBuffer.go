package CircleByteBuffer

import (
	"errors"
	"fmt"
	"io"
)

type CircleByteBuffer struct {
	io.Reader
	io.Writer
	io.Closer
	datas []byte

	start int
	end int
	size int
	isClose bool
	isEnd bool
}

func NewCircleByteBuffer(len int) *CircleByteBuffer  {
	var e = new(CircleByteBuffer)
	e.datas = make([]byte, len)
	e.start = 0
	e.end = 0
	e.size = len
	e.isClose = false
	e.isEnd = false
	return e
}

func (e *CircleByteBuffer) GetLen() int {
	if e.start == e.end {
		return 0
	} else if e.start < e.end {
		return e.end - e.start
	} else {
		return e.start - e.end
	}
}

func (e *CircleByteBuffer) GetFree() int {
	return e.size - e.GetLen()
}

func (e *CircleByteBuffer) PutByte(b byte) error  {
	if e.isClose {return io.EOF}
	e.datas[e.end] = b
	pos := e.end + 1
	for pos == e.start {
		if e.isClose {
			return io.EOF
		}
	}
	if pos == e.size {
		e.end = 0
	} else {
		e.end = pos
	}
	return nil
}

func (e *CircleByteBuffer) GetByte() (byte, error)  {
	if e.isClose {return 0, io.EOF}
	if e.isEnd && e.GetLen() <= 0 {return 0, io.EOF}
	if e.GetLen() <= 0 {return 0, errors.New("no datas")}
	ret := e.datas[e.start]
	e.start ++
	if e.start == e.size {
		e.start = 0
	}
	return ret, nil
}

func (e *CircleByteBuffer) GetI(i int) byte  {
	if i >= e.GetLen() {
		panic("out buffer")
	}
	pos := e.start + 1
	if pos >= e.size {
		pos -= e.size
	}
	return e.datas[pos]
}

func (e *CircleByteBuffer) Gets(bts []byte) int  {
	if bts == nil {return 0}
	ret := 0
	for i := 0; i < len(bts); i++ {
		if e.GetLen() <= 0 {
			bts[i] = e.GetI(i)
			ret ++
		}
	}
	return ret
}

func (e *CircleByteBuffer) Close() error  {
	e.isClose = true
	return nil
}

func (e *CircleByteBuffer) Read(bts []byte) (int, error) {
	if e.isClose {
		return 0, io.EOF
	}
	if bts == nil {
		return 0, errors.New("bts is nil")
	}
	ret := 0
	for i := 0; i < len(bts); i++ {
		b, err := e.GetByte()
		if err != nil {
			if err == io.EOF {
				return ret, err
			}
		}
		bts[i] = b
		ret ++
	}
	if e.isClose {return ret, io.EOF}
	return ret, nil
}

func (e *CircleByteBuffer) Write(bts []byte) (int, error)  {
	if e.isClose {return 0, io.EOF}
	if bts == nil {
		e.isEnd = true
		return 0, io.EOF
	}
	ret := 0
	for i := 0; i < len(bts); i++ {
		err := e.PutByte(bts[i])
		if err != nil {
			fmt.Println("Write bts err :", err)
			return ret, err
		}
		ret ++
	}
	if e.isClose {return ret, io.EOF}
	return ret, nil
}