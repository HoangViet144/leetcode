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
    ListNode* deleteDuplicates(ListNode* head) {
        if(head==NULL)return head;
        ListNode*newHead=new ListNode(-1);
        ListNode* cur=head;
        newHead->next=head;
        ListNode* pre=newHead;
        while(cur)
        {
            while(cur->next  && pre->next->val==cur->next->val)
                cur=cur->next;
            if(pre->next==cur)
                pre=pre->next;
            else
                pre->next=cur->next;
            cur=cur->next;
        }
        return newHead->next;
    }
};
