/** @type {import("esbuild").BuildOptions} */
export default {
  logLevel: "info",
  entryPoints: ["src/index.tsx"],
  bundle: true,
  outdir: "dist",
  define: {
    "process.env.NODE_ENV": '"development"',
    "process.env.REACT_APP_STAKING_ADDR": '"0x959922be3caee4b8cd9a407cc3ac1c251c2007b1"',
    "process.env.REACT_APP_ETH_NODE_URL": '"http://0.0.0.0:8545"',
    "process.env.REACT_APP_MOR_ADDR": '"0x5fbdb2315678afecb367f032d93f642f64180aa3"',
    "process.env.REACT_APP_LMR_ADDR": '"0x0B306BF915C4d645ff596e518fAf3F9669b97016"',
  },
  inject: ["./src/react-shim.ts"],
  plugins: [],
  loader: {
    ".png": "file",
  },
};
