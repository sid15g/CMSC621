/**
* Author : Siddhant Goenka
*/
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
	wc := os.Args[1]
	fname := os.Args[2]

	if strings.Contains(wc, "-log=") {
		wc = os.Args[2]
		fname = os.Args[3]
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


func ParseInt(str string) uint64	{

	if len(str) == 0 {
		return uint64(0)
	}else {
		val,err := strconv.ParseUint(str, 10, 64)

		if err!=nil	{
			logger.Error("Unable to parse", str, ":", err)
			return val
		}else {
			return val
		}
	}

}

/* math.pow10() fails for uint64 */
func powOfTen(exp string) uint64	{

	var res uint64 = 1

	for i:=1; i<=len(exp); i++ {
		res *= 10
	}
	return res
}

/* Checks for error and prints if found */
func check(err error) bool {
	if err != nil {
		logger.Alert(err)
		return false
	}
	return true
}