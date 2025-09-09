package main

import (
	"fmt"
	"sync"
)

func runConcurrency() {
	runFoo()
	runFooBar()
}

// 1114. Print in Order
func runFoo() {
	fmt.Println("Foo Concurrency Start")

	foo := newFoo()
	wg := new(sync.WaitGroup)
	wg.Add(3)
	go func() {
		defer wg.Done()
		foo.First(func() { fmt.Println("First") })
	}()
	go func() {
		defer wg.Done()
		foo.Second(func() { fmt.Println("Second") })
	}()
	go func() {
		defer wg.Done()
		foo.Third(func() { fmt.Println("Third") })
	}()
	wg.Wait()

	fmt.Println("Foo Concurrency Finish")
}

type Foo struct {
	two   chan struct{}
	three chan struct{}
}

func newFoo() *Foo {
	return &Foo{
		two:   make(chan struct{}),
		three: make(chan struct{}),
	}
}

func (f *Foo) First(printFirst func()) {
	printFirst()
	close(f.two)
}

func (f *Foo) Second(printSecond func()) {
	<-f.two
	printSecond()
	close(f.three)
}

func (f *Foo) Third(printThird func()) {
	<-f.three
	printThird()
}

/*
1115. Print Foo Bar Alternately

Key Points:

Buffered Channels: The fooChan and barChan channels are buffered with a capacity of 1. This allows the initial signal to be sent without blocking.

Initial Signal: The initial signal fooBar.fooChan <- struct{}{} starts the sequence by allowing the Foo method to proceed.

Alternating Signals: The Foo and Bar methods alternate sending and receiving signals, ensuring that "Foo" and "Bar" are printed alternately.
*/
func runFooBar() {
	fmt.Println("\nFooBar Concurrency Start")
	wg := new(sync.WaitGroup)
	fooBar := newFooBar()
	n := 10
	wg.Add(2)
	go fooBar.Foo(func() { fmt.Print("Foo") }, n, wg)
	go fooBar.Bar(func() { fmt.Print("Bar") }, n, wg)
	// Start the sequence by sending an initial signal to fooChan
	fooBar.fooChan <- struct{}{}
	// Wait for all goroutines to finish
	wg.Wait()
	// Close the channels after the WaitGroup has finished waiting
	// <-fooBar.fooChan // you can do this because the last iteration of Bar will populate this channel
	// close(fooBar.fooChan)
	// close(fooBar.barChan)

	fmt.Println("\nFooBar Concurrency Finish")
}

type FooBar struct {
	fooChan chan struct{}
	barChan chan struct{}
}

func newFooBar() *FooBar {
	return &FooBar{
		fooChan: make(chan struct{}, 1),
		barChan: make(chan struct{}, 1),
	}
}

func (fb *FooBar) Foo(printFoo func(), n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range n {
		<-fb.fooChan
		printFoo()
		fb.barChan <- struct{}{} // send empty struct to signal

	}
}
func (fb *FooBar) Bar(printBar func(), n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range n {
		<-fb.barChan
		printBar()
		if i < n {
			fb.fooChan <- struct{}{} // send empty struct to signal
		}
	}
}
