class Solution {
public:
    int majorityElement(vector<int>& nums) {
        map<int,int>freq;
        for(int i=0;i<nums.size();i++)
            if(freq.count(nums[i])==0)freq[nums[i]]=1;
            else freq[nums[i]]++;
        int ans=-1, f=0;
        for(auto it=freq.begin();it!=freq.end();it++)
        {
            if(f<it->second)
            {
                ans=it->first;
                f=it->second;
            }
        }
        return ans;
    }
};
