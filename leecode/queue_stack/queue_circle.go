package main

func main() {

}

type MyCircularQueue struct {
	maxSize int
	front   int
	end     int
	queue   []int
	length  int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		k,
		0,
		0,
		make([]int, k),
		0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.length == this.maxSize {
		return false
	}
	this.queue[this.end] = value
	this.end = (this.end + 1) % this.maxSize
	this.length += 1
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.length == 0 {
		return false
	}
	this.front = (this.front + 1) % this.maxSize
	this.length -= 1
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	if this.end == 0 {
		return this.queue[this.maxSize-1]
	}
	return this.queue[this.end-1]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.length == 0
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.length == this.maxSize
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
