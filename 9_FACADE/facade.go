package main

import "fmt"

func main() {

	// we hide a lot of implementation details of buffer and viewports behind Console facade
	c := NewConsole()
	u := c.GetCharacterAt(1)
	fmt.Println(u)
}

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width: width, height: height, buffer: make([]rune, width*height)}
}

func (b *Buffer) At(idx int) rune {
	return b.buffer[idx]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(idx int) rune {
	return v.buffer.At(v.offset + idx)
}

type Console struct {
	buffer    []*Buffer
	viewports []*Viewport
	offset    int
}

// default implementation of Console
func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{buffer: []*Buffer{b}, viewports: []*Viewport{v}, offset: 0}
}

func (c *Console) GetCharacterAt(idx int) rune {
	return c.viewports[0].GetCharacterAt(idx)
}
