# Enum

Enum provides a simple enum implementation for Golang.

# Usage

First of all you need to import the package as

```go
import "github.com/piteego/enum"
```
Then install it using go get as following:
```bash
go get "github.com/piteego/enum"
```

# Example

Let's say you want to create an enum for the traffic light colors.
Here is how we usually do it.

```go
package traffic

type Light int8

const (
	Red Light = iota // 0
	Yellow           // 1
	Green            // 2
)
```
By implementing the identifier interface as following:

```go
func(Light) EnumUid() string {return "traffic.Light"}
```
You can register your enum as following:

```go
package traffic

import "github.com/piteego/enum"

func init() {
    enum.Register(map[Light]string{
        Red:    "Red",
        Yellow: "Yellow",
        Green:  "Green",
    })
}
```

And enjoy some fancy functions to interact with your enum values:

- ***Is***: checks if the given value is equal to your target enum values.

```go
    enum.Is(traffic.Light(0), traffic.Red)                   // Output: true
    enum.Is(traffic.Light(1), traffic.Red, traffic.Green)    // Output: false
```

- ***Validate***: checks if the given value is a valid enum value.

```go
    enum.Validate(traffic.Light(0))  // Output: nil
	enum.Validate[traffic.Light](0) // Output: nil
    enum.Validate(traffic.Red)  // Output: nil
	err := enum.Validate(traffic.Light(3))  // Output: [Enum] invalid enum value for traffic.Light: must be one of [0,1,2], got 3 
	errors.Is(err, enum.ErrInvalidValue) // Output: true
```

- ***New***: creates a new enum value from the given string.

```go
    enum.New[traffic.Light]("Red")  // Output: enum=pointer to traffic.Red, err= nil 
	red, err := enum.New[traffic.Light]("red") 
	// Output: nil, [Enum] invalid enum value for traffic.Light: must be one of [Red,Yellow,Green], got red
	errors.Is(err, enum.ErrInvalidValue) // Output: true
```

- ***String***: returns the string representation of the given enum value.

```go
    enum.String(traffic.Red)  // Output: "Red"
	enum.String(traffic.Light(3))  // Output: ""
```

- ***Strings***: returns slice of string for the given enum type.

```go
    enum.Strings[traffic.Light]() // Output: []string{"Red", "Yellow", "Green"}
```

- ***Values***: returns slice of enum values.

```go
    enum.Values[traffic.Light]() // Output: []traffic.Light{0, 1, 2}
```