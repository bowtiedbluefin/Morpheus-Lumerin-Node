'use strict';
const axios = require('axios').default;

// Mostly follows etherscan api definition except for few cases
// Decided to have a code-duplication instead of writing complex 
// mapping logic
class BlockscoutApi {
  constructor({ baseURL }) {
    this.api = axios.create({baseURL})
  }

  /**
   * Returns a list of ERC20 Token Transfer Events hashes by Address
   * @param {string} from start block
   * @param {string} to end block
   * @param {string} address wallet address
   * @param {string} tokenAddress  address
   * @param {number} [page] page number
   * @param {number} [pageSize] page size
   * @returns {Promise<string[]>} array of transaction hashes
   */
  async getTokenTransactions(from, to, address, tokenAddress, page = 1, pageSize = 10) {
    const params = {
      module: 'account',
      action: 'tokentx',
      sort: 'desc',
      contractaddress: tokenAddress,
      start_block: from,
      end_block: to,
      address,
      page,
      offset: pageSize
    }
    const { data } = await this.api.get('', { params })
    const { status, message, result } = data
    if (status !== '1' && message !== 'No transactions found') {
      throw new Error(result)
    }
    return result
  }

  /**
   * Returns a list of transactions for a specific address
   * @param {string} from start block
   * @param {string} to end block
   * @param {string} address wallet address
   * @param {number} [page] page number
   * @param {number} [pageSize] page size
   * @returns {Promise<string[]>} array of transaction hashes
   */
  async getEthTransactions(from, to, address, page = 1, pageSize = 1000) {
    const params = {
      module: 'account',
      action: 'txlist',
      sort: 'desc',
      start_block: from,
      end_block: to,
      address,
      page,
      offset: pageSize
    }

    const { data } = await this.api.get('', { params })

    const { status, message, result } = data
    if (status !== '1' && message !== 'No transactions found') {
      throw new Error(result)
    }

    return result
  }
}

module.exports = { BlockscoutApi };