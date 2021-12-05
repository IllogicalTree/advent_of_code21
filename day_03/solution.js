const fs = require('fs')

const readDiagnostics = (filePath) => {
  let diagnostics = [];
  try {
    const fileData = fs.readFileSync(filePath, 'utf8');
    diagnostics = fileData.split(/\r?\n/);
  } catch (err) {
    console.error(err);
  }
  return diagnostics;
}

const mostCommonBit = (diagnostics, position) => {
  let totalOneBits = 0; 
  let totalZeroBits = 0;
  diagnostics.forEach(diagnostic => {
    let bitAtPosition = diagnostic[position];
    bitAtPosition == 1 ? totalOneBits++ : totalZeroBits++;
  })
  return totalOneBits >= totalZeroBits ? 1 : 0;
}
const leastCommonBit = (diagonstics, position) => mostCommonBit(diagonstics, position) === 1 ? 0 : 1;

const diagnosticsWithBitAtColumn = (diagnostics, column, bit) => {
  let newDiagnostics = [];
  diagnostics.forEach(diagnostic => {
    if (diagnostic[column] == bit) {
      newDiagnostics.push(diagnostic);
    }
  });
  return newDiagnostics;
}

const part1 = diagnostics => {
  let gammaRate = '';
  let epsilonRate = '';
  for (let i = 0; i < diagnostics[0].length; i++) {
    gammaRate += mostCommonBit(diagnostics, i);
    epsilonRate += leastCommonBit(diagnostics, i);
  }
  gammaRate = parseInt(gammaRate, 2);
  epsilonRate = parseInt(epsilonRate, 2);
  return gammaRate * epsilonRate;
}

const part2 = diagnostics => {
  let oxygenValues = diagnostics;
  let co2Values = diagnostics;
  let length = diagnostics[0].length;

  for (let i = 0; i < length; i++) {
    oxygenValues = diagnosticsWithBitAtColumn(oxygenValues, i, mostCommonBit(oxygenValues, i));
    if (oxygenValues.length === 1) {
      break
    }
  }

  for (let i = 0; i < length; i++) {
    co2Values = diagnosticsWithBitAtColumn(co2Values, i, leastCommonBit(co2Values, i));
    if (co2Values.length === 1) {
      break
    }
  }

  oxygenRating = oxygenValues[0]
  co2Rating = co2Values[0]
  oxygenRating = parseInt(oxygenRating, 2);
  co2Rating = parseInt(co2Rating, 2);

  return oxygenRating * co2Rating;
}

const main = async () => {
  const testInputData = await readDiagnostics('./day_03/test_input.txt');
  console.log(`Part 1 (test): ${part1(testInputData)}`);
  console.log(`Part 2 (test): ${part2(testInputData)}`);

  const inputData = await readDiagnostics('./day_03/input.txt');
  console.log(`Part 1: ${part1(inputData)}`);
  console.log(`Part 2: ${part2(inputData)}`);
}

main();
