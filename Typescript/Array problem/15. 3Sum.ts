function mergeSort(array) {
  if (array.length === 1) {
    return array;
  }
  const length = array.length;
  const middle = Math.floor(length / 2);
  const left = array.slice(0, middle);
  const right = array.slice(middle);
  return merge(mergeSort(left), mergeSort(right));
}
function merge(left, right) {
  const result = [];
  let leftIndex: number = 0;
  let rightIndex: number = 0;
  while (leftIndex < left.length && rightIndex < right.length) {
    if (left[leftIndex] < right[rightIndex]) {
      result.push(left[leftIndex]);
      leftIndex++;
    } else {
      result.push(right[rightIndex]);
      rightIndex++;
    }
  }
  return result.concat(left.slice(leftIndex)).concat(right.slice(rightIndex));
}

function threeSum(nums: number[]): number[][] {
  let answer: number[][] = [];
  nums = mergeSort(nums);

  let target_idx: number = 0;
  for (; target_idx < nums.length; target_idx++) {
    let target: number = -nums[target_idx];
    let left = target_idx + 1;
    let right = nums.length - 1;
    while (left < right) {
      let breaktrigger = 0;
      if (nums[left] + nums[right] < target) left += 1;
      else if (nums[left] + nums[right] > target) right -= 1;
      else {
        let x: number[] = [-target, nums[left], nums[right]];
        answer.forEach((arr) => {
          if (arr[1] == x[1] && arr[2] == x[2]) {
            breaktrigger = 1;
          }
        });
        left += 1;
        if (breaktrigger == 0) answer.push(x);
      }
    }
  }

  return answer;
}

console.log(threeSum([-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4]));
