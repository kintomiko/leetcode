#include <iostream>
#include <math.h>
#include <string>

using namespace std;

class Solution {
public:
    int countPrimeSetBits(int L, int R) {
        int count = 0;
        for(int i=L;i<=R;i++){
            if(isPrime(numOfOne(i))){
                count++;
            }
        }
        return count;
    }
    
private:
    bool isPrime(int num){
	if(num == 2 || num == 3){
	    return true;
	}
        if(num == 1){
	    return false;
	}
        for(int i=2;i<=sqrt(num);i++){
            if(num % i == 0){
                return false;
            }
        }
        return true;
    }
    
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
        getline(cin, line);
        int R = stringToInteger(line);
        
        int ret = Solution().countPrimeSetBits(L, R);

        string out = to_string(ret);
        cout << out << endl;
    }
    return 0;
}
