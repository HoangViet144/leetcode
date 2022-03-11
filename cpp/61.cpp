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
    ListNode* rotateRight(ListNode* head, int k) {
        if(head==NULL)return head;
        int cnt=0;
        ListNode*pCur=head;
        while(pCur)
        {
            cnt++;
            pCur=pCur->next;
        }
        k%=cnt;
        while(k--)
        {
            ListNode*pCur=head;
            int curVal=0,preVal=0;
            while(pCur)
            {
                curVal=pCur->val;
                pCur->val=preVal;
                preVal=curVal;
                pCur=pCur->next;
            }
            head->val=preVal;    
        }
        return head;
    }
};
