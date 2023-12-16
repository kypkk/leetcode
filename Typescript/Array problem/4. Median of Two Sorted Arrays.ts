function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
  let total_len: number = nums1.length + nums2.length;
  let idx_nums1: number = 0;
  let idx_nums2: number = 0;
  let ans_array: number[] = [];
  for (let i: number = 0; i < total_len; i++) {
    if (nums1[idx_nums1] != undefined && nums2[idx_nums2] != undefined) {
      if (nums1[idx_nums1] < nums2[idx_nums2]) {
        ans_array.push(nums1[idx_nums1]);
        idx_nums1 += 1;
      } else {
        ans_array.push(nums2[idx_nums2]);
        idx_nums2 += 1;
      }
    } else if (nums2[idx_nums2] == undefined) {
      ans_array.push(nums1[idx_nums1]);
      idx_nums1 += 1;
    } else {
      ans_array.push(nums2[idx_nums2]);
      idx_nums2 += 1;
    }
  }
  if (total_len % 2 == 1) {
    let ans: number = ans_array[Math.floor(total_len / 2)];

    return Number(ans.toFixed(5));
  } else {
    let ans: number =
      (ans_array[total_len / 2] + ans_array[total_len / 2 - 1]) / 2;
    return Number(ans.toFixed(5));
  }
}

console.log(findMedianSortedArrays([0, 0], [0, 0]));
