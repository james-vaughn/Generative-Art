class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }

    withinDistance(point, dist) {
        return (Math.pow((this.x - point.x), 2) + Math.pow((this.y - point.y), 2)) < dist;
    }

    static random(maxX, maxY) {
        const x_pos = Math.floor(Math.random() * maxX);
        const y_pos = Math.floor(Math.random() * maxY);
        return new Point(x_pos, y_pos);
    }
}

exports.Point = Point;