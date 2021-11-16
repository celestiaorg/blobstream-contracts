import { HardhatUserConfig } from 'hardhat/types';
import '@nomiclabs/hardhat-waffle';
import 'hardhat-deploy';
import 'solidity-coverage';

const config: HardhatUserConfig = {
  defaultNetwork: 'hardhat',
  solidity: {
    compilers: [
      {
        version: '0.8.10',
      },
    ],
  },
  networks: {
    coverage: {
      url: 'http://127.0.0.1:8555',
    },
  },
};

export default config;
