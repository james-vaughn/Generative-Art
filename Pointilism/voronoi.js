const fs = require("fs");
const canvas = require("canvas");
const {Point} = require("../lib/point");

const context = canvas.getContext("2d");
const pngStream = canvas.pngStream();
const outputImageFile = fs.createWriteStream(__dirname + "/../output/pointilism");

pngStream.on("data", function (chunk) {
    outputImageFile.write(chunk);
});

fs.readFile(_dirname + "../input/pointilism.png", function (err, input_img) {
   if(err) throw err;

   const img = new Image();
   img.src = input_img;
   context.drawImage(img, 0, 0, img.width, img.height);

   points = samplePoints(context);
   // colorGrid = voronoi(context, points);
   // drawImage(colorGrid);
});

//returns an array of points to sample
function samplePoints(context) {
    const num_random_points_per_sample = 20;
    const poisson_inner_radius = 30;
    const poisson_outer_radius = 2 * poisson_inner_radius;

    const activePoints = [Point.random(context.width, context.height)];
    const samplePoints = [];

    while(activePoints.length > 0) {
        const possible_sample_points = [];

        // TODO change this to be random selection
        // Select a point to work off of
        const samplePoint = activePoints.shift();

        // generate possible sample points
        for(let i = 0; i < num_random_points_per_sample; i++) {
            // TODO fix this so it properly calcs values withing the Poisson disc
            // TODO Also, dont let points go outside grid
            const delta_x = Math.floor(Math.random() * (poisson_outer_radius - poisson_inner_radius) + poisson_inner_radius);
            const delta_y = Math.floor(Math.random() * (poisson_outer_radius - poisson_inner_radius) + poisson_inner_radius);
            possible_sample_points.push(new Point(samplePoint.x - delta_x, samplePoint.y - delta_y));
        }

        // remove bad points
        possible_sample_points.filter(function (point) {
           return activePoints.some((activePoint) => activePoint.withinDistance(point, poisson_inner_radius));
        });

        // If no points are left, mark the point as an inactive sample point
        if(possible_sample_points.length === 0) {
            samplePoints.push(samplePoint);
            continue;
        }

        // add the remaining points to the active points list
        for(const point of possible_sample_points) {
            activePoints.push(point);
        }
    }

    return samplePoints;
}

// Returns a color matrix representing the voronoi image
function voronoi(context, points) {

}

function drawImage(colorGrid) {

}
