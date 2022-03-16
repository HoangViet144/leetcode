class Solution {
public:
    int romanToInt(string s) {
        map<char, int> char2val;
        
        char2val.insert({ 'I', 1 });
        char2val.insert({ 'V', 5 });
        char2val.insert({ 'X', 10 });
        char2val.insert({ 'L', 50 });
        char2val.insert({ 'C', 100 });
        char2val.insert({ 'D', 500 });
        char2val.insert({ 'M', 1000 });
        
        int num = 0;
        for (int i = 0; i < s.length(); i++){
            if (i+1 <s.length() && char2val[s[i]] < char2val[s[i + 1]]) {
                num += char2val[s[i+1]] - char2val[s[i]];
                i++;
            } else {
                num += char2val[s[i]];            
            }
        }

        return num;
    }
};
