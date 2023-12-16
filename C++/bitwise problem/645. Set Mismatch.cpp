#include<iostream>
#include<vector>

using namespace std;

vector<int> findErrorNums(vector<int>& nums);

int main(){
    vector<int> nums({2, 2});
    vector<int> ans = findErrorNums(nums);
    for (auto &x: ans)
        cout << x << " ";
    cout << endl;
    return 0;
}

vector<int> findErrorNums(vector<int>& nums){
    vector<int> ans({});
    for(int i = 0; i < nums.size(); i++){
        if((nums[i] > i+1)){
            ans.push_back(i+2);
            ans.push_back(i+1);
            break;
        }else if((nums[i] < i+1)){
            ans.push_back(i);
            ans.push_back(i+1);
        }
    }
    return ans;
}