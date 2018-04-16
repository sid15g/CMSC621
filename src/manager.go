/**
* Author : Siddhant Goenka
*/
package main

import (
	"sync"
	"fmt"
)

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
	total := m.finalizeSum()

	fmt.Println()
	logger.Info("Total sum is:", total)

}//end of method


func (m *manager) scheduleWorkers()	{

	offset :=  round( float64(finfo.filesize)/float64(m.wcount) )
	logger.Debugf("Offset for each worker: %dbytes", offset)

	var i uint64 = 0
	var count uint8 = 0
	var settle uint64 = 0
	workers = make([]*worker, m.wcount);
	extra := finfo.filesize - offset*m.wcount

	for i=0; i < finfo.filesize;		{

		if( extra > 0 )	{
			settle = 1
			extra--
		}else {
			settle = 0
		}

		end := min(i+offset-1+settle, finfo.filesize-1)
		di := dataInfo{finfo.filename, i, end}
		str := di.marshal()

		ch := make(chan string, len(str))
		w := &worker{ch:ch, id: count, lock:&lock}
		workers[count] = w
		ch <- str
		// DO NOT close(ch), need to get back the response

		lock.Add(1)
		go w.start()

		i += (offset+settle)
		count++

	}//end of loop

}//end of method


/* Get response from all workers, since wait is over, and return the final sum */
func (m *manager) finalizeSum() uint64	{

	logger.Info("Finalizing sum...")
	response := make([]*result, m.wcount);

	for _, w := range workers	{

		res := <- w.ch

		if len(res) > 0	{
			r := &result{}
			logger.Infof("Received JSON from W%d : %s", w.id, res)
			r.unmarshal(res)
			response[w.id] = r
			close(w.ch)
		}else {
			logger.Errorf("Unknown Response from W%d : %s", w.id, res)
		}

	}//end of worker loop


	var wid int16 = 1
	var length int16 = int16(len(response))

	var total_sum uint64 = response[0].Value
	var prev_suffix string = response[0].Suffix

	if response[0].Chunk != 0	{
		prev_suffix = fmt.Sprintf("%d", response[0].Chunk)
	}

	if len(response[0].Prefix) > 0	{
		ip := ParseInt(response[0].Prefix)
		logger.Info("Initial-Prefix Added:", ip)
		total_sum += ip
	}

	for ; wid<length; wid++ {

		r := response[wid]

		if wid > 0	{
			/* Check previous response also for suffix and prefix */

			if len(prev_suffix)>0 && len(r.Prefix)>0 {
				//sf := ParseInt(prev_suffix)
				//pr := ParseInt(r.Prefix)
				//temp := powOfTen(r.Prefix)
				//num := (sf * temp ) + pr
				temp := fmt.Sprintf("%s%s", prev_suffix, r.Prefix)
				num := ParseInt(temp)
				logger.Info("Suffix-Prefix Merged:", num)
				total_sum += num
			}else if r.Chunk > 0 {
				//sf := ParseInt(prev_suffix)
				//temp := powOfTen( fmt.Sprintf("%d", r.Chunk) )
				//num := (sf * temp) + r.Chunk
				temp := fmt.Sprintf("%s%d", prev_suffix, r.Chunk)
				num := ParseInt(temp)
				logger.Infof("Chunk combined : %d", num)
				r.Suffix = fmt.Sprintf("%d", num)
			}else if len(prev_suffix) > 0	{
				sf := ParseInt(prev_suffix)
				logger.Info("Suffix Added to sum:", sf)
				total_sum += sf
			}else if len(r.Prefix) > 0 {
				pr := ParseInt(r.Prefix)
				logger.Info("Prefix Added to sum:", pr)
				total_sum += pr
			}

			total_sum += r.Value
			prev_suffix = r.Suffix

		}else {
			logger.Errorf("Unknown Worker ID W%d ", wid)
		}

	}//end of result loop

	if len(prev_suffix) > 0	{
		sf := ParseInt(prev_suffix)
		logger.Info("Suffix Added to sum:", sf)
		total_sum += sf
	}

	return total_sum

}//end of method


func min(a uint64, b uint64 ) uint64	{
	if a < b {
		return a
	}else	{
		return b
	}
}