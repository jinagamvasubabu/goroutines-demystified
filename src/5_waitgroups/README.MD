### Waitgroups
* Waitgroups is used for waiting for a collection of goroutines to completes their tasks 
* It provides a goroutine synchronization mechanism

```go
wg.Add(i) - Adds i number of gorutines to the wait group 
wg.wait() - wait for all goroutines to finish
```

