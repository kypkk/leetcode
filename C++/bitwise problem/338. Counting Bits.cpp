#include<iostream>
#include<vector>

using namespace std;

vector<int> solution(int n);

int main(){

    int n;
    cin >> n;
    vector<int> ans = solution(n);    
    for(auto &x: ans)
        cout << x << " ";
    cout << endl;
    return 0;
}

vector<int> solution(int n){
    vector<int> ans({});
    for(int i = 0; i < n + 1; i++){
        int32_t num = i;
        int count = 0;
        for(int j = 0; j < 32; j++){
            if((num & 0x00000001) == 1)
                count += 1;
            num = num >> 1;
        }
        ans.push_back(count);
    }
    return ans;
}