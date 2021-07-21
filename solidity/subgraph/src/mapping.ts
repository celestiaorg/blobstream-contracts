import { BigInt, Bytes, log } from "@graphprotocol/graph-ts";
import {
  ERC20DeployedEvent,
  SendToCosmosEvent,
  SubmitBatchCall,
  UpdateValsetCall,
} from "../generated/Peggy/Peggy";
import {
  DeployedERC20,
  Deposit,
  Valset,
  Withdrawal,
} from "../generated/schema";

// TODO tracking the up-to-date state variables for convenience

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
}

export function handleSendToCosmosEvent(event: SendToCosmosEvent): void {
  log.debug(
    "handleSendToCosmosEvent: Withdrawal of {} at token with address {}",
    [event.params._amount.toString(), event.params._tokenContract.toHex()]
  );

  let depositID =
    event.transaction.hash.toHex() + "-" + event.logIndex.toString();
  let deposit = new Deposit(depositID);

  deposit.tokenContract = event.params._tokenContract;
  deposit.amount = event.params._amount;
  deposit.destination = event.params._destination;
  deposit.sender = event.params._sender;

  deposit.save();
}

export function handleUpdateValset(call: UpdateValsetCall): void {
  log.info("handleUpdateValset: New valset length {} ", [
    BigInt.fromI32(call.inputs._newValset.length).toString(),
  ]);

  let valsetID = call.transaction.hash.toHex() + "-call";
  let valset = new Valset(valsetID);

  valset.valsetNonce = call.inputs._currentValset.valsetNonce;
  valset.rewardAmount = call.inputs._currentValset.rewardAmount;
  valset.rewardToken = call.inputs._currentValset.rewardToken;
  valset.powers = call.inputs._currentValset.powers;

  let addressTypeValidators = call.inputs._currentValset.validators;
  let validators = new Array<Bytes>(
    call.inputs._currentValset.validators.length
  );
  for (let i = 0; i < addressTypeValidators.length; i++) {
    validators[i] = addressTypeValidators[i] as Bytes;
  }
  valset.validators = validators;

  let bigIntTypePowers = call.inputs._currentValset.powers;
  let powers = new Array<BigInt>(call.inputs._currentValset.powers.length);
  for (let i = 0; i < bigIntTypePowers.length; i++) {
    powers[i] = bigIntTypePowers[i] as BigInt;
  }
  valset.powers = powers;

  valset.save();
}

export function handleSubmitBatch(call: SubmitBatchCall): void {
  log.info(
    "handleSubmitBatch: Withdrawal with nonce {} of token with address {}",
    [call.inputs._batchNonce.toString(), call.inputs._tokenContract.toHex()]
  );

  let withdrawalID = call.transaction.hash.toHex() + "-call";
  let withdrawal = new Withdrawal(withdrawalID);

  withdrawal.amounts = call.inputs._amounts;
  withdrawal.batchNonce = call.inputs._batchNonce;
  withdrawal.fees = call.inputs._fees;
  withdrawal.tokenContract = call.inputs._tokenContract;

  let addressTypeDestinations = call.inputs._destinations;
  let destinations = new Array<Bytes>(call.inputs._destinations.length);
  for (let i = 0; i < addressTypeDestinations.length; i++) {
    destinations[i] = addressTypeDestinations[i] as Bytes;
  }
  withdrawal.destinations = destinations;

  withdrawal.save();
}
