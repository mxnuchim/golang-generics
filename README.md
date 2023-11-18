# Golang Generics Linked List and Queue

This project provides generic implementations of a linked list and a queue in Go. The `list` package contains a generic linked list, and the `queue` package contains a generic queue. The `main` package demonstrates the usage of these data structures with different data types.

## Linked List Package

### `List.go`

The `list` package provides a generic linked list implementation. It includes the following types:

- `node[T any]`: Represents a node in the linked list.
- `linkedList[T any]`: Represents the linked list itself.

#### Methods:

1. **`New[T any]() *linkedList[T]`**: Creates a new linked list.
2. **`PushFront(value T)`**: Adds a new node with the specified value to the front of the list.
3. **`PushBack(value T)`**: Adds a new node with the specified value to the back of the list.
4. **`Head() *node[T]`**: Returns the head node of the list.
5. **`Tail() *node[T]`**: Returns the tail node of the list.

#### Usage:

````go
// 1. Create a linked list for integers
intList := list.New[int]()
intList.PushFront(90)
intList.PushFront(243)
intList.PushFront(876) // add value to start (head) of linked list
intList.PushBack(1034) // add value to end (tail) of linked list

// 2. Iterate through the linked list
node := intList.Head()
for node != nil {
    fmt.Println("value --> ", node.Value())
    node = node.Next()
}

// 3. Real world use case with users
users := list.New[User]()
users.PushBack(User{name: "Angela Simmons", email: "angela@example.com", password: "xxxx"})
users.PushBack(User{name: "Manuchim Oliver", email: "manuchim@example.com", password: "xxxx"})
users.PushBack(User{name: "Posha Alabi", email: "posha@example.com", password: "xxxx"})

// 4. Iterate through the linked list of users
userNode := users.Head()
for userNode != nil {
    fmt.Println("User name --> ", userNode.Value().name)
    userNode = userNode.Next()
}

## Queue Package

### `queue.go`

The `queue` package provides a generic queue implementation. It includes the following type:

- `q[T any]`: Represents the queue.

#### Methods:

1. **`New[T any]() *q[T]`**: Creates a new queue.
2. **`Enqueue(item T) *q[T]`**: Adds an item to the queue.
3. **`Dequeue() (T, error)`**: Removes and returns an item from the front of the queue.
4. **`String() string`**: Converts the queue to a string for printing.

#### Usage:

```go
// 1. Create a queue for integers
q := queue.New[int]()
for i := 1; i <= 11; i++ {
    q.Enqueue(i * 2)
}

// 2. Print the contents of the queue
fmt.Println(q.String())

// 3. Realistic use case with users in a queue
userQueue := queue.New[User]()
userQueue.Enqueue(User{name: "Posha Alabi", email: "posha@example.com", password: "xxxx"}).
    Enqueue(User{name: "Manuchim Oliver", email: "manuchim@example.com", password: "xxxx"}).
    Enqueue(User{name: "Angela Simmons", email: "angela@example.com", password: "xxxx"})

// 4. Print the contents of the user queue
fmt.Println(userQueue.String())


````

## Main Package

### `Main.go`

The `main` package demonstrates the usage of the generic linked list and queue with different data types.

#### Types:

- `User`: Represents a simple user structure with name, email, and password fields.

#### Usage:

1. **Linked List Example:**

   ```go
   fmt.Println("\n\n <-- Linked List -->")
   // Golang generics example with linked list for any data type
   linkedList := list.New[int]()
   linkedList.PushFront(90)
   linkedList.PushFront(243)
   linkedList.PushFront(876) // add value to start (head) of linked list
   linkedList.PushBack(1034) // add value to end (tail) of linked list

   // Iterate through the linked list
   node := linkedList.Head()
   for node != nil {
       fmt.Println("value --> ", node.Value())
       node = node.Next()
   }

   // Real-world use case with users
   users := list.New[User]()
   users.PushBack(User{name: "Angela Simmons", email: "angela@example.com", password: "xxxx"})
   users.PushBack(User{name: "Manuchim Oliver", email: "manuchim@example.com", password: "xxxx"})
   users.PushBack(User{name: "Posha Alabi", email: "posha@example.com", password: "xxxx"})

   // Iterate through the linked list of users
   userNode := users.Head()
   for userNode != nil {
       fmt.Println("User name --> ", userNode.Value().name)
       userNode = userNode.Next()
   }
   ```

2. **Queue Example:**

   ```go
   fmt.Println("\n\n <-- Queue -->")

   // In this case, the queue will be expanded as the number of items to be added exceeds the initial length of items array (10)
   q := queue.New[int]()
   for i := 1; i <= 11; i++ {
       q.Enqueue(i * 2)
   }

   // Print the contents of the queue
   fmt.Println(q)

   // Realistic use case with users in a queue
   userQueue := queue.New[User]()
   userQueue.Enqueue(User{name: "Posha Alabi", email: "posha@example.com", password: "xxxx"}).
       Enqueue(User{name: "Manuchim Oliver", email: "manuchim@example.com", password: "xxxx"}).
       Enqueue(User{name: "Angela Simmons", email: "angela@example.com", password: "xxxx"})

   // Print the contents of the user queue
   fmt.Println(userQueue)
   ```

Feel free to adapt and integrate this code snippet into your project's documentation.
