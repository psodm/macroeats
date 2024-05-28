ALTER TABLE macros ADD CONSTRAINT energy_check CHECK (energy >= 0);
ALTER TABLE macros ADD CONSTRAINT calories_check CHECK (calories >= 0);
ALTER TABLE macros ADD CONSTRAINT protein_check CHECK (protein >= 0);
ALTER TABLE macros ADD CONSTRAINT carbohydrates_check CHECK (carbohydrates >= 0);
ALTER TABLE macros ADD CONSTRAINT fat_check CHECK (fat >= 0);

ALTER TABLE foods ADD CONSTRAINT quantity_check CHECK (serving_size > 0);

ALTER TABLE recipes ADD CONSTRAINT servings_check CHECK (servings >= 1);

ALTER TABLE recipe_ingredients ADD CONSTRAINT ingredient_quantity_check CHECK (quantity > 0);