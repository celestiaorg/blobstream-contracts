// TODO not working in assembly script
// import { bech32 } from "bech32";
// import { Address as EthereumJSAddress } from "ethereumjs-util";

import {
  Address,
  BigInt,
  Bytes,
  dataSource,
  log,
} from "@graphprotocol/graph-ts";

import {
  ERC20DeployedEvent,
  SendToCosmosEvent,
  SubmitBatchCall,
  TransactionBatchExecutedEvent,
  ValsetUpdatedEvent,
} from "../generated/Peggy/Peggy";
import {
  DeployedERC20,
  Deposit,
  State,
  StateLastBatchNonces,
  Valset,
  BatchWithdrawal,
  Withdrawal,
} from "../generated/schema";

// setting peggy id directly as using call handler initialize was causing problems
let STATE_PEGGY_ID_KOVAN = Bytes.fromHexString(
  "0x696e6a6563746976652d70656767796964000000000000000000000000000000"
);
let STATE_PEGGY_ID_MAINNET = Bytes.fromHexString(
  "0x696e6a6563746976652d70656767796964000000000000000000000000000000"
);

let STATE_STORE_ID = "1";

function getCelestiaAddress(address: Bytes): string {
  return address.toHexString();

  // not working in assembly script
  // const addressBuffer = EthereumJSAddress.fromString(
  //   address.toHexString()
  // ).toBuffer();
  // return bech32.encode("inj", bech32.toWords(addressBuffer));
}

// below code not working, creating different checkpoints than Peggy contract for unknown reason
// function getCheckpoint(
//   peggyId: Bytes,
//   valsetNonce: BigInt,
//   validators: Address[],
//   powers: BigInt[],
//   rewardAmount: BigInt,
//   rewardToken: Address
// ): ByteArray {
//   // bytes32 encoding of the string "checkpoint"
//   let methodName = Bytes.fromHexString(
//     "0x636865636b706f696e7400000000000000000000000000000000000000000000"
//   ) as Bytes;

//   let tupleArray: Array<ethereum.Value> = [
//     ethereum.Value.fromBytes(peggyId),
//     ethereum.Value.fromBytes(methodName),
//     ethereum.Value.fromUnsignedBigInt(valsetNonce),
//     ethereum.Value.fromAddressArray(validators),
//     ethereum.Value.fromUnsignedBigIntArray(powers),
//     ethereum.Value.fromUnsignedBigInt(rewardAmount),
//     ethereum.Value.fromAddress(rewardToken),
//   ];

//   let encodedCheckpointData = ethereum.encode(
//     ethereum.Value.fromTuple(tupleArray as ethereum.Tuple)
//   )!;
//   let checkpoint = crypto.keccak256(encodedCheckpointData);

//   return checkpoint;
// }

function updateState(
  validators: Address[],
  powers: BigInt[],
  rewardAmount: BigInt,
  rewardToken: Address,
  lastEventNonce: BigInt,
  lastValsetNonce: BigInt
): void {
  let state = new State(STATE_STORE_ID);
  state.peggyId =
    dataSource.network() == "mainnet"
      ? (STATE_PEGGY_ID_MAINNET as Bytes)
      : (STATE_PEGGY_ID_KOVAN as Bytes);

  // let checkpoint = getCheckpoint(
  //   state.peggyId,
  //   lastValsetNonce,
  //   validators,
  //   powers,
  //   rewardAmount,
  //   rewardToken
  // );
  // state.lastValsetCheckpoint = checkpoint as Bytes;

  state.lastEventNonce = lastEventNonce.toI32();
  state.lastValsetNonce = lastValsetNonce.toI32();

  state.save();
}

function updateLastBatchNoncesState(
  tokenAddress: Address,
  lastBatchNonce: BigInt
): void {
  let stateLastBatchNonces = new StateLastBatchNonces(tokenAddress.toHex());
  stateLastBatchNonces.nonce = lastBatchNonce.toI32();
  stateLastBatchNonces.save();
}

export function handleNewERC20Deployed(event: ERC20DeployedEvent): void {
  log.info("handleNewERC20Deployed: {} at token with address {}", [
    event.params._symbol,
    event.params._tokenContract.toHex(),
  ]);

  let erc20ID =
    event.transaction.hash.toHex() + "-" + event.logIndex.toString();
  let erc20 = new DeployedERC20(erc20ID);

  erc20.cosmosDenom = event.params._cosmosDenom;
  erc20.tokenContract = event.params._tokenContract;
  erc20.name = event.params._name;
  erc20.symbol = event.params._symbol;
  erc20.decimals = event.params._decimals;

  erc20.save();

  let state = new State(STATE_STORE_ID);
  state.lastEventNonce = event.params._eventNonce.toI32();
  state.save();
}

export function handleSendToCosmosEvent(event: SendToCosmosEvent): void {
  log.info(
    "handleSendToCosmosEvent: Withdrawal of {} at token with address {}",
    [event.params._amount.toString(), event.params._tokenContract.toHex()]
  );

  let depositID =
    event.transaction.hash.toHex() + "-" + event.logIndex.toString();
  let deposit = new Deposit(depositID);

  deposit.tokenContract = event.params._tokenContract;
  deposit.amount = event.params._amount;
  deposit.destination = getCelestiaAddress(event.params._destination);
  deposit.sender = event.params._sender;
  deposit.timestamp = event.block.timestamp.toI32();
  deposit.blockHeight = event.block.number.toI32();

  deposit.save();

  let state = new State(STATE_STORE_ID);
  state.lastEventNonce = event.params._eventNonce.toI32();
  state.save();
}

export function handleUpdateValset(event: ValsetUpdatedEvent): void {
  log.info("handleUpdateValset: New valset length {} ", [
    BigInt.fromI32(event.params._validators.length).toString(),
  ]);

  let valsetID =
    event.transaction.hash.toHex() + "-" + event.logIndex.toString();
  let valset = new Valset(valsetID);

  valset.valsetNonce = event.params._newValsetNonce;
  valset.rewardAmount = event.params._rewardAmount;
  valset.rewardToken = event.params._rewardToken;
  valset.powers = event.params._powers;

  let addressTypeValidators = event.params._validators;
  let validators = new Array<Bytes>(event.params._validators.length);
  for (let i = 0; i < addressTypeValidators.length; i++) {
    validators[i] = addressTypeValidators[i] as Bytes;
  }
  valset.validators = validators;

  let bigIntTypePowers = event.params._powers;
  let powers = new Array<BigInt>(event.params._powers.length);
  for (let i = 0; i < bigIntTypePowers.length; i++) {
    powers[i] = bigIntTypePowers[i] as BigInt;
  }
  valset.powers = powers;
  valset.timestamp = event.block.timestamp.toI32();
  valset.blockHeight = event.block.number.toI32();

  valset.save();

  updateState(
    event.params._validators,
    event.params._powers,
    event.params._rewardAmount,
    event.params._rewardToken,
    event.params._eventNonce,
    event.params._newValsetNonce
  );
}

export function handleSubmitBatchEvent(
  event: TransactionBatchExecutedEvent
): void {
  let state = new State(STATE_STORE_ID);
  state.lastEventNonce = event.params._eventNonce.toI32();
  state.save();
}

export function handleSubmitBatch(call: SubmitBatchCall): void {
  log.info(
    "handleSubmitBatch: Withdrawal with nonce {} of token with address {}",
    [call.inputs._batchNonce.toString(), call.inputs._tokenContract.toHex()]
  );

  let batchWithdrawalID =
    call.transaction.hash.toHex() + "-" + call.inputs._batchNonce.toString();
  let batchWithdrawal = new BatchWithdrawal(batchWithdrawalID);

  batchWithdrawal.amounts = call.inputs._amounts;
  batchWithdrawal.batchNonce = call.inputs._batchNonce;
  batchWithdrawal.fees = call.inputs._fees;
  batchWithdrawal.tokenContract = call.inputs._tokenContract;

  let addressTypeDestinations = call.inputs._destinations;
  let destinations = new Array<string>(call.inputs._destinations.length);
  for (let i = 0; i < addressTypeDestinations.length; i++) {
    destinations[i] = getCelestiaAddress(addressTypeDestinations[i] as Bytes);
  }
  batchWithdrawal.destinations = destinations;
  batchWithdrawal.sender = call.from;

  let totalFee = BigInt.fromI32(0);
  let fees = call.inputs._fees;

  for (let i = 0; i < fees.length; i++) {
    totalFee = totalFee.plus(fees[i]);
  }
  batchWithdrawal.totalFee = totalFee;
  batchWithdrawal.timestamp = call.block.timestamp.toI32();
  batchWithdrawal.blockHeight = call.block.number.toI32();

  batchWithdrawal.save();

  let amounts = call.inputs._amounts;

  for (let i = 0; i < amounts.length; i++) {
    let withdrawalID =
      call.transaction.hash.toHex() +
      "-" +
      call.inputs._batchNonce.toString() +
      "-" +
      i.toString();
    let withdrawal = new Withdrawal(withdrawalID);

    withdrawal.amount = amounts[i];
    withdrawal.destination = getCelestiaAddress(
      addressTypeDestinations[i] as Bytes
    );
    withdrawal.fee = fees[i];
    withdrawal.tokenContract = call.inputs._tokenContract;
    withdrawal.timestamp = call.block.timestamp.toI32();
    withdrawal.blockHeight = call.block.number.toI32();

    withdrawal.save();
  }

  updateLastBatchNoncesState(
    call.inputs._tokenContract,
    call.inputs._batchNonce
  );
}
