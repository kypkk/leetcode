#include<iostream>

using namespace std;

int solution(int n);
int polluteMSBRight(int n);

int main(){
    int n;
    cin >> n;
    cout << solution(n);    

    return 0;
}

int polluteMSBRight(int n){
    n |= n >> 1;
    n |= n >> 2;
    n |= n >> 4;
    n |= n >> 8;
    n |= n >> 16;
    return n;
}

int solution(int n){
    if(n ==0)
        return 1;
    return (polluteMSBRight(n) ^ n);
}