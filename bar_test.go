/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-29 13:45:28
 * @ LastEditTime: 2020-07-29 18:47:56
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /bar/bar_test.go
 * @
 */
package bar

import (
	"fmt"
	"testing"
)

func TestBar(t *testing.T) {
	b := NewBar(0, 100)
	b.Cur <- CurrentStep{
		Cur:      50,
		Message:  fmt.Sprintf("step %v", 50),
		IsReturn: true,
	}
	b.Cur <- CurrentStep{
		Cur:     86,
		Message: fmt.Sprintf("step %v", 86),
	}
	b.Cur <- CurrentStep{
		Cur:     87,
		Message: fmt.Sprintf("step %v", 87),
	}
	b.Cur <- CurrentStep{
		Cur:     100,
		Message: fmt.Sprintf("step %v", 100),
	}

}
