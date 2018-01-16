//modules
const Canvas = require("canvas");
const fs = require("fs");
const {Grid} = require("./grid");


const filename = "langston.png";
const outputFile = fs.createWriteStream(__dirname + "/../output/" + filename);

// Create the play board
const width = 100,
      height = 100,
      steps = 50,
      num_ants = 3;

const grid = new Grid(width, height, num_ants);
grid.simulate(steps);

// Draw the image

// set up the canvas
const canvas = new Canvas(width, height);
const context = canvas.getContext("2d");
context.fillStyle = "black";
context.fillRect(0, 0, canvas.width, canvas.height);

let pngStream = canvas.pngStream();

pngStream.on("data", function (chunk) {
    outputFile.write(chunk);
});