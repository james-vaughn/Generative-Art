//modules
const fs = require("fs");
const {Grid} = require("./grid");
const {Color} = require("../lib/color");
const {drawImage} = require("./drawing");

const backgroundImage = fs.createWriteStream(__dirname + "/../output/termite.png");
const foregroundImage = fs.createWriteStream(__dirname + "/../output/termite2.png");

// Create the play board
const width = 1920,
      height = 1080;

const parametersBackground = {
    "steps" : 2000000,
    "num_ants" : 16,
    "simu_type" : "random",
    "overwrite" : false,
    "alpha" : .4
};

const parametersForeground = {
    "steps" : 2000000,
    "num_ants" : 16,
    "simu_type" : "random",
    "overwrite" : true,
    "alpha" : .6
};

console.log("Generating background...");
genTermiteArt(backgroundImage, parametersBackground, __dirname + "/../output/bg_");

console.log("Generating foreground...");
genTermiteArt(foregroundImage, parametersForeground, __dirname + "/../output/fg_", true);


function smoothGrid(grid, iter, radius, outputImgPrefix) {
    // smoothing
    for (let smoothing_iter = 0; smoothing_iter < iter; smoothing_iter++) {
        //generate image for every other smoothing iteration
        if(smoothing_iter % 2 === 0) {
            console.log(`\tGenerating smoothing image ${smoothing_iter}`);

            const imageFile = fs.createWriteStream(`${outputImgPrefix}smootingIter${smoothing_iter}.png`);
            drawImage(grid, imageFile);
        }

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

function genTermiteArt(outputImageFile, parameters, smoothingPrefix, reduce=false) {
    const numAnts = parameters["num_ants"];
    const alpha = parameters["alpha"];

    // generate the colors
    let colors = [];
    colors.push(Color.random(alpha));

    for(let i = 0; i < numAnts - 1; i++) {
        colors.push(Color.mutation_of(colors[i]));
    }


    if (reduce) {
        for(let i = 0; i < numAnts / 3; i++) {
            //colors[randInt(numAnts)] = new Color(0, 0, 0, 0);
            colors[i] = new Color(0, 0, 0, 0);
        }

    }

    //temporarily make the colors opaque
    colors.map(color => {
        if(color.a > 0) {
            color.a = 1
        }
    });

    const grid = new Grid(width, height, numAnts, colors);
    grid.simulate(parameters["steps"], parameters["simu_type"], parameters["overwrite"]);


    smoothGrid(grid, 10, 6, smoothingPrefix);

    //return the colors back to normal alpha level
    grid.colors.map(color => {
        if(color.a > 0) {
            color.a = alpha
        }
    });

    drawImage(grid, outputImageFile, colors);
}