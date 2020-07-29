/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-29 13:45:28
 * @ LastEditTime: 2020-07-29 18:30:56
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
	b := NewBar(0, 150)
	b.Cur <- CurrentStep{
		Cur:     88,
		Message: fmt.Sprintf("step %v", 88),
	}

}
