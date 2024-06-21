# Go-Essentials

Go Program Key Components

## Files and Packages

- Go Code splits into multiple packages
- we can export and import packages

### Importance of Main

- main entry point to run an application
- Go Modules and Buidling Go Programs
- one module consists of multiple packages
- ```go mod init``` : To initialize a module
- ```go build``` : to make the module runnable

### Onwards to a new Project

- Variables and Types

  - int ,  float64 , boolean , string
  - shortand operator(:=)
  - Constant Values ( fixed value )
  - Importance of variables
  
### Go Pointers

- What are Pointers:

  - Pointers are variables that store value address instead of values
  - code ----> Computer memory
  - Variables ---> Value (Address)
  - * represents the dereferencing the variable
  - & represents the address of a variable
  - 
- Why Pointers:

### Structs
- Grouping Data and Functions into collections

### Interfaces :
- Interfaces are like combining repeated methods into a single type
- Mentioning the names of an methods to an inteface type
- Repeated methods 
type <inteface-name> interface {

}
- Embedded Intefaces : interfaces inside the interface

### Interfaces Dynamic Types and Limitations:

### Generics :  
- func [Placholder name]() function-name <genaric> {}

### Arrays
- collections of similar kind of data is called arrrays 
- where we can store multiple values

### Slices
- Cutting the arrays is called arrays
- len : No of elements and cap will gives capacity
- 
