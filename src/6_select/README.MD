### Select:
* Select statement blocks until one of the send/receive operation is ready. Takes any random one if both of them ready
* Select statement mainly used with the channels when you dont know which one is ready either send/receive
* Select statement default case is used to avoid panics if none of them ready and avoid blocking

