class Solution {
public:
    vector<int> countBits(int num) {
        if (num==0)return vector<int>(1,0);
        if (num == 1)
        {
            vector<int>res;
            res.push_back(0);
            res.push_back(1);
            return res;
        }
        vector<int>res (num+1,0);
        res[0]=0;
        int cur=1;
        while(cur*2<=num)
        {
            res[cur]=1;
            for(int i=cur+1; i<cur*2;i++) res[i]=res[i-cur]+1;
            cur*=2;
        }
        for(int i=cur; i<=num;i++) res[i]=res[i-cur]+1;
        return res;
    }
};