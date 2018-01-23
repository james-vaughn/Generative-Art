class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }

    withinDistance(point, dist) {
        const v = Math.sqrt(Math.pow((this.x - point.x), 2) + Math.pow((this.y - point.y), 2)) < dist;
        return v;
    }

    static random(maxX, maxY) {
        const x_pos = Math.floor(Math.random() * maxX);
        const y_pos = Math.floor(Math.random() * maxY);
        return new Point(x_pos, y_pos);
    }

    static adjust(point, x_delta, y_delta, x_upper, y_upper) {
        let x_new = point.x + x_delta;
        if (x_new > x_upper) {
            x_new = x_upper;
        } else if (x_new < 0) {
            x_new = 0;
        }

        let y_new = point.y + y_delta;
        if (y_new > y_upper) {
            y_new = y_upper;
        } else if (y_new < 0) {
            y_new = 0;
        }

        return new Point(x_new, y_new);

    }
}

exports.Point = Point;