#include<iostream>

using namespace std;

uint32_t solution(uint32_t n);

int main(){

    uint32_t n = 0xffffffff;
    cout << solution(n);    

    return 0;
}

uint32_t solution(uint32_t n){
    int ans = 0;
    for(int i = 0; i < 31; i++){
        if((n & 0x00000001) == 1)
            ans += 1;
        n = n >> 1;
    }
    if((n & 0x00000001) == 1)
        ans += 1;
    return ans;
    
}