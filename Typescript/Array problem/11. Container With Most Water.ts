function maxArea(height: number[]): number {
  let right: number = height.length - 1;
  let left: number = 0;
  let maxarea = 0;
  while (left < right) {
    let width = right - left;
    maxarea = Math.max(maxarea, Math.min(height[left], height[right]) * width);
    if (height[left] < height[right]) left += 1;
    else right -= 1;
  }
  return maxarea;
}

const heigh: number[] = [1, 8, 6, 2, 5, 4, 8, 3, 7];

console.log(maxArea(heigh));
