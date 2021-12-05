const fs = require('fs');

const readDepths = (filePath) => {
  let depths = [];
  try {
    const fileData = fs.readFileSync(filePath, 'utf8');
    depths = fileData.split(/\r?\n/).map(depth => Number(depth));
  } catch (err) {
    console.error(err);
  }
  return depths;
}

const countSlidingWindowIncrease = (depths, windowSize) => {
  let count = 0;
  for (let i = 0; i < depths.length - windowSize; i++) {
    if (depths[i + windowSize] > depths[i]) {
      count++;
    }
  }
  return count;
}

const part1 = depths => countSlidingWindowIncrease(depths, 1);
const part2 = depths => countSlidingWindowIncrease(depths, 3);

const main = async () => {
  const testInputData = await readDepths('./day_01/test_input.txt');
  console.log(`Part 1 (test): ${part1(testInputData)}`);
  console.log(`Part 2 (test): ${part2(testInputData)}`);

  const inputData = await readDepths('./day_01/input.txt');
  console.log(`Part 1: ${part1(inputData)}`);
  console.log(`Part 2: ${part2(inputData)}`);
}

main();