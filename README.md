# Multimap

A Go package implementing a generic multimap data structure, allowing multiple values per key with support for immutable, mutable, and sequential access.

## Features

- **Immutable Multimap**: Key-value mappings with read-only operations.
- **Mutable Multimap**: Add, remove, and modify key-value pairs.
- **Sequential Multimap**: Iterate over values for a key sequentially.

## Installation

```bash
go get github.com/nkamenev/multimap
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/nkamenev/multimap"
)

func main() {
	// Create a mutable multimap
	mutRecipes := Make[string, string](1)
	mutRecipes.Set("pancakes", "flour", "milk", "eggs")

	// Freeze multimap
	recipes := mutRecipes.Immutable()

	// Get values
	ingredients, _ := recipes.Get("pancakes")
	fmt.Println(ingredients) // Output: [flour milk eggs]

	// Sequential access
	s := recipes.Sequential()
	fmt.Println(s.Next("pancakes")) // Output: flour true
	fmt.Println(s.Next("pancakes")) // Output: milk true
	fmt.Println(s.Next("pancakes")) // Output: eggs true
	fmt.Println(s.Next("pancakes")) // Output:  false
}
```

## Testing

Run tests with:

```bash
go test -v
```

## License

# MIT
