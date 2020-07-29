/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-29 13:45:28
 * @ LastEditTime: 2020-07-29 15:04:57
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /bar/bar_test.go
 * @
 */
package bar

import (
	"errors"
	"fmt"
	"testing"
)

func TestBar(t *testing.T) {
	b := NewBar(0, 150)
	go b.RunBar()
	for i := int64(0); i <= 150; i++ {
		if i == 88 {
			b.Cur <- CurrentStep{
				Cur: i,
				Err: errors.New("err"),
			}
		}
		b.Cur <- CurrentStep{
			Cur:     i,
			Message: fmt.Sprintf("step %v", i),
		}
	}

}
