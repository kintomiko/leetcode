#include <iostream>
#include <math.h>
#include <string>

using namespace std;

class Solution {

public:
    bool isPrime(int num){
        for(int i=2;i<=sqrt(num);i++){
            if(num % i == 0){
                return true;
            }
        }
        return false;
    }

};

int stringToInteger(string input) {
    return stoi(input);
}

int main() {
    if(true == 0){
	cout << "true is 0";
    }
    string line;
    while (getline(cin, line)) {
        int L = stringToInteger(line);
        bool ret = Solution().isPrime(L);

        string out = to_string(ret);
        cout << out << endl;
    }
    return 0;
}
