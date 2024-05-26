ALTER TABLE macros DROP CONSTRAINT IF EXISTS energy_check;
ALTER TABLE macros DROP CONSTRAINT IF EXISTS calories_check;
ALTER TABLE macros DROP CONSTRAINT IF EXISTS protein_check;
ALTER TABLE macros DROP CONSTRAINT IF EXISTS carbohydrate_check;
ALTER TABLE macros DROP CONSTRAINT IF EXISTS fat_check;

ALTER TABLE ingredients DROP CONSTRAINT IF EXISTS quantity_check;

ALTER TABLE recipes DROP CONSTRAINT IF EXISTS servings_check;

ALTER TABLE recipe_ingredients DROP CONSTRAINT IF EXISTS ingredient_quantity_check;