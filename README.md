# Golang Generics Project

This is a simple implementation of a linked list and queue with generic types in Go. The linked list is defined in the `list` package, and the main application in the `main` package demonstrates the use of the linked list with different data types.

## List Package

### `list.go`

The `list` package provides a generic linked list implementation. It includes the following types:

#### `node[T any]`

- Represents a node in the linked list.
- Contains a value of type `T` and references to the next and previous nodes.

#### `linkedList[T any]`

- Represents the linked list itself.
- Includes methods for creating a new linked list, pushing elements to the front and back, and accessing the head and tail nodes.

#### Methods:

- `New[T any]() *linkedList[T]`: Creates a new linked list.
- `PushFront(value T)`: Adds a new node with the specified value to the front of the list.
- `PushBack(value T)`: Adds a new node with the specified value to the back of the list.
- `Head() *node[T]`: Returns the head node of the list.
- `Tail() *node[T]`: Returns the tail node of the list.

## Main Package

### `Main.go`

The `main` package demonstrates the usage of the generic linked list with different data types.

#### Types:

- `User`: Represents a simple user structure with name, email, and password fields.

#### Main Function:

- Creates a linked list for integers and adds values to both the front and back.
- Creates a linked list for `User` structures and adds user objects to the list.
- Prints the values of the linked lists.

## Usage

To incorporate the generic linked list into your Go project, follow these steps:

1. Import the `github.com/mxnuchim/golang-generics/list` package.

   ```go
   import (
       "fmt"
       "github.com/mxnuchim/golang-generics/list"
   )
   ```

2. Create a new linked list for your desired data type. For example, to create a linked list for integers:

   ```go
   intList := list.New[int]()
   ```

3. Add elements to the linked list using `PushFront` or `PushBack` methods:

   ```go
   intList.PushFront(42)
   intList.PushBack(99)
   ```

4. Access the head and tail nodes as needed:

   ```go
   headNode := intList.Head()
   tailNode := intList.Tail()
   ```

5. Iterate through the linked list to retrieve values:

   ```go
   node := intList.Head()
   for node != nil {
       fmt.Println("Value:", node.Value())
       node = node.Next()
   }
   ```

6. Customize the linked list for other data types. For example, to create a linked list for `User` structures:

   ```go
   users := list.New[User]()
   ```

7. Add user objects to the list and iterate through the linked list:

   ```go
   users.PushBack(User{name: "John Doe", email: "john@example.com", password: "secure123"})
   ```

   ```go
   userNode := users.Head()
   for userNode != nil {
       fmt.Println("User Name:", userNode.Value().name)
       userNode = userNode.Next()
   }
   ```

Feel free to adapt and integrate this generic linked list implementation into your Go projects to efficiently manage dynamic collections of various data types.
# golang-generics
# golang-generics
# golang-generics
