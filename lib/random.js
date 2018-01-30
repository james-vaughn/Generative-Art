function randInt(max, min = 0) {
    return Math.floor(Math.random() * max) + min;
}

function randSign() {
    return Math.random() < .5 ? -1 : 1;
}

function randSignedInt(max, min = 0) {
    return randSign() * (randInt(max, min));
}

exports.randInt = randInt;
exports.randSign = randSign;
exports.randSignedInt = randSignedInt;