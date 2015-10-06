/**
 * Created by Michael on 9/14/2015.
 */
//Sample code to read in test cases:
var fs = require("fs");

if (process.argv[2].length < 1) {
    throw new Error('Include filename as first parameter');
}
fs.readFileSync(process.argv[2]).toString().split('\n').forEach(function (line) {
    if (line != "") {
        var parsed = line.split(' ');
        console.log(FizzBuzz(parsed[0], parsed[1], parsed[2]));
    }
});

function FizzBuzz(f, b, l) {
    if (f < 1 || f > 20) {
        throw new Error("Fizz is out of bounds.  range [1, 20]");
        return;
    }
    if (b < 1 || b > 20) {
        throw new Error("Buzz is out of bounds.  range [1, 20]");
        return;
    }
    if (l < 21 || l > 100) {
        throw new Error("Length is out of bounds.  range [21, 100]");
        return;
    }
    var return_string = '';
    for (var i = 1; i <= l; i++) {
        if (i%f === 0 && i%b === 0) {
            return_string += 'FB ';
        } else if (i%f === 0) {
            return_string += 'F ';
        } else if (i%b === 0) {
            return_string += 'B ';
        } else {
            return_string += i + ' ';
        }
    }
    return return_string;
}
