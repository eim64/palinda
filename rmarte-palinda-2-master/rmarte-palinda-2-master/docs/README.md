# Task 1 
## [bug01.go](../src/bug01.go)

Unbuffered channels will not write unless it's currently being read. 

Fixed by making the channel buffered.

This means a few values can be stored before it requires a channel read. 



## [bug02.go](../src/bug02.go)
The main process is closed before the subroutine can print a value.

Fixed by adding "I'm finished you can go die now" channel that's piped to whenever the subroutine is finished. 

Since main process halts until the "I'm finished you can go die now", it will not go die before the subroutine is finished.

# Task 2
* What happens if you switch the order of the statements
  `wgp.Wait()` and `close(ch)` in the end of the `main` function?

The channel will close just as all subroutines have been created, and the producers will try to write to a closed channel and then panic.

* What happens if you move the `close(ch)` from the `main` function
  and instead close the channel in the end of the function
  `Produce`?

The channel will close when the first produce subroutine is finished and the remaining ones will try to write to a closed channel and then panic. 

* What happens if you remove the statement `close(ch)` completely?

each for in the Consume function should run until the channel is closed. If there was a wait for each Consume to finish, it would halt before priting the time differential and never finish running.

* What happens if you increase the number of consumers from 2 to 4?
The rate at which producers and consumers would coincide and it should get about twice as fast, as long as there's enough cpu threads. The expected speed should be around strings / min(producers, consumers, threads), since produce and consume are just as fast. 


* Can you be sure that all strings are printed before the program
  stops?

As mentioned earlier, close only makes sure the loops in Consume can finish. Not that they will finish. Thus the main thread should be finished and kill the Consume routines before they're finished.  