/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* l1, ListNode* l2) {
        ListNode *head=new ListNode(0);
        ListNode *pCur=head;
        while(l1 && l2)
        {
            if(l1->val <l2->val)
            {
                pCur->next=new ListNode(l1->val);
                l1=l1->next;
            }
            else
            {
                pCur->next=new ListNode(l2->val);
                l2=l2->next;
            }
            pCur=pCur->next;
        }
        while(l1)
        {
            pCur->next=new ListNode(l1->val);
            l1=l1->next;
            pCur=pCur->next;
        }
        while(l2)
        {
            pCur->next=new ListNode(l2->val);
            l2=l2->next;
            pCur=pCur->next;
        }
        return head->next;
    }
};
