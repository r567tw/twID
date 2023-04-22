# twID

The `twID` package provides a function for verifying Taiwan national identification numbers. This is useful for validating the format and checksum of the ID number.

## Installation

You can install the `twID` package using `go get`:
```bash
go get github.com/r567tw/twID
```

## Usage

Here's an example of how to use the `twID` package:

```go
package main

import (
	"fmt"
	"github.com/r567tw/twID"
)

func main() {
	id := "A123456789"

	if twID.Verify(id) {
		fmt.Printf("%s is a valid Taiwan national identification number.\n", id)
	} else {
		fmt.Printf("%s is not a valid Taiwan national identification number.\n", id)
	}
}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)
