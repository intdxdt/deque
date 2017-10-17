# deque
double ended queue for go
implemented on top of a array base to support indexing.

# motivation 
* allocates sufficient space if necessary for:`append`, `appendLeft`,  `popleft` , `popright`
* cache friendly
* indexing (`deck[i]`)

