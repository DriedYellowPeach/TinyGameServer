/*
@Time    : 12/5/21 18:12
@Author  : nil
@File    : core_test
*/

package module

import (
	"fmt"
	"testing"
	"time"
)

func TestCore(t *testing.T) {
	c1 := NewCore()

	c1.Server.Register("c1f0", func(i, j string) string {
		return i+j
	})

	c1.Server.Register("c1f1", func(index, t int) int {
		time.Sleep(time.Second * time.Duration(t))
		//fmt.Println("f3 end")
		return index
	})

	c2 := NewCore()
	c2.Server.Register("c2f0", func(i, j string) string {
		return i+j
	})
	c2.Server.Register("c2f1", func(index, t int) int {
		time.Sleep(time.Second * time.Duration(t))
		//fmt.Println("f3 end")
		return index
	})

	go func() {
		c1.Run()
	}()

	go func() {
		c2.Run()
	}()

	c1.Client.Bind(c2.Server)
	c2.Client.Bind(c1.Server)

	index := 0
	for {

		fmt.Println(c1.Client.Call("c2f0", fmt.Sprintf("%d turn:", index), "c1 call c2"))
		fmt.Println(c2.Client.Call("c1f0", fmt.Sprintf("%d turn:", index), "c2 call c1"))
		c1.Client.AsyncCall("c2f1", index, 2, func(index int) {
			fmt.Printf("%d turn: you finally awake, you sleep %d seconds\n", index, 1)
		})
		c2.Client.AsyncCall("c1f1", index, 3, func(index int) {
			fmt.Printf("%d turn you finally awake, you sleep %d seconds\n", index, 1)
		})
		index++
		time.Sleep(1 * time.Second)
	}
}
