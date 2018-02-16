class Ant {
    constructor(point, dir, id) {
        this.point = point
        this.dir = dir;
        this.id = id;
    }

    canMoveForward(maxX, maxY) {
        return (this.dir === "up" && this.point.y > 0) ||
               (this.dir === "left" && this.point.x > 0) ||
               (this.dir === "right" && this.point.x < maxX) ||
               (this.dir === "down" && this.point.y < maxY);
    }

    // Returns the new coordinates
    move(maxX, maxY){
        if(this.canMoveForward(maxX, maxY) === false) {
            // this.turn("right");
            if(Math.random() < .5) {
                this.turn("right");
            } else {
                this.turn("left");
            }
            return this.move(maxX, maxY);
        }

        switch (this.dir) {
            case "up":
                this.point.y -= 1;
                break;
            case "left":
                this.point.x -= 1;
                break;
            case "down":
                this.point.y += 1;
                break;
            case "right":
                this.point.x += 1;
                break;
        }

        return this.point;
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