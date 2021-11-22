/*
@Time    : 11/21/21 21:14
@Author  : nil
@File    : example_test.go
*/

package chaninvoke_test

import (
	"chaninvoke"
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	s := chaninvoke.StartServer(10)

	done := make(chan struct{})

	go func(){
		s.Register("f0", func(args ...interface{}) interface{}{
			sum := 0
			for _, v := range args {
				sum += v.(int)
			}
			return interface{}(sum)
		})

		done <- struct{}{}

		for {
			s.Exec(<- s.CallChan)
		}
	}()
	fmt.Println("register done")
	<-done

	go func() {
		c := chaninvoke.StartClient(10, s)
		fmt.Println(c.Call("f0", 1, 2, 3, 4))
		done <- struct{}{}
	}()

	<-done
}

