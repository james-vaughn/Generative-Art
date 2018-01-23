const fs = require("fs");
const Canvas = require("canvas");
const {Point} = require("../lib/point");

const Image = Canvas.Image;

const outputImageFile = fs.createWriteStream(__dirname + "/../output/pointilism.png");

fs.readFile(__dirname + "/../input/pointilism.png", function (err, input_img) {
   if(err) throw err;

   const img = new Image();
   img.src = input_img;

   const canvas = new Canvas(img.width, img.height);
   const context = canvas.getContext("2d");

   context.drawImage(img, 0, 0, img.width, img.height);

   points = samplePoints(img.width, img.height);
   console.log(points);
   // colorGrid = voronoi(context, points);
   // drawImage(colorGrid);

   const pngStream = canvas.pngStream();
   pngStream.on("data", function (chunk) {
       outputImageFile.write(chunk);
   });
});

//returns an array of points to sample
function samplePoints(width, height) {
    const num_random_points_per_sample = 20;
    const poisson_inner_radius = 50;
    const poisson_outer_radius = 2 * poisson_inner_radius;

    const activePoints = [Point.random(width, height)];
    const samplePoints = [];

    while(activePoints.length > 0) {
        const possible_sample_points = [];

        // TODO change this to be random selection
        // Select a point to work off of
        const samplePoint = activePoints.shift();

        // generate possible sample points
        for(let i = 0; i < num_random_points_per_sample; i++) {
            // TODO fix this so it properly calcs values withing the Poisson disc
            const delta_x = Math.floor(2 * (Math.random() - .5) * (poisson_outer_radius - poisson_inner_radius) + poisson_inner_radius);
            const delta_y = Math.floor(2 * (Math.random() - .5) * (poisson_outer_radius - poisson_inner_radius) + poisson_inner_radius);
            const pt = new Point(samplePoint.x + delta_x, samplePoint.y + delta_y);

            if(pt.x >= 0 && pt.x < width && pt.y >= 0 && pt.y < height) {
                possible_sample_points.push(pt);
            }
        }

        // remove bad points
        const points = possible_sample_points.filter(function (point) {
           return activePoints.some((activePoint) => activePoint.withinDistance(point, poisson_inner_radius)) === false &&
                  samplePoints.some((samplePoint) => samplePoint.withinDistance(point, poisson_inner_radius)) === false;
        });

        // If no points are left, mark the point as an inactive sample point
        if(points.length === 0) {
            samplePoints.push(samplePoint);
            continue;
        }

        // add one of the possible points to the points list
        activePoints.push(points[0]);
        activePoints.push(samplePoint);
    }

    return samplePoints;
}

// Returns a color matrix representing the voronoi image
function voronoi(context, points) {

}

function drawImage(colorGrid) {

}
