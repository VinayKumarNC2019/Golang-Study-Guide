# test

Here's an example of a slice of structs:

```go
package main

import (
	"fmt"
)

func main() {
	type Employee struct {
		ID        int
		FirstName string
		Lastname  string
	}

	/*
		Here we are defining a variable called myEmployees which is slice of structs.
	*/
	var myEmployees []Employee
	fmt.Println(myEmployees)  // prints []

	myEmployees = append(myEmployees,
		Employee{
			ID:        007,
			FirstName: "James",
			Lastname:  "Bond",
		},
    )
    
    myEmployees = append(myEmployees,
		Employee{
			ID:        004,
			FirstName: "Stuart",
			Lastname:  "Thomas",
		},
	)

	fmt.Println(myEmployees) // prints [{7 James Bond} {4 Stuart Thomas}]
	fmt.Println(myEmployees[0]) // {7 James Bond}

}

```

Running this in the [golang playground](https://play.golang.org/) gives:

```
[]
[{7 James Bond}]
{7 James Bond}
```


Here's how to make it more memory efficient using pointers. 

```golang
package main

import (
	"fmt"
)

func main() {
	type Employee struct {
		ID        int
		FirstName string
		Lastname  string
	}

	/*
		Here we are defining a variable called myEmployees which is slice of structs.
	*/
	var myEmployees []*Employee
	fmt.Println(myEmployees) // prints []

	/*
	   Now create a couple struct to be added to our slice
	*/

	var james Employee // just initialised but not declared yet
	// another struct, 'jack' which we initialised plus declared
	var jack Employee
	jack = Employee{
		ID:        23,
		FirstName: "Jack",
		Lastname:  "Bauer",
	}

    // now we add the objects to our array. 
	myEmployees = append(myEmployees, &james)
	myEmployees = append(myEmployees, &jack)

	fmt.Println(myEmployees)     // prints [0x446280 0x4462a0] // i.e. an array of memory pointers
	fmt.Println(*myEmployees[0]) // prints {0  }
	fmt.Println(*myEmployees[1]) // {23 Jack Baur}
}
```

This outputs:

```
[]
[0x43e280 0x43e2a0]
{0  }
{23 Jack Bauer}
```

Here's another example: [https://github.com/Sher-Chowdhury/gsg_slices_pointers_structs_demo](https://github.com/Sher-Chowdhury/gsg_slices_pointers_structs_demo)