import { ethers } from 'hardhat';

async function main() {
  const qgbFactory = await ethers.getContractFactory('Peggy');
  const qgb = await qgbFactory.deploy();
  await qgb.deployed();
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error); // eslint-disable-line no-console
    process.exit(1);
  });
