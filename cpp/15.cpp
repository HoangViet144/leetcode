class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>>ans;
        set<vector<int>> uniqueAns;

        int n= nums.size();
        
        int ind1,ind2,ind3;
        sort(nums.begin(), nums.end());
        for(ind1=0;ind1<n-2;ind1++){
            if(ind1 > 0 && nums[ind1] == nums[ind1-1]) {
		    	continue;
		    }
            ind2= ind1+1;
            ind3= n-1;
            while (ind2<ind3){
                int sum = nums[ind1]+nums[ind2]+nums[ind3];
                if(sum==0){
                    vector<int> triplet;
                    triplet.push_back(nums[ind1]);
                    triplet.push_back(nums[ind2]);
                    triplet.push_back(nums[ind3]);
                    sort(triplet.begin(), triplet.end());
                    if(uniqueAns.find(triplet)==uniqueAns.end()){
                        uniqueAns.insert(triplet);
                        ans.push_back(triplet);
                    }
                    ind3--;
                } else if(sum>0) {
                    ind3--;
                } else {
                    ind2++;
                }
            }
        }
        
        return ans;
    }
};
