func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    Head := &ListNode{}
    end := Head

    for l1 != nil && l2 != nil {
        if l1.Val >= l2.Val {
            end.Next = l2
            l2 = l2.Next
        } else {
            end.Next = l1
            l1 = l1.Next
        }

        end = end.Next
    }

    if l1 != nil {
        end.Next = l1
    }else {
        end.Next = l2
    }

    return Head.Next
}