const {Ant} = require("./ant");
const {Point} = require("../lib/point");

class Grid {
    constructor(width, height, num_ants) {
        this.width = width;
        this.height = height;
        this.grid = [];
        for (let i = 0; i < height; i ++) {
            this.grid[i] = Array(width).fill(-1);
        }

        this.ants = this.generateAnts_RandomApproach(width, height, num_ants, this.grid);
    }

    // Run the algorithm for n steps
    simulate(steps) {
        for (let i = 0; i < steps; i++) {
            for(const ant of this.ants) {

                // advance the ant
                const point = ant.move(this.width - 1, this.height - 1);

                // turn it appropriately
                if (this.grid[point.y][point.x] === ant.id) {
                    ant.turn("left");
                    this.grid[point.y][point.x] = -1;
                } else if(this.grid[point.y][point.x] === -1) {
                    this.grid[point.y][point.x] = ant.id;
                }
                else {
                    ant.turn("right");
                    // this.grid[y][x] = ant.id;
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

        const num_clusters = Math.sqrt(num_ants);

        // create the basis of the ant clusters
        for (let i = 0; i < num_ants; i++) {
            const ant_point = Point.random(width, height);
            const ant = new Ant(ant_point,
                Ant.randomDirection(),
                i);

            ants.push(ant);
            grid[ant.point.y][ant.point.x] = ant.id;
        }

        // create the rest of the ants close to the existing ant clusters
        for (let i = 0; i < num_ants; i++) {

        }

        return ants;
    }
}

exports.Grid = Grid;
