/*Package lists contains a series of data structures based on lists.

Stack - O(1) FILO based on linked lists, any values interface{}
Queue - O(1) FIFO based on linked lists, any values interface{}

Scenario 1:
Faster, but not safe for concurrency.
var listNotSafe := lists.NewStack(false) //Stack,Queue

Scenario 2:
If you use goroutines create one using
var listSafe := lists.NewStack(true) //Stack,Queue
Most common error is "stack was empty", check Stack.Empty() or ignore it in highly-concurrent funcs.
Because the state may change between the HasElement() call and Pop/Peek.

Scenario 3:
Manual lock the struct, 100% reability, prune to mistakes/bugs
var listNotSafe := lists.NewStack(false) //Stack,Queue
listNotSafe.Lock()
//do stuff with the list
listNotSafe.Unlock()

For more details see the README and *_test.go*/
package lists
