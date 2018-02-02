const {randInt, randSignedInt} = require("./random");

class Color{
    constructor(r, g, b, a) {
        this.r = r;
        this.g = g;
        this.b = b;
        this.a = a;
    }

    static random() {
        return new Color(Math.floor(Math.random() * 255),
            Math.floor(Math.random() * 255),
            Math.floor(Math.random() * 255),
            .3);
    }

    static mutation_of(color) {
        const new_color = new Color(color.r, color.g, color.b, color.a);

        // pick a color value to mutate
        const color_val = randInt(3);
        const amt = randSignedInt(45, 15);

        switch (color_val) {
            //then mutate that value by up to 30
            case 1:
                new_color.r = Color.adjust(new_color.r, amt);
                break;
            case 2:
                new_color.g = Color.adjust(new_color.g, amt);
                break;
            case 3:
                new_color.b = Color.adjust(new_color.b, amt);
                break;
        }

        return new_color;
    }

    static adjust(color_val, amount) {
        color_val += amount;

        if(color_val > 255) {
            color_val = 255;
        } else if(color_val < 0) {
            color_val = 0;
        }

        return color_val
    }
}

exports.Color = Color;
