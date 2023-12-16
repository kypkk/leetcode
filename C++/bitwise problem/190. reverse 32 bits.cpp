#include<iostream>

using namespace std;

uint32_t solution(uint32_t n);

int main(){

    uint32_t n = 32768;
    cout << solution(n);    

    return 0;
}

uint32_t solution(uint32_t n){
    n = (n << 16) | (n >> 16);
    n = ((n & 0xff00ff00) >> 8 | (n & 0x00ff00ff) << 8);
    n = ((n & 0xf0f0f0f0) >> 4 | (n & 0x0f0f0f0f) << 4);
    n = ((n & 0xcccccccc) >> 2 | (n & 0x33333333) << 2);
    n = ((n & 0xaaaaaaaa) >> 1 | (n & 0x55555555) << 1);

    return n;
    
}