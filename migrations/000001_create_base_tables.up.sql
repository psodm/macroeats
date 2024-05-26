CREATE TABLE IF NOT EXISTS measurement_units (
    measurement_unit_id SERIAL PRIMARY KEY,
    measurement_name TEXT NOT NULL UNIQUE,
    measurement_abbreviation TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS macros (
    macros_id SERIAL PRIMARY KEY,
    energy_kj NUMERIC NOT NULL,
    protein_g NUMERIC NOT NULL,
    carbohydrate_g NUMERIC NOT NULL,
    fat_g NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS ingredients (
    ingredient_id SERIAL PRIMARY KEY,
    ingredient_name TEXT NOT NULL UNIQUE,
    ingredient_serving_measurement_quantity NUMERIC NOT NULL,
    ingredient_serving_measurement_unit_id INTEGER NOT NULL REFERENCES measurement_units(measurement_unit_id),
    ingredient_macros_id INTEGER NOT NULL REFERENCES macros(macros_id)
);

CREATE TABLE IF NOT EXISTS cuisines (
    cuisine_id SERIAL PRIMARY KEY,
    cuisine_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS meal_types (
    meal_type_id SERIAL PRIMARY KEY,
    meal_type_name TEXT NOT NULL UNIQUE 
);

CREATE TABLE IF NOT EXISTS recipes (
    recipe_id SERIAL PRIMARY KEY,
    recipe_name TEXT NOT NULL UNIQUE,
    recipe_description TEXT NOT NULL,
    recipe_meal_type_id INTEGER NOT NULL REFERENCES meal_types(meal_type_id),
    recipe_cuisine_id INTEGER NOT NULL REFERENCES cuisines(cuisine_id),
    servings NUMERIC NOT NULL DEFAULT 1,
    recipe_macros_id INTEGER NOT NULL REFERENCES macros(macros_id),
    recipe_notes TEXT[] NOT NULL,
    created_at DATE NOT NULL DEFAULT CURRENT_DATE,
    version INTEGER NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS recipe_ingredient_sections (
    section_id SERIAL PRIMARY KEY,
    section_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS recipe_ingredients (
    recipe_id INTEGER NOT NULL REFERENCES recipes(recipe_id),
    ingredient_id INTEGER NOT NULL REFERENCES ingredients(ingredient_id),
    ingredient_section_id INTEGER NOT NULL REFERENCES recipe_ingredient_sections(section_id),
    quantity NUMERIC NOT NULL,
    measurement_unit_id INTEGER NOT NULL REFERENCES measurement_units(measurement_unit_id),
    PRIMARY KEY (recipe_id, ingredient_id)
);

CREATE TABLE IF NOT EXISTS recipe_instructions (
    instruction_id SERIAL PRIMARY KEY,
    recipe_id INTEGER NOT NULL REFERENCES recipes(recipe_id),
    step TEXT NOT NULL,
    instruction TEXT NOT NULL
);