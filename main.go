package main

import (
	"fmt"

	"github.com/recor-glitch/LinkList/singleLinkList"
)

func main() {

	data := []any{
		map[string]string{"A": "value1", "B": "value2"},
		"value1",
		1,
		true,
		[]int{1, 2, 3, 4, 5},
		"value2",
	}

	// SINGLY LINK LIST
	ll := singleLinkList.LinkList{}

	// INSERT DATA INTO LINK LIST
	for i, d := range data {
		_, err := ll.Insert(fmt.Sprintf("key%v", i), d)
		if err != nil {
			fmt.Println(err)
		}
	}

	// DISPLAY DATA FROM LINK LIST
	ll.Display()
}
