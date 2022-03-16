class Solution {
public:
    bool validateStackSequences(vector<int>& pushed, vector<int>& popped) {
        stack<int> st;
        int indPu=0,indPo=0;
        while (indPu < pushed.size() && indPo < popped.size()){
            if(!st.empty() && st.top()==popped[indPo]){
                st.pop();
                indPo++;
            } else {
                st.push(pushed[indPu++]);                
            }
        }
        while(indPo < popped.size()){
            if(!st.empty() && st.top()==popped[indPo]){
                st.pop();
                indPo++;
            } else {
                return false;
            }
        }
        return true;
    }
};
