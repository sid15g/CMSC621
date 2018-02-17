# [CMSC-621 Advanced Operating Systems](https://www.csee.umbc.edu/~kalpakis/Courses/621-sp18/cmsc621.php)

Implemented the academic project(s), under [Prof. K. Kalpakis](https://www.csee.umbc.edu/~kalpakis/) in `CMSC-621`


### [HW1](https://www.csee.umbc.edu/~kalpakis/Courses/621-sp18/homeworks/hw1c.php):
Write a GoLang multithreaded applications to compute the sum of integers stored in a file.

#### Prerequisites:
* GoLang
* GO library [stdlog](https://github.com/alexcesaro/log)
* Linux Environment

#### Usage:
* Install `stdlog` library using `make configure`
* `make build`
* `./main -log=<loglevel> <#workers> <filename>`
* Specifying log level is optional, default log level is 'info'

