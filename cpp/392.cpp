class Solution {
public:
    bool isSubsequence(string s, string t) {
        int n=s.length(), m=t.length();
        if(n>m)return false;
        vector<vector<int>> nxt(m+1,vector<int>(26,m+1));
        for(int i=m-1;i>=0;i--)
        {
            for(int j=0;j<26;j++)nxt[i][j]=nxt[i+1][j];
            nxt[i][t[i]-'a']=i+1;
        }
        int cur=0;
        for(int i=0;i<n;i++)
        {
            if(nxt[cur][s[i]-'a']==m+1)return false;
            cur=nxt[cur][s[i]-'a'];
        }
        return true;
    }
};
