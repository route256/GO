package main

import "fmt"

type RingBuffer struct {
	inSinkChannel chan int

	outSinkChannel chan int
}

func NewRingBuffer(
	bufferSize int,
) *RingBuffer {
	rb := &RingBuffer{
		inSinkChannel:  make(chan int),
		outSinkChannel: make(chan int, bufferSize),
	}

	go rb.run()

	return rb
}

func (rb *RingBuffer) run() {
	for value := range rb.inSinkChannel {
		select {
		case rb.outSinkChannel <- value:
		default:
			select {
			case v := <-rb.outSinkChannel:
				fmt.Println("Dropped frame", v)
			default:
			}

			rb.outSinkChannel <- value
		}
	}

	close(rb.outSinkChannel)
}

func (rb *RingBuffer) Push(value int) {
	rb.inSinkChannel <- value
}

func (rb *RingBuffer) Close() {
	close(rb.inSinkChannel)
}

func (rb *RingBuffer) Consume() <-chan int {
	return rb.outSinkChannel
}

func main() {
	ringBuffer := NewRingBuffer(4)

	go func() {
		for i := 0; i < 100; i++ {
			ringBuffer.Push(i)
		}

		ringBuffer.Close()
	}()

	for v := range ringBuffer.Consume() {
		fmt.Println("Consumed frame", v)
	}
}
