
## Concurrency is not parallelism
##### **Parallelism**:
* Doing Multiple things at once
* Number of tasks you are doing in parallel is directly proportional to the number of cores in your system
* Order of execution is not predictable and we should not rely on them to execute in any specific order

Eg:
* Dreaming and sleeping at the same time
* Drinking and Dancing at the same time


##### Concurrency:
* Concurrency is nothing but designing of your code in to an independently executable function and should not have any side effects even if you  run 100,000 times
* Running independently executable function separately and can communicate in between to synchronize their tasks

Eg: Multitasking on single core machine
 