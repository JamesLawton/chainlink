// SPDX-License-Identifier: MIT
pragma solidity 0.8.6;

import {iKeeperRegistryExecutable} from "../interfaces/iKeeperRegistry.sol";
import {ConfirmedOwner} from "../ConfirmedOwner.sol";

/**
 * @notice This contract serves as a wrapper around a keeper registry's checkUpkeep function.
 */
contract KeeperRegistryCheckUpkeepGasUsageWrapper is ConfirmedOwner {
  iKeeperRegistryExecutable private immutable i_keeperRegistry;

  /**
   * @param keeperRegistry address of a keeper registry
   */
  constructor(iKeeperRegistryExecutable keeperRegistry) ConfirmedOwner(msg.sender) {
    i_keeperRegistry = keeperRegistry;
  }

  /**
   * @return the keeper registry
   */
  function getKeeperRegistry() external view returns (iKeeperRegistryExecutable) {
    return i_keeperRegistry;
  }

  /**
   * @notice This function is called by monitoring service to estimate how much gas checkUpkeep functions will consume.
   * @param id identifier of the upkeep to check
   * @param from the address to simulate performing the upkeep from
   */
  function measureCheckGas(uint256 id, address from)
    external
    returns (
      bool,
      bytes memory,
      uint256
    )
  {
    uint256 startGas = gasleft();
    try i_keeperRegistry.checkUpkeep(id, from) returns (
      bytes memory performData,
      uint256 maxLinkPayment,
      uint256 gasLimit,
      uint256 adjustedGasWei,
      uint256 linkEth
    ) {
      uint256 gasUsed = startGas - gasleft();
      return (true, performData, gasUsed);
    } catch {
      uint256 gasUsed = startGas - gasleft();
      return (false, "", gasUsed);
    }
  }
}
