class Solution {
public:
    string minRemoveToMakeValid(string s) {
        stack<int> st;
        vector<int> valid, invalid;
    
        for(int i=0;i<s.length();i++){
            if(s[i]=='('){
                st.push(i);
            } else if(s[i]==')'){
                if(st.empty()) invalid.push_back(i);
                else {
                    valid.push_back(st.top());
                    valid.push_back(i);
                    st.pop();
                }
            }
        }
        sort(valid.begin(),valid.end());
        int ind=0,indIn=0;
        int szValid = valid.size(), szInvalid = invalid.size();
        string ans="";
        for(int i=0;i<s.length();i++){
            if(s[i]!='(' && s[i]!=')')ans+=s[i];
            else {
                if(indIn<szInvalid && invalid[indIn]==i){
                    indIn++;
                    continue;
                }
                if(ind<szValid && valid[ind]==i){
                    ans+=s[i];
                    ind++;
                }
            }
        }
        return ans;
    }
};
