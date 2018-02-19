/**
* Author : Siddhant Goenka
*/
package main

import (
	"sync"
	"strings"
)

/* Worker structure to store necessary details */
type worker struct	{
	id uint8
	ch chan string
	fptr *fileInfo
	lock *sync.WaitGroup
}

const (
	space = string(' ')
	newln = string('\n')
)

/* Starts worker thread - use 'go' keyword to start it */
func (w *worker) start()	{

	w.getInfo()
	res := w.calculateSum()

	w.lock.Done()
	w.ch <- res
	logger.Debugf(" Worker[%d] Partial sum sent : %s", w.id, res)

}//end of method


/* Get JSON from the channel, parse it and create file pointer for worker to read the file */
func (w *worker) getInfo()	{
	di := &dataInfo{}
	strJson := <- w.ch
	logger.Infof(" Worker[%d] JSON received : %s", w.id, strJson)
	di.unmarshal(strJson)
	w.fptr = &fileInfo{filename:di.Filename, start:di.Start, end:di.End}
}//end of method


/* Calculates partial sum and returns prefix, sum, suffix */
func (w *worker) calculateSum() string	{

	byts := w.fptr.read()
	data := strings.Replace(string(byts), newln, space, -1)
	logger.Infof(" Worker[%d]: data -> [%s]", w.id, string(data))

	var sum uint64 = 0
	vals := string(data)

	hasSpace := strings.Contains(vals, space)

	if hasSpace == false {
		num := ParseInt(vals)
		r := &result{Chunk: num}
		return r.marshal()
	}else {
		str := strings.Split(vals, space)
		last := len(str)-1
		r := &result{Prefix:str[0], Suffix: str[last]}

		for i:=1; i<last; i++ {

			e := str[i]

			if len(e) > 0 {
				num := ParseInt(e)
				sum += num
			}else {
				// found space, do not do anything
			}

		}//end of loop

		logger.Debugf(" Worker[%d]: Partial Sum %d",w.id, sum )
		r.finalize(sum)
		return r.marshal()
	}//end of space check

}//end of method


func (r *result) finalize(sum uint64)	{
	r.Value = sum
	//TODO
}//end of method