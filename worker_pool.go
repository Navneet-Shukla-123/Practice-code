package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Way to process the email task
func(t *EmailTask)Process(){
	fmt.Printf("Sending Email to %s\n",t.Email)
	// simulate the time consuming process
	time.Sleep(2 *time.Second)
}

//Image processing task
type ImageProcessingTask struct{
	ImageUrl string
}

// way to process the image

func(t *ImageProcessingTask)Process(){
	fmt.Printf("Processing the image %s\n",t.ImageUrl)
	// simulate the time consuming task
	time.Sleep(5 *time.Second)
}

// Worker pool definition

type WorkerPool struct{
	Tasks []Task
	concurrency int // number of goroutines running
	tasksChan chan Task
	wg sync.WaitGroup
}

// Functions to execute the worker pool

func(wp *WorkerPool)worker(){
	for task:=range wp.tasksChan{
		task.Process()
		wp.wg.Done()
	}
}

func(wp  *WorkerPool)Run(){
	// Initialize the tasks channel

	wp.tasksChan=make(chan Task,len(wp.Tasks))

	// start workers

	for i:=0;i<wp.concurrency;i++{
		go wp.worker()
	}

	//Send tasks to the Tasks channel
	wp.wg.Add(len(wp.Tasks))
	for _,task:=range wp.Tasks{
		wp.tasksChan<-task
	}
	close(wp.tasksChan)

	//wait for all tasks to finish
	wp.wg.Wait()
}