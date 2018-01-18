class Ant {
    constructor(x, y, dir, id) {
        this.x = x;
        this.y = y;
        this.dir = dir;
        this.id = id;
    }

    canMoveForward(maxX, maxY) {
        return (this.dir === "up" && this.y > 0) ||
               (this.dir === "left" && this.x > 0) ||
               (this.dir === "right" && this.x < maxX) ||
               (this.dir === "down" && this.y < maxY);
    }

    // Returns the new coordinates
    move(maxX, maxY){
        if(this.canMoveForward(maxX, maxY) === false) {
            if(Math.random() < .5) {
                this.turn("right");
            } else {
                this.turn("left");
            }
            return this.move(maxX, maxY);
        }

        switch (this.dir) {
            case "up":
                this.y -= 1;
                break;
            case "left":
                this.x -= 1;
                break;
            case "down":
                this.y += 1;
                break;
            case "right":
                this.x += 1;
                break;
        }

        return [this.x, this.y];
    }

    turn(direction) {
        if(direction === "left") {
            switch (this.dir) {
                case "up":
                    this.dir = "left";
                    break;
                case "left":
                    this.dir = "down";
                    break;
                case "down":
                    this.dir = "right";
                    break;
                case "right":
                    this.dir = "up";
                    break;
            }
        } else if(direction == "right") {
            switch (this.dir) {
                case "up":
                    this.dir = "right";
                    break;
                case "left":
                    this.dir = "up";
                    break;
                case "down":
                    this.dir = "left";
                    break;
                case "right":
                    this.dir = "down";
                    break;
            }
        }
    }

    static randomDirection() {
        const val = Math.floor(Math.random() * 4);
        switch (val) {
            case 0:
                return "up";
            case 1:
                return "left";
            case 2:
                return "right";
            case 3:
                return "down";
        }
    }
}

exports.Ant = Ant;