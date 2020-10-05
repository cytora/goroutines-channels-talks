package main
//START OMIT
//➜  goroutines-channels-talks git:(master) ✗ go run -race code-session3/failed/racy_infinite_hat.go
//==================
//WARNING: DATA RACE
//Read at 0x00c0000d2000 by goroutine 9:
//main.hatMagician()
///Users/adelinasimion/code/goroutines-channels-talks/code-session3/failed/racy_infinite_hat.go:12 +0x8d
//
//Previous write at 0x00c0000d2000 by goroutine 8:
//main.hatMagician()
///Users/adelinasimion/code/goroutines-channels-talks/code-session3/failed/racy_infinite_hat.go:12 +0x17c
//
//Goroutine 9 (running) created at:
//main.main()
///Users/adelinasimion/code/goroutines-channels-talks/code-session3/failed/racy_infinite_hat.go:31 +0x1fe
//
//Goroutine 8 (running) created at:
//main.main()
///Users/adelinasimion/code/goroutines-channels-talks/code-session3/failed/racy_infinite_hat.go:31 +0x1fe
//==================
//Found 1 data race(s)
//exit status 66
//END OMIT
