class Solution {
public:
    string simplifyPath(string path) {
        stack<string> dirSt;
        string canonicalPath = "";
        string dir;
        int n = path.length();
        
        for(int i = 0; i < n; ++i) {
            if(path[i] == '/') continue;
            dir = "";
            while(i < n && path[i] != '/') {
                dir += path[i++];
            }
            if(dir == ".") continue;
            else if(dir == "..") {
                if(!dirSt.empty()) dirSt.pop();
            }
            else dirSt.push(dir);
        }
        
        while(!dirSt.empty()) {
            canonicalPath = "/" + dirSt.top() + canonicalPath;
            dirSt.pop();
        }

        if(canonicalPath == "") canonicalPath = "/";
        
        return canonicalPath;
    }
};
