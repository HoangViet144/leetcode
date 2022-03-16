class RecentCounter {
    queue<int> reqLst;
public:
    RecentCounter() {
        
    }
    
    int ping(int t) {
        while(!reqLst.empty() && reqLst.front() < t- 3000)reqLst.pop();
        reqLst.push(t);
        return reqLst.size();
    }
};

/**
 * Your RecentCounter object will be instantiated and called as such:
 * RecentCounter* obj = new RecentCounter();
 * int param_1 = obj->ping(t);
 */
