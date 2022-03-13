class Solution {
public:
    bool isValid(string s) {
        stack<char> st;
        for(char i: s){
            if(i=='('|| i == '[' || i == '{') st.push(i);
            else {
                if(st.empty())return false;
                char cur = st.top();
                st.pop();
                if(i==')' && cur != '(') return false;
                if(i==']' && cur != '[') return false;
                if(i=='}' && cur != '{') return false;
            }
        }
        return st.size() == 0;
    }
};
