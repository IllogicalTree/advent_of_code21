const fs = require('fs')

const readCommands = (filePath) => {
  let commands = [];
  try {
    const fileData = fs.readFileSync(filePath, 'utf8');
    commands = fileData.split(/\r?\n/).map(command => {
        let commandParts = command.split(' ');
        let direction = commandParts[0];
        let distance = parseInt(commandParts[1]);
        return {direction, distance};
    });
  } catch (err) {
    console.error(err);
  }
  return commands;
}

const part1 = commands => {
  let horizontal = 0;
  let depth = 0;
  commands.forEach(command => {
      switch (command.direction) {
        case 'forward':
            horizontal += command.distance;
            break
        case 'up':
            depth -= command.distance;
            break
        case 'down':
            depth += command.distance;
            break
      }; 
  });
  return horizontal * depth;
}

const part2 = commands => {
  let horizontal = 0;
  let depth = 0;
  let aim = 0;
  commands.forEach(command => {
    switch (command.direction) {
      case 'forward':
          horizontal += command.distance;
          depth += command.distance * aim;
          break
      case 'up':
          aim -= command.distance;
          break
      case 'down':
          aim += command.distance;
          break
    }; 
  });
  return horizontal * depth;
}

const main = async () => {
  const testInputData = await readCommands('./day_02/test_input.txt');
  console.log(`Part 1 (test): ${part1(testInputData)}`);
  console.log(`Part 2 (test): ${part2(testInputData)}`);

  const inputData = await readCommands('./day_02/input.txt');
  console.log(`Part 1: ${part1(inputData)}`);
  console.log(`Part 2: ${part2(inputData)}`);
}

main();

/*
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
*/

