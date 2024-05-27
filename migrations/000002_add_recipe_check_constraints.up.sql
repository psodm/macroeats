ALTER TABLE macros ADD CONSTRAINT energy_check CHECK (energy > 0);
ALTER TABLE macros ADD CONSTRAINT calories_check CHECK (calories BETWEEN 4.1 * energy AND 4.2 * energy);
ALTER TABLE macros ADD CONSTRAINT protein_check CHECK (protein > 0);
ALTER TABLE macros ADD CONSTRAINT carbohydrate_check CHECK (carbohydrate > 0);
ALTER TABLE macros ADD CONSTRAINT fat_check CHECK (fat > 0);

ALTER TABLE foods ADD CONSTRAINT quantity_check CHECK (serving_quantity > 0);

ALTER TABLE recipes ADD CONSTRAINT servings_check CHECK (servings >= 1);

ALTER TABLE recipe_ingredients ADD CONSTRAINT ingredient_quantity_check CHECK (quantity > 0);