const fs = require("fs");
const Canvas = require("canvas");

function gridMidsimulationDraw(grid, filename) {
    const imageFile = fs.createWriteStream(__dirname + "/../output/" + filename);
    drawImage(grid, imageFile);
}


function drawPixel(context, x, y, color) {
    context.fillStyle = `rgba(${color.r}, ${color.g}, ${color.b}, ${color.a})`;
    context.fillRect(x, y, 1, 1);
}

function drawImage(grid, outputImageFile) {
    let colors = grid.colors;

    // set up the canvas
    const canvas = new Canvas(grid.width, grid.height);
    const context = canvas.getContext("2d");
    context.fillStyle = `rgba(${colors[0].r}, ${colors[0].g}, ${colors[0].b}, ${colors[0].a})`;
    context.fillRect(0, 0, grid.width, grid.height);

    // drawing
    for (let y = 0; y < grid.height; y++) {
        for (let x = 0; x < grid.width; x++) {
            const ant = grid.grid[y][x];

            if (ant === -1) continue;

            const color = colors[ant];
            drawPixel(context, x, y, color);
        }
    }

    //draw canvas to png file
    let pngStream = canvas.pngStream();

    pngStream.on("data", function (chunk) {
        outputImageFile.write(chunk);
    });
}

exports.gridMidsimulationDraw = gridMidsimulationDraw;
exports.drawImage = drawImage;