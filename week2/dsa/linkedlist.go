package dsa

import "fmt"

//linkedlist (insert, delete, print)
type Node struct {
    Value int
    Next  *Node
}

type LinkedList struct {
    Head *Node
}

func NewNode(value int) *Node {
    return &Node{Value: value, Next: nil}
}

func NewLinkedList(Head *Node) *LinkedList {
    return &LinkedList{Head: Head}
}

func (l *LinkedList) Insert(Value int) {
    node := NewNode(Value)

    if l.Head == nil {
        l.Head = node
    } else {
        last := l.Head

        for last.Next != nil {
            last = last.Next
        }
        last.Next = node
    }
    // TODO: insert at end
}

func (l *LinkedList) Delete(Value int) {
    //if the list is empty list
    if l.Head == nil {
        return
    }

    //if Value is Head of the list, change the Head to the Next (nil)
    if l.Head.Value == Value {
        l.Head = l.Head.Next
        return
    }

    //if not we traverse through the list
    current := l.Head
    for current.Next != nil {
        if current.Next.Value == Value {
            current.Next = current.Next.Next
            return
        }

        current = current.Next
    }
    // TODO: delete first matching Value
}

func (l *LinkedList) Print() {
    // TODO: print all nodes
    current := l.Head
    for current != nil {
        fmt.Printf("-- > %d ", current.Value)
        current = current.Next
    }

    fmt.Println("The end of the list")
}
