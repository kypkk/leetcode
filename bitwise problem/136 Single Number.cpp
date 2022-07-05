#include<iostream>
#include<vector>

using namespace std;

int solution(vector<int>& nums);

int main(){

    vector<int> nums({4,4,2,2,3,1,1});
    cout << solution(nums);

    return 0;
}

int solution(vector<int>& nums){
    int ans = 0;
    for(auto& x: nums)
        ans ^= x;
    return ans;
}