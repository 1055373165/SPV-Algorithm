var createHash = require('crypto').createHash;

/**
 * @param {Buffer} buf1
 * @return {Buffer}
 */
//双 hash
function sha256x2 (buf1, buf2) {
  var buf = createHash('sha256').update(buf1).update(buf2).digest();
  return createHash('sha256').update(buf).digest();
}

/**
 * Reverse CURRENT buffer
 * @param {Buffer} buf
 * @return {Buffer}
 */
function reverse (buf) {
  return Array.prototype.reverse.call(buf);
}

/**
 * @param {Buffer} buf1
 * @param {Buffer} buf2
 * @return {Buffer}
 */
function isEqual (buf1, buf2) {
  if (buf1.length !== buf2.length) {
    return false;
  }

  for (var i = 0; i < buf1.length; ++i) {
    if (buf1[i] !== buf2[i]) {
      return false;
    }
  }

  return true;
}

/**
 * @typedef {Object} ProofObject
 * @param {string} txId
 * @param {number} txIndex
 * @param {string[]} sibling
 */

/**
 * @param {string[]} txIds
 * @param {number} txIndex
 * @return {ProofObject}
 */
exports.getProof = function (txIds, txIndex) {
  var proof = {
    txId: txIds[txIndex],
    txIndex: txIndex,
    sibling: []
  };

  var tree = new Array(txIds.length);
  for (var i = 0; i < tree.length; ++i) {
		// 小端存储
    tree[i] = reverse(new Buffer(txIds[i], 'hex'));
  }
  var target = tree[txIndex];
	
// 求 merkle path
  while (tree.length !== 1) {
		// 逐层求 hash，最终求出根 hash 值
    var newTree = new Array(~~((tree.length + 1) / 2));
    for (var j = 0; j < tree.length; j += 2) {
      var hash1 = tree[j];
			// 如果长度为偶数，那么最后一组 j+1 = tree.length-1，
			// 如果长度为奇数，那么最后一组 tree.length - 1 < j + 1, 重用最后一个 hash 值
      var hash2 = tree[Math.min(j + 1, tree.length - 1)];
      // 计算两个 hash 级联的 double sha256 值
      newTree[j / 2] = sha256x2(hash1, hash2);
			
			// 如果 hash1 是目标 hash 值，那么将 hash2 压入merkle path（小端方式）然后将 target 更新为级联后做 hash 的值，从而寻找上一层的 merkle path
      if (isEqual(target, hash1)) {
        proof.sibling.push(reverse(hash2).toString('hex'));
        target = newTree[j / 2];
      } else if (isEqual(target, hash2)) {
        proof.sibling.push(reverse(hash1).toString('hex'));
        target = newTree[j / 2];
      }
    }

    tree = newTree;
  }

  return proof;
};

/**
 * @param {ProofObject} proofObj
 * @return {string}
 */
// 根据 merkle proof 计算 merkle 根 hash 值
exports.getTxMerkle = function (proofObj) {
  var target = reverse(new Buffer(proofObj.txId, 'hex'));
  var txIndex = proofObj.txIndex;
  var sibling = proofObj.sibling;
	// txIndex 用来计算 target || sibling 计算出的 hash 对应的上一层的索引
  for (var i = 0; i < proofObj.sibling.length; ++i, txIndex = ~~(txIndex / 2)) {
		// 如果 txIndex 位于奇数索引，说明其位于一对 hash 中的后面，那么就将 sibling 中的 hash 放到前面
    if (txIndex % 2 === 1) {
      target = sha256x2(reverse(new Buffer(sibling[i], 'hex')), target);
    } else { 
      target = sha256x2(target, reverse(new Buffer(sibling[i], 'hex')));
    }
  }

  return reverse(target).toString('hex');
};

/**
 * @param {string[]} txIds
 * @return {string}
 */
exports.getMerkleRoot = function (txIds) {
// 使用副本数组（存储的是小端格式的数据）进行计算
  var tree = new Array(txIds.length);
  for (var i = 0; i < tree.length; ++i) {
    tree[i] = reverse(new Buffer(txIds[i], 'hex'));
  }

  while (tree.length !== 1) {
    var newTree = new Array(~~((tree.length + 1) / 2));
    for (var j = 0; j < tree.length; j += 2) {
      var hash1 = tree[j];
      var hash2 = tree[Math.min(j + 1, tree.length - 1)];

      newTree[j / 2] = sha256x2(hash1, hash2);
    }

    tree = newTree;
  }

  return reverse(tree[0]).toString('hex');
};