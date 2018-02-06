//modules
const Canvas = require("canvas");
const fs = require("fs");
const {Grid} = require("./grid");
const {Color} = require("../lib/color");
//const {randInt} = require("../lib/random");

const backgroundImage = fs.createWriteStream(__dirname + "/../output/termite.png");
const foregroundImage = fs.createWriteStream(__dirname + "/../output/termite2.png");

// Create the play board
const width = 1000,
      height = 1000;

const parametersBackground = {
    "steps" : 1000000,
    "num_ants" : 16,
    "simu_type" : "random",
    "overwrite" : false,
    "alpha" : .6
};

const parametersForeground = {
    "steps" : 1000000,
    "num_ants" : 20,
    "simu_type" : "rigid",
    "overwrite" : true,
    "alpha" : .7
};

console.log("Generating background...");
genTermiteArt(backgroundImage, parametersBackground);

console.log("Generating foreground...");
genTermiteArt(foregroundImage, parametersForeground);


function smoothGrid(grid, iter, radius) {
    // smoothing
    for (let smoothing_iter = 0; smoothing_iter < iter; smoothing_iter++) {
        for (let y = 0; y < height; y++) {
            for (let x = 0; x < width; x++) {
                grid.grid[y][x] = avgAntVal(grid, x, y, radius);
            }
        }
    }
}

function avgAntVal(grid, x, y, radius) {
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

function genTermiteArt(outputImageFile, parameters, reduce=false) {
    const numAnts = parameters["num_ants"];

    const grid = new Grid(width, height, numAnts);
    grid.simulate(parameters["steps"], parameters["simu_type"], parameters["overwrite"]);


    // Draw the image from the grid

    // generate the colors
    colors = [];
    colors.push(Color.random(parameters["alpha"]));

    for(let i = 0; i < numAnts - 1; i++) {
        colors.push(Color.mutation_of(colors[i]));
    }


    if (reduce) {
        for(let i = 0; i < numAnts / 3; i++) {
            //colors[randInt(numAnts)] = new Color(0, 0, 0, 0);
            colors[i] = new Color(0, 0, 0, 0);
        }

    }

    // set up the canvas
    const canvas = new Canvas(width, height);
    const context = canvas.getContext("2d");
    context.fillStyle = `rgba(${colors[0].r}, ${colors[0].g}, ${colors[0].b}, ${colors[0].a})`;
    context.fillRect(0, 0, canvas.width, canvas.height);

    smoothGrid(grid, 5, 6);

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
}