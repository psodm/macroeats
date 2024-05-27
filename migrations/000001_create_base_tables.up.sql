CREATE TABLE IF NOT EXISTS measurements (
    measurement_id SERIAL PRIMARY KEY,
    measurement_name TEXT NOT NULL UNIQUE,
    measurement_abbreviation TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS macros (
    macros_id SERIAL PRIMARY KEY,
    energy NUMERIC NOT NULL,
    calories NUMERIC NOT NULL,
    protein NUMERIC NOT NULL,
    carbohydrate NUMERIC NOT NULL,
    fat NUMERIC NOT NULL
);

CREATE TABLE IF NOT EXISTS foods (
    food_id SERIAL PRIMARY KEY,
    food_name TEXT NOT NULL UNIQUE,
    serving_quantity NUMERIC NOT NULL,
    serving_measurement_id INTEGER NOT NULL REFERENCES measurements(measurement_id),
    macros_id INTEGER NOT NULL REFERENCES macros(macros_id)
);

CREATE TABLE IF NOT EXISTS cuisines (
    cuisine_id SERIAL PRIMARY KEY,
    cuisine_name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS meals (
    meal_id SERIAL PRIMARY KEY,
    meal_name TEXT NOT NULL UNIQUE 
);

CREATE TABLE IF NOT EXISTS recipes (
    recipe_id SERIAL PRIMARY KEY,
    recipe_name TEXT NOT NULL UNIQUE,
    recipe_description TEXT NOT NULL,
    recipe_meal_type_id INTEGER NOT NULL REFERENCES meals(meal_id),
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
    food_id INTEGER NOT NULL REFERENCES foods(food_id),
    ingredient_section_id INTEGER NOT NULL REFERENCES recipe_ingredient_sections(section_id),
    quantity NUMERIC NOT NULL,
    measurement_id INTEGER NOT NULL REFERENCES measurements(measurement_id),
    PRIMARY KEY (recipe_id, food_id)
);

CREATE TABLE IF NOT EXISTS recipe_instructions (
    instruction_id SERIAL PRIMARY KEY,
    recipe_id INTEGER NOT NULL REFERENCES recipes(recipe_id),
    step TEXT NOT NULL,
    instruction TEXT NOT NULL
);