class Solution {
public:
    string removeDuplicateLetters(string s) {
        vector<int> letter(26,0);
        for(char i:s){
            letter[i-'a']++;
        }
        vector<bool>visited(26,0);
        string ans="";
        for(char i: s){
            letter[i-'a']--;
            if(!visited[i-'a']){
                while(ans.size() && ans.back()>i && letter[ans.back()-'a']){
                    visited[ans.back()-'a']=false;
                    ans.pop_back();
                }
                visited[i-'a']=true;
                ans+=i;
            }
        }
        return ans;
    }
};
