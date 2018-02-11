package main

import (
	log "github.com/alexcesaro/log/stdlog"
	"strconv"
	"math"
	"os"
	"strings"
)

var logger = log.GetFromFlags()

func main() {

	/* Parsing command line arguments */
	fname := os.Args[1]
	wc := os.Args[2]

	if strings.Contains(fname, "-log=") {
		fname = os.Args[2]
		wc = os.Args[3]
	}

	tcount, _ := strconv.Atoi(wc)

	/* Starting manager... */
	m := manager {filename:fname, wcount:uint64(tcount)}
	m.start()

	logger.Close()

}//end of main method


/* Rounds up the value of a float64 */
func round(val float64) uint64	{
	return uint64( math.Floor(val) )
}


/* Checks for error and prints if found */
func check(err error) bool {
	if err != nil {
		logger.Alert(err)
		return false
	}
	return true
}