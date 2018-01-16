//modules
const Canvas = require("canvas");
const fs = require("fs");
const {Grid} = require("./grid");

const filename = "langston.png";
const outputImageFile = fs.createWriteStream(__dirname + "/../output/" + filename);

// Create the play board
const width = 1000,
      height = 1000,
      steps = 100000000,
      num_ants = 50;

const grid = new Grid(width, height, num_ants);
grid.simulate(steps);


// Draw the image from the grid

class Color{
    constructor(r, g, b, a) {
        this.r = r;
        this.g = g;
        this.b = b;
        this.a = a;
    }

    static random() {
        return new Color(Math.floor(Math.random() * 255),
            Math.floor(Math.random() * 255),
            Math.floor(Math.random() * 255),
            1);
    }
}

// set up the canvas
const canvas = new Canvas(width, height);
const context = canvas.getContext("2d");
context.fillStyle = "black";
context.fillRect(0, 0, canvas.width, canvas.height);

colors = [];

for(let i = 0; i < num_ants; i++) {
    colors.push(Color.random());
}

for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
        const ant = grid.grid[y][x];

        if (ant === -1) continue;

        const color = colors[ant];
        drawPixel(context, x, y, color);
    }
}


let pngStream = canvas.pngStream();

pngStream.on("data", function (chunk) {
    outputImageFile.write(chunk);
});


function drawPixel(context, x, y, color) {
    context.fillStyle = `rgba(${color.r}, ${color.g}, ${color.b}, ${color.a})`;
    context.fillRect(x, y, 1, 1);
}
