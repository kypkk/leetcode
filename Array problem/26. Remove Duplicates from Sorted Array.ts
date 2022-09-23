function removeDuplicates(nums: number[]): number {
  if (nums.length <= 1) {
    return nums.length;
  }
  let i: number = 0;
  for (let j: number = 1; j < nums.length; j++) {
    if (nums[i] != nums[j]) {
      i += 1;
      nums[i] = nums[j];
    }
  }
  return i + 1;
}

removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]);
// 0 1 2 2
