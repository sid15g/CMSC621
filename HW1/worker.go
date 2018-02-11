/**
* Author : Siddhant Goenka
*/
package main

import (
	"sync"
	"strings"
	"strconv"
)

/* Worker structure to store necessary details */
type worker struct	{
	id uint8
	ch chan string
	fptr *fileInfo
	lock *sync.WaitGroup
}

const space = string(' ')

/* Starts worker thread - use 'go' keyword to start it */
func (w *worker) start()	{
	w.getInfo()
	w.calculateSum()
	w.lock.Done()
}


/* Get JSON from the channel, parse it and create file pointer for worker to read the file */
func (w *worker) getInfo()	{
	di := &dataInfo{}
	strJson := <- w.ch
	logger.Infof(" Worker[%d] JSON received : %s", w.id, strJson)
	di.unmarshal(strJson)
	w.fptr = &fileInfo{filename:di.Filename, start:di.Start, end:di.End}
}//end of method


/* Calculates partial sum and returns prefix, sum, suffix */
func (w *worker) calculateSum() uint64	{

	data := w.fptr.read()
	logger.Debugf(" Worker[%d]: %d bytes data read -> %v", w.id, len(data), data)

	var sum uint64 = 0
	vals := string(data)
	str := strings.Split(vals, space)

	//TODO ignore first and last element, for prefix and suffix

	for _, e := range str {

		num, err := strconv.Atoi(e)

		if len(e)>0 && check(err)	{
			sum += uint64(num)
			//TODO calculate prefix, suffix and partial sum
		}else if len(e) > 0 {
			//TODO print some error
		}else {
			// found space, do not do anything
		}

	}//end of loop

	logger.Infof(" Worker[%d]: Partial Sum %d",w.id, sum )
	return sum
	//TODO return prefix, sum, suffix (in order)

}//end of method