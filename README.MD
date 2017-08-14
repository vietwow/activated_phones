## Task 1: Find unique number and actually activated day
### strategy and algorithm
- Read csv file
- Make a hashmap to store phone number and sorted activation history
  * time complexity O(nlog(n)) (usually, for each number (insert/update) to hashmap O(log(n)))
  * memory complexity O(n)
  * activation history is stored as a sorted list (from high to low)
  * insert history to existed phone number in hashmap get time complexity O(log(n))
- traverse over hashmap
  * get phone number, time complexity O(n)
  * each phone number, traverse over history list(sorted), actual activation is in history record,
    which has activated day != deactivated day of record after it. time complexity O(n)

### source code
Checkout source code in this repo
```
# Help
$ go run main.go -help
# example to run
$ go run main.go -in sample.csv -out result.csv
```

## Task 2: Architecture design
### Design
Look at pdf: [Download](https://github.com/duythinht/activated_phones/raw/master/task2/arch.pdf) | [View](https://github.com/duythinht/activated_phones/blob/master/task2/arch.pdf) (may display too small to view)
### explain:
N = 6000, S = 6, C = 300
- Install 3 server as DB replaca set (for scalable and fault tolerance)
- Install 1 server as Load balancer (nginx or HAProxy)
- We have 2 servers left for application, on each server:
  * Install local cache (redis or memcache)
  * build up applicaton with separate 2 components:
    + View model, which read only from local cache for service product list, product detail and order detail
    + Command model, which make order, update product and fill/update cache when neccessary

- Forward traffic from internet from load balancer to 2 instance

Note: This design assume for each application instance, local cache and view model can serve up to 3000 req/s (N/2) and command model can handle up to 150 req/s (C/2), read always more often than write. In my opinion this is entirely feasible
