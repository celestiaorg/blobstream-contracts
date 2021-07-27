require("@nomiclabs/hardhat-waffle");

task("accounts", "Prints the list of accounts", async () => {
  const accounts = await ethers.getSigners();

  console.log("Available Accounts");
  console.log("==================");
  for (const account of accounts) {
    console.log(account.address);
  }

  console.log("");

  console.log("HD Wallet");
  console.log("==================");
  console.log(
    "Mnemonic:      concert load couple harbor equip island argue ramp clarify fence smart topic"
  );
  console.log("Base HD Path:  m/44'/60'/0'/0/{account_index}");
  console.log("==================");
});

module.exports = {
  solidity: "0.8.2",

  paths: {
    root: "../..",
    cache: "test/ethereum/cache",
    artifacts: "test/ethereum/artifacts",
    sources: "solidity/contracts",
  },

  networks: {
    hardhat: {
      chainId: 50,
      hardfork: "muirGlacier",
      loggingEnabled: true,
      allowUnlimitedContractSize: true,
      blockGasLimit: 10000000,
      accounts: {
        mnemonic:
          "concert load couple harbor equip island argue ramp clarify fence smart topic",
        initialIndex: 0,
        path: "m/44'/60'/0'/0",
        count: 10,
        accountsBalance: "100000000000000000000",
      },
    },
  },
};
