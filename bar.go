/**
 * @ Author: Daniel Tan
 * @ Date: 2020-07-29 13:29:54
 * @ LastEditTime: 2020-07-29 14:02:11
 * @ LastEditors: Daniel Tan
 * @ Description:
 * @ FilePath: /bar/bar.go
 * @
 */

package bar

import (
	"fmt"
)

var (
	_defaultGraph = "█"
)

// CurrentStep current step information
type CurrentStep struct {
	Cur     int64
	Message string
	Err     error
}

// Bar bar instance
type Bar struct {
	percent int64  //百分比
	cur     int64  //当前进度位置
	total   int64  //总进度
	rate    string //进度条
	graph   string //显示符号

	Cur  chan CurrentStep
	Done chan error
}

// NewOption init start and total
func (bar *Bar) NewOption(start, total int64) {
	bar.cur = start
	bar.total = total
	if bar.graph == "" {
		bar.graph = _defaultGraph
	}
	bar.percent = bar.getPercent()
	for i := 0; i < int(bar.percent); i += 2 {
		bar.rate += bar.graph //初始化进度条位置
	}
}

func (bar *Bar) getPercent() int64 {
	return int64(float32(bar.cur) / float32(bar.total) * 100)
}

// Play play bar
func (bar *Bar) Play(cur int64, message string) {
	bar.cur = cur
	last := bar.percent
	bar.percent = bar.getPercent()
	if bar.percent != last && bar.percent%2 == 0 {
		bar.rate += bar.graph
	}
	fmt.Printf("\r[%-50s]%3d%%  %8d/%d %v", bar.rate, bar.percent, bar.cur, bar.total, message)
}

// RunBar run bar instance
func (bar *Bar) RunBar() {
Loop:
	for {
		select {
		case i := <-bar.Cur:
			if i.Err != nil {
				bar.Play(i.Cur, i.Err.Error())
				bar.Done <- i.Err
			}
			bar.Play(i.Cur, i.Message)
			if i.Cur == bar.total {
				bar.Done <- nil
				break Loop
			}
		case err := <-bar.Done:
			bar.Finish(err)
			break Loop
		}
	}
}

// Finish finish bar
func (bar *Bar) Finish(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println()
}

// NewBar new bar with default
func NewBar(start, total int64) *Bar {
	var b Bar
	b.NewOption(start, total)
	b.Cur = make(chan CurrentStep)
	return &b
}

// NewBarWithGraph new bar with customize graph
func NewBarWithGraph(start, total int64, graph string) *Bar {
	b := NewBar(start, total)
	b.graph = graph
	return b
}
