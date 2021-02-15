import { Peggy } from "../typechain";
import { ethers } from "ethers";
import fs from "fs";
import commandLineArgs from "command-line-args";


const args = commandLineArgs([
  // the ethernum node used to deploy the contract
  { name: "eth-node", type: String },
  // the Ethereum private key that will contain the gas required to pay for the contact deployment
  { name: "eth-privkey", type: String },
]);

async function deposit() {
  const provider = await new ethers.providers.JsonRpcProvider(args["eth-node"]);
  let wallet = new ethers.Wallet(args["eth-privkey"], provider);

  let proposalPath = './scripts/ERC20Tokens.json'

  let {erc20TestAddress, erc20TestAddress1, erc20TestAddress2, peggy} = JSON.parse(fs.readFileSync(proposalPath, "utf8"));

  let Peggy;
  {
    let { abi } = getContractArtifacts("./artifacts/contracts/Peggy.sol/Peggy.json");
    Peggy = new ethers.Contract(peggy, abi, wallet)
  }

  let { abi } = getContractArtifacts("./artifacts/contracts/TestERC20A.sol/TestERC20A.json");
  let erc20 = new ethers.Contract(erc20TestAddress, abi, wallet);
  const approvalAmount = "99999999999999999999999999999999";
  console.log(`Approving ${approvalAmount} of ERC-20 token ${erc20TestAddress} to Peggy at ${peggy}`)
  await erc20.approve(peggy, ethers.utils.parseEther(approvalAmount))
  erc20 = new ethers.Contract(erc20TestAddress1, abi, wallet);
  console.log(`Approving ${approvalAmount} of ERC-20 token ${erc20TestAddress1} to Peggy at ${peggy}`)
  await erc20.approve(peggy, ethers.utils.parseEther(approvalAmount))
  erc20 = new ethers.Contract(erc20TestAddress2, abi, wallet);
  console.log(`Approving ${approvalAmount} of ERC-20 token ${erc20TestAddress2} to Peggy at ${peggy}`)
  await erc20.approve(peggy, ethers.utils.parseEther(approvalAmount))

  let tokenASendAmount = 69;
  let tokenBSendAmount = 79;
  let tokenCSendAmount = 89;
  console.log(`Send To Cosmos ${tokenASendAmount} of ERC-20 token ${erc20TestAddress} for ${wallet.address} to Peggy at ${peggy}`)
  await Peggy.sendToCosmos(erc20TestAddress, wallet.address, tokenASendAmount);
  console.log(`Send To Cosmos ${tokenBSendAmount} of ERC-20 token ${erc20TestAddress1} for ${wallet.address} to Peggy at ${peggy}`)
  await Peggy.sendToCosmos(erc20TestAddress1, wallet.address, tokenBSendAmount);
  console.log(`Send To Cosmos ${tokenCSendAmount} of ERC-20 token ${erc20TestAddress2} for ${wallet.address} to Peggy at ${peggy}`)
  await Peggy.sendToCosmos(erc20TestAddress2, wallet.address, tokenCSendAmount);
}

function getContractArtifacts(path: string): { bytecode: string; abi: string } {
  const {bytecode, abi} = JSON.parse(fs.readFileSync(path, "utf8").toString());
  return { bytecode, abi };
}

async function main() {
  await deposit();
}

main();
