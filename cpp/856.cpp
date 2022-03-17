class Solution {
public:
    int scoreOfParentheses(string s) {
        stack<int> st;
        st.push(0);
        for(char i: s){
            if(i=='(') st.push(0);
            else {
                int cur= st.top();
                st.pop();
                int cur2= st.top();
                st.pop();
                st.push(cur2+max(2*cur,1));
            }
        }
        return st.top();
    }
};
