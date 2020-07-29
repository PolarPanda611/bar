/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-29 13:45:28
 * @ LastEditTime: 2020-07-29 14:00:14
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
	go b.RunBar()
	for i := int64(0); i <= 150; i++ {
		b.Cur <- CurrentStep{
			Cur:     i,
			Message: fmt.Sprintf("step %v", i),
		}
	}

}
