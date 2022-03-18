class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        left = 0
        right = 0
        
        ans = 0
        
        letter = [0]*256
        while right < len(s):
            c = s[right]
            letter[ord(c)]+=1
            
            while letter[ord(c)] > 1:
                t = s[left]
                letter[ord(t)] -=1
                left+=1
            
            ans = max(ans, right-left+1)
            right+=1
        return ans
            
            
