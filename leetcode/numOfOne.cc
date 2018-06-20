#include <iostream>
#include <math.h>
#include <string>

using namespace std;

class Solution {

public:
    int numOfOne(int num){
        int count = 0;
        int current = num;
        int rem;
        while(current > 0){
            rem = current%2;
            current = current >> 1;
            if(rem == 1){
                count++;
            }
        }
        return count;
    }


};

int stringToInteger(string input) {
    return stoi(input);
}

int main() {
    string line;
    while (getline(cin, line)) {
        int L = stringToInteger(line);
        int ret = Solution().numOfOne(L);

        string out = to_string(ret);
        cout << out << endl;
    }
    return 0;
}
