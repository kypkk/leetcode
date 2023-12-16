#include<iostream>

using namespace std;

bool solution(int n);

int main(){
    int n;
    cin >> n;
    // cout << res;
    string ans;
    if(solution(n))
        ans = "true";
    else 
        ans = "false";
    cout << ans;
    return 0;
}

bool solution(int n){
    return (n > 0) && ((n & -n) == n) && ((n & 0xaaaaaaaa) == 0);
}
