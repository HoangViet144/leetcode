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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int digit=0;
        int remainder=0;
        ListNode *ans=new ListNode(0);
        ListNode *cur=ans;
        while(l1 || l2)
        {
            int x= (l1?l1->val :0);
            int y= (l2?l2->val :0);
            digit= (x+y+remainder)%10;
            remainder=(x+y+remainder)/10;
            cur->next=new ListNode(digit);
            cur=cur->next;
            if(l1)l1=l1->next;
            if(l2)l2=l2->next;
        }
        if(remainder)cur->next=new ListNode(remainder);
        return ans->next;
    }
};
