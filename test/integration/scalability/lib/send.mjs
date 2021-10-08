import { getAddress } from "./getAddress.mjs";

export async function send({
  node,
  chainId,
  binary,
  src,
  dst,
  denom,
  amount,
  fees,
  home = undefined,
  dryRun = false,
}) {
  if (!node) throw new Error("missing requirement argument: --node");
  if (!chainId) throw new Error("missing requirement argument: --chain-id");
  if (!binary) throw new Error("missing requirement argument: --binary");
  if (!src) throw new Error("missing requirement argument: --src");
  if (!dst) throw new Error("missing requirement argument: --dst");
  if (!denom) throw new Error("missing requirement argument: --denom");
  if (!amount) throw new Error("missing requirement argument: --amount");

  console.log(`
node    ${node}
chainId ${chainId}
binary  ${binary}
src     ${src}
dst     ${dst}
denom   ${denom}
amount  ${amount}
fees    ${fees}
`);

  const srcAddr = await getAddress({ binary, name: src, home });
  const dstAddr = await getAddress({ binary, name: dst, home });

  await $`
${binary} \
    tx \
    bank \
    send \
    ${srcAddr} \
    ${dstAddr} \
    ${amount}${denom} \
    ${fees ? `--fees` : ``} ${fees ? `${fees}${denom}` : ``} \
    --keyring-backend test \
    --node ${node} \
    --chain-id ${chainId} \
    ${!dryRun ? "--yes" : ""} \
    ${home ? `--home` : ``} ${home ? `${home}` : ``} \
    --broadcast-mode block
    `;
}
