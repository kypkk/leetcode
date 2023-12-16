package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// dealing with one of the list is empty then just return other
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// dealing with the overflow bit
	over := 0
	if l1.Val+l2.Val >= 10 {
		over = int((l1.Val + l2.Val) / 10)
	}

	/*
	 * Create the answer ListNode:
	 * Val equals to (listone's value + listtwo's value) mod 10
	 * Next null
	 * set l1 to l1 Next, l2 to l2 Next
	 * assign answer to a listnode pointer l
	 */
	answer := ListNode{Val: (l1.Val + l2.Val) % 10, Next: nil}
	l1 = l1.Next
	l2 = l2.Next
	l := &answer

	/*
	 * Dealing with the second to the final value of l1 and l2 with the following for loop
	 */

	for {
		/*
		 * Base case of the loop:
		 * if l1 and l2 are both empty
		 * check if there are overflow bit
		 * if there are an overflow bit create a ListNode with val equals to the overflow bit and Next nil
		 * append the overflowbit node to the l(answer)
		 * loop end
		 */
		if l1 == nil && l2 == nil {
			if over > 0 {
				tmp := ListNode{Val: over, Next: nil}
				l.Next = &tmp
			}
			break
		}

		/*
		 * if one of the list is empty check if there are overflow bit
		 * if no, append the other list to l(answer)
		 * if yes, create a listnode with val: the (other's val + overflow bit) mod 10, Next: nil
		 * append this node to l(answer)
		 * calculate if there are more overflow bit
		 * set l to next and the non empty list to next, and continue
		 */
		if l1 == nil {
			if over == 0 {
				l.Next = l2
				break
			} else {
				tmp := ListNode{Val: (l2.Val + over) % 10, Next: nil}
				l.Next = &tmp

				if l2.Val+over >= 10 {
					over = int((l2.Val + over) / 10)
				} else {
					over = 0
				}
				l = l.Next
				l2 = l2.Next
			}
		} else if l2 == nil {
			if over == 0 {
				l.Next = l1
				break
			} else {
				tmp := ListNode{Val: (l1.Val + over) % 10, Next: nil}
				l.Next = &tmp
				if l1.Val+over >= 10 {
					over = int((l1.Val + over) / 10)
				} else {
					over = 0
				}
				l = l.Next
				l1 = l1.Next
			}
		} else {
			/*
			 * case when both list are non empty
			 * create a listnode val: (l1's value + l2's value + overflow bit) mod 10. Next: nil
			 * append this node to l(answer)
			 * calculate for the next overflow bit
			 * set l to l's next l1, l2 to their next
			 */
			tmp := ListNode{Val: (l1.Val + l2.Val + over) % 10, Next: nil}
			if l1.Val+l2.Val+over >= 10 {
				over = int((l1.Val + l2.Val + over) / 10)
			} else {
				over = 0
			}
			l.Next = &tmp
			l = l.Next
			l1 = l1.Next
			l2 = l2.Next
		}
	}
	return &answer
}

/*
* to verify my answer
 */
func TraverseNode(l1 *ListNode) int {
	var answer int
	count := 0
	for {
		if l1 == nil {
			break
		}
		answer += l1.Val * int(math.Pow(10, float64(count)))
		count += 1
		l1 = l1.Next
	}
	return answer
}

func main() {
	tail := ListNode{Val: 2, Next: nil}
	head := ListNode{Val: 9, Next: &tail}

	tail2 := ListNode{Val: 7, Next: nil}
	head2 := ListNode{Val: 3, Next: &tail2}

	answer := addTwoNumbers(&head, &head2)
	answerVal := TraverseNode(answer)
	fmt.Println(answerVal)
}
