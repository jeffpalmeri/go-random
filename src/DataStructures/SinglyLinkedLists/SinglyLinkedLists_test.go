package singlylinkedlists

import (
	"testing"
)

func TestAdd(t *testing.T) {
	ssl := NewList(1)
	t.Logf("ddd")

	if ssl.size != 1 {
		t.Errorf("Expected list size to be one, but instead got %v", ssl.size)
	}

	t.Logf("%v", ssl)
}

func TestAppend(t *testing.T) {
	ssl := NewList(1)

	ssl.Append(2);
	
	if(ssl.size != 2) {
		t.Errorf("Expected list size to be two, but instead got %v", ssl.size)
	}

	t.Logf("List length is: %v", ssl.size)
}

func TestLoop(t *testing.T) {
	ssl := NewList(1)

	for i := 2; i <= 5; i++ {
		ssl.Append(i)
	}

	res := ssl.Remove(0)

	if res == false {
		t.Logf("This value aint here: %v", res)
	}

	i := 0
	cur := ssl.head;
	for cur != nil {
		t.Logf("Node at position %v has value %v", i, cur.value)

		cur = cur.next
		i++
	}
}

func TestDifferentLoop(t *testing.T) {
	ssl := NewList(1)

	for i := 2; i <= 5; i++ {
		ssl.Append(i)
	}

	i := 0
	cur := ssl.head;
	for cur.next != nil {
		t.Logf("Node at position %v has value %v", i, cur.value)

		cur = cur.next
		i++
	}
}