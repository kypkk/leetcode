#include<iostream>

using namespace std;

bool solution(int n);

int main(){
    int n;
    cin >> n;

    if(solution(n))
        cout << "true";
    else
        cout << "false";   


    return 0;
}

bool solution(int n){
    int tmp = ((n >> 1) ^ n);
    return (tmp & (tmp + 1)) == 0;
}