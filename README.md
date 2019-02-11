# linked-list

## singly linked-list
Singly linked-list (stack) works in LIFO order.
##### Usage
```
sl := singly.List{}
```
### Methods
#### Read()
This method will give back a slice of interface with the data in the list in LIFO order.
##### Usage
```
data := sl.Read()
```
#### Push()
Push method inserts data into the list.
##### Usage
```
sl.Push(5) 
// sl.Read() = [5]

or

sl.Push(5, 6, 7, 8)
// sl.Read() = [8 7 6 5]
```
#### Pop()
This method will pops out the recently pushed data from the list. The popped data will be deleted from the list and data before that will be the new head.
##### Usage
```
// sl.Read() = [8 7 6 5]

poppedData := sl.Pop()
// poppedData = 8
// sl.Read() = [7 6 5]
```
#### Reverse()
Revese will reverse the list.
##### Usage
```
// sl.Read() = [8 7 6 5]

sl.Reverse()
// sl.Read() = [5 6 7 8]
```

## doubly linked-list
Doubly linked-list can work in both LIFO and FIFO (queue) order.
##### Usage
```
dl := doubly.List{}
```
#### ReadLIFO()
This method will give back a slice of interface with the data in the list in LIFO order.
##### Usage
```
data := dl.ReadLIFO()
```
#### ReadFIFO()
This method will give back a slice of interface with the data in the list in FIFO order.
##### Usage
```
data := dl.ReadFIFO()
```
#### Push()
Push method inserts data into the list.
##### Usage
```
dl.Push(5) 
// dl.Read() = [5]

or

dl.Push(5, 6, 7, 8)
// dl.ReadLIFO() = [8 7 6 5]
```
#### PopHead()
This method will pops out the recently pushed data from the list. The popped data will be deleted from the list and data before that will be the new head.
##### Usage
```
// dl.ReadLIFO() = [8 7 6 5]

poppedData := dl.PopHead()
// poppedData = 8
// dl.ReadLIFO() = [7 6 5]
```
#### PopTail()
This method will pops out the initially pushed data from the list. The popped data will be deleted from the list and data after that will be the new tail.
##### Usage
```
// dl.ReadLIFO() = [8 7 6 5]

poppedData := dl.PopTail()
// poppedData = 5
// dl.ReadLIFO() = [8 7 6]
```
