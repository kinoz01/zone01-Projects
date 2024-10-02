const blockChain = (data, prev = { index: 0, hash: "0" }) => {
    const block = {
      index: prev.index + 1,
      hash: hashCode(`${prev.index + 1}${prev.hash}${JSON.stringify(data)}`),
      data,
      prev,
      chain: newData => blockChain(newData, block)
    };
    return block;
};
  