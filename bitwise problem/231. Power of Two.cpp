#include<iostream>

using namespace std;

bool solution(int n);

int main(){

    int n;
    cin >> n;
    if(solution(n))
        cout << "true" << endl;
    else
        cout << "false" << endl;    

    return 0;
}

bool solution(int n){
    long long l = n;
    return l && ((l & -l) == l);
}