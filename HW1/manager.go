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
	m.finalizeSum()

}//end of method


func (m *manager) scheduleWorkers()	{

	offset :=  round( float64(finfo.filesize)/float64(m.wcount-1) )
	logger.Debugf("Offset for each worker: %dbytes", offset)

	var i uint64 = 0
	var count uint8 = 0
	workers = make([]*worker, m.wcount);


	for i=0; i <= finfo.filesize;		{

		end := min(i+offset-1, finfo.filesize-1)
		di := dataInfo{finfo.filename, i, end}
		str := di.marshal()

		ch := make(chan string, 100)						// buffer length = 100 bytes
		w := &worker{ch:ch, id: count, lock:&lock}
		workers[count] = w
		ch <- str
		close(ch)

		lock.Add(1)
		go w.start()

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