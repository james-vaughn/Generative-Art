//modules
const Canvas = require("canvas");
const fs = require("fs");
const {Grid} = require("./grid");
const {Color} = require("../lib/color");

const filename = "termite.png";
const outputImageFile = fs.createWriteStream(__dirname + "/../output/" + filename);

// Create the play board
const width = 1000,
      height = 1000,
      steps = 500000,
      num_ants = 25;

const grid = new Grid(width, height, num_ants);
grid.simulate(steps);


// Draw the image from the grid

// generate the colors
colors = [];
colors.push(Color.random());

for(let i = 0; i < num_ants - 1; i++) {
    colors.push(Color.mutation_of(colors[i]));
}

// set up the canvas
const canvas = new Canvas(width, height);
const context = canvas.getContext("2d");
context.fillStyle = `rgba(${colors[0].r}, ${colors[0].g}, ${colors[0].b}, ${colors[0].a})`;
context.fillRect(0, 0, canvas.width, canvas.height);

// smoothing
for (let smoothing_iter = 0; smoothing_iter < 5; smoothing_iter++) {
    for (let y = 0; y < height; y++) {
        for (let x = 0; x < width; x++) {
            grid.grid[y][x] = avg_ant_val(grid, x, y, 4);
        }
    }
}

// drawing
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


function avg_ant_val(grid, x, y, radius) {
    neighboring_ants = [];

    // add all of the neighboring pixels
    for(let y_val = Math.max(0, y - radius); y_val <= Math.min(height - 1, y + radius); y_val++) {
        for(let x_val = Math.max(0, x - radius); x_val <= Math.min(width -1, x + radius); x_val++) {
            neighboring_ants.push(grid.grid[y_val][x_val]);
        }
    }

    // grab the most common value
    neighboring_ants.sort();

    let max = 1,
        freq = 1,
        result = neighboring_ants[0];

    for(let i = 0; i < neighboring_ants.length-1; i++) {
        if(neighboring_ants[i] === neighboring_ants[i+1]) {
            freq += 1;
        } else {
            freq = 1;
        }

        if(freq > max) {
            result = neighboring_ants[i];
            max = freq;
        }
    }

    return result;
}

function drawPixel(context, x, y, color) {
    context.fillStyle = `rgba(${color.r}, ${color.g}, ${color.b}, ${color.a})`;
    context.fillRect(x, y, 1, 1);
}
