## Task 1: Find unique number and actually activated day
### strategy and algorithm
- Read csv file
- Make a hashmap to store phone number and sorted activation history
  * time complexity O(nlog(n)) (usually, for each number to hashmap O(log(n)))
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
$ go run main -help
# example to run
$ go run main -in sample.csv -out result.csv
```

## Task 2: Architecture design
### chart
