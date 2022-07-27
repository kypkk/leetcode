function twoSum(nums: number[], target: number): number[] {
  let dict = {};
  for (let i: number = 0; i < nums.length; i++) {
    dict[nums[i]] = i;
  }
  for (let i: number = 0; i < nums.length; i++) {
    let tmp: number = nums[i];
    let goal: number = target - tmp;
    if (dict[goal] && dict[goal] != i) return [i, dict[goal]];
  }
}

console.log(twoSum([3, 2, 4], 6));
