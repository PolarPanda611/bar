## bar 

Usage 
```
package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/PolarPanda611/bar"
)

func main() {
	b := bar.NewBar(0, 150)
	for i := int64(0); i <= 150; i++ {
		time.Sleep(100 * time.Millisecond)
		if i == 88 {
			b.Cur <- bar.CurrentStep{
				Cur: i,
				Err: errors.New("wocao"),
			}
			break
		} else {
			b.Cur <- bar.CurrentStep{
				Cur:     i,
				Message: fmt.Sprintf("step %v", i),
			}
		}
	}
}

```

Result 
```
# go run main.go
[█████████████████████████████████████████████████ ] 99%       149/150 step 149

```