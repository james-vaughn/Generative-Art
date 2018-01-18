const {Ant} = require("./ant")

class Grid {
    constructor(width, height, num_ants) {
        this.width = width;
        this.height = height;
        this.grid = [];
        for (let i = 0; i < height; i ++) {
            this.grid[i] = Array(width).fill(-1);
        }

        this.ants = [];

        // create all of the ants
        for (let i = 0; i < num_ants; i++) {
            const ant = new Ant(Math.floor(Math.random() * width),
                Math.floor(Math.random() * height),
                Ant.randomDirection(),
                i);
            this.ants.push(ant);
            this.grid[ant.y][ant.x] = ant.id;
        }
    }

    // Run the algorithm for n steps
    simulate(steps) {
        for (let i = 0; i < steps; i++) {
            for(const ant of this.ants) {

                // advance the ant
                const [x, y] = ant.move(this.width - 1, this.height - 1);

                // turn it appropriately
                if (this.grid[y][x] === ant.id) {
                    ant.turn("left");
                    this.grid[y][x] = -1;
                } else if(this.grid[y][x] === -1) {
                    this.grid[y][x] = ant.id;
                }
                else {
                    ant.turn("right");
                    // this.grid[y][x] = ant.id;
                }
            }
        }
    }
}

exports.Grid = Grid;
