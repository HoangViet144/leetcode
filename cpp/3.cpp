class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n=s.length();
        set<char> ss;
        int ans=0,start=0,end=0;
        while(start<n && end<n)
        {
            if(ss.find(s[end])==ss.end())
            {
                ss.insert(s[end++]);
                ans=max(ans,end-start);
            }
            else
                ss.erase(s[start++]);
        }
        return ans;
    }
};
