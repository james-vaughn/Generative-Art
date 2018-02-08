const {Ant} = require("./ant");
const {Point} = require("../lib/point");
const {randSignedInt} = require("../lib/random");
const {gridMidsimulationDraw} = require("./drawing");

class Grid {
    constructor(width, height, num_ants, colors, filePrefix) {
        this.width = width;
        this.height = height;
        this.num_ants = num_ants;
        this.colors = colors;
		this.filePrefix = filePrefix;
        this.grid = [];
        for (let i = 0; i < height; i ++) {
            this.grid[i] = Array(width).fill(-1);
        }

        this.ants = this.generateAnts_ClusterApproach(width, height, num_ants, this.grid);
    }

    // Run the algorithm for steps number of steps
    // The simulation can either be of "random" type, meaning the termites will random walk
    // or of "rigid" type, meaning the termites will always follow a given pattern with little randomness
    // Overwrite is a boolean indicating if termites can change other termites' colors
    simulate(steps, type, overwrite=false) {
        if(type !== "random" && type !== "rigid") {
            throw "Bad simulation type";
        }
		
		let captureAt = 100;
        for (let i = 0; i < steps; i++) {
            if (i !== 0 && i === captureAt) {
				console.log(`\tCreating mid-simulation image for step ${i}`);
                gridMidsimulationDraw(this, `${this.filePrefix}step${i}.png`);
				captureAt *= 5;
            }

            for(const ant of this.ants) {

                // advance the ant
                const point = ant.move(this.width - 1, this.height - 1);

                // turn it appropriately
                if (this.grid[point.y][point.x] === ant.id) {

                    if(type === "random") {
                        if(Math.random() < .5) {
                            ant.turn("right");
                        } else {
                            ant.turn("left");
                        }
                    } else if(type === "rigid") {
                     ant.turn("left");
                    }

                } else if(this.grid[point.y][point.x] === -1) {
                    this.grid[point.y][point.x] = ant.id;
                }
                else {

                    if(type === "random") {
                        if(Math.random() < .5) {
                            ant.turn("right");
                        } else {
                            ant.turn("left");
                        }
                    } else if(type === "rigid") {
                         ant.turn("right");
                    }

                    if(overwrite === true) {
                                this.grid[point.y][point.x] = ant.id;
        		    }

                }
            }
        }
    }

    generateAnts_RandomApproach(width, height, num_ants, grid) {
        const ants = [];

        // create all of the ants
        for (let i = 0; i < num_ants; i++) {
            const ant_point = Point.random(width, height);
            const ant = new Ant(ant_point,
                                Ant.randomDirection(),
                                i);
            ants.push(ant);
            grid[ant.point.y][ant.point.x] = ant.id;
        }

        return ants;
    }

    generateAnts_ClusterApproach(width, height, num_ants, grid) {
        const ants = [];

        const num_clusters = Math.floor(Math.sqrt(num_ants));

        // create the basis of the ant clusters
        for (let i = 0; i < num_clusters; i++) {
            const ant_point = Point.random(width, height);
            const ant = new Ant(ant_point,
                Ant.randomDirection(),
                i);

            ants.push(ant);
            grid[ant.point.y][ant.point.x] = ant.id;
        }

        // create the rest of the ants close to the existing ant clusters
        for (let i = num_clusters; i < num_ants; i++) {
            const base_ant = ants[Math.floor(Math.random() * num_clusters)];
            const new_point = Point.adjust(base_ant.point, randSignedInt(25) + 5, randSignedInt(25) + 5, width - 1, height - 1);
            const ant = new Ant(new_point,
                Ant.randomDirection(),
                i);

            ants.push(ant);
            grid[ant.point.y][ant.point.x] = ant.id;
        }

        return ants;
    }
}

exports.Grid = Grid;
