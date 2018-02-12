/**
* Author : Siddhant Goenka
*/
package main

import "sync"

type manager struct	{
	filename string
	wcount uint64
}

var (
	finfo *fileInfo					// file information pointer
	workers []*worker				// all worker details
	lock sync.WaitGroup				// object to wait for all the workers
)


func (m *manager) start()	{

	/* Gets file size information to schedule worker */
	finfo = &fileInfo{filename:m.filename}
	finfo.info()

	m.scheduleWorkers()
	/* Wait for workers to finish calculating partial sums */
	lock.Wait()

	/* Get response from all workers, since wait is over */
	response := make([]string, m.wcount);

	for _, w := range workers	{
		res := <- w.ch
		response[w.id] = res
		logger.Infof("Received JSON from W%d : %s", w.id, res)
		close(w.ch)
	}//end of loop

	m.finalizeSum()

}//end of method


func (m *manager) scheduleWorkers()	{

	offset :=  round( float64(finfo.filesize)/float64(m.wcount) )
	logger.Debugf("Offset for each worker: %dbytes", offset)

	var i uint64 = 0
	var count uint8 = 0
	workers = make([]*worker, m.wcount);


	for i=0; i < finfo.filesize;		{

		end := min(i+offset-1, finfo.filesize-1)
		di := dataInfo{finfo.filename, i, end}
		str := di.marshal()

		ch := make(chan string, len(str))						// buffer length = 100 bytes
		w := &worker{ch:ch, id: count, lock:&lock}
		workers[count] = w
		ch <- str
		// DO NOT close(ch), need to get back the response

		lock.Add(1)
		w.start()

		i += offset
		count++

	}//end of loop

}//end of method


func (m *manager) finalizeSum() uint64	{
	logger.Info("Finalizing sum...")
	//TODO
	return 0
}//end of method


func min(a uint64, b uint64 ) uint64	{
	if a <= b {
		return a
	}else	{
		return b
	}
}