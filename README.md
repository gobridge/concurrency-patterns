## Go Concurrency Patterns

Examples taken from Rob Pike's talk about concurrency patterns.

### Go Concurrency Patterns video:

https://www.youtube.com/watch?v=f6kdp27TYZs

### Notes from Rob's talk

"The composition of independently executing computations"
-- Rob Pike

Concurrency not parallelism.
On a single core, you can't have parallelism.

#### Concurrency
- Easy To Understand
- Easy To Use
- Easy To Reason
- Work at a Higher Level

#### Concurrency is not new
Hoare's CSP Paper in 1978
- Occam ('83), Erlang ('86), Newsqueak ('88), Concurrent ML ('93), Alef ('95), Limbo ('96)

#### Go is a Branch of:
Newsqueak-Alef-Limbo   using Channels

#### Goroutines
- Independently executing function
- Has its own stack which grows and shrinks
- Very Cheap, could have thousands or more
- Not a thread
   - Could have only one thread running thousands of goroutines
- Goroutines are multiplexed dynamically onto threads as needed to keep routines running
- Could think of it as a very cheap thread

#### Buffered Channels
- Channels can be created with a buffer
- Buffering removes synchronization
- Can be important for some problems but they are more subtle to reason about
- Not using them today

#### Summary
- Started With
  - Slow, Sequential and Failure-Sensitive code
- Ended With
  - Fast, Concurrent, Replicated and Robust code

#### Other Patterns

**Chatroulette Toy:**  
http://tinyurl.com/gochatroulette

**Load Balancer:**  
http://tinyurl.com/goloadbalancer

**Concurrent Prime Sieve:**  
http://tinyurl.com/gosieve

**Concurrent Power Series (by Mcllroy)**  
http://tinyurl.com/gopowerseries
