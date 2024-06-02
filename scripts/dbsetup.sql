INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Milligram', 'mg');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Ounce', 'oz');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Gram', 'g');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Pound', 'lb');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Kilogram', 'kg');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Pinch', 'pinch');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Litre', 'l');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Fluid Ounce ', 'fl oz');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Gallon', 'gal');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Pint', 'pt');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Quart', 'qt');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Millilitre', 'ml');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Drop', 'drop');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Cup', 'cup');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Teaspoon', 'tsp');
INSERT INTO measurements (measurement_name, measurement_abbreviation) VALUES ('Tablespoon', 'tbsp');

INSERT INTO cuisines (cuisine_name) VALUES ('American');
INSERT INTO cuisines (cuisine_name) VALUES ('Australian');
INSERT INTO cuisines (cuisine_name) VALUES ('British');
INSERT INTO cuisines (cuisine_name) VALUES ('Dutch');
INSERT INTO cuisines (cuisine_name) VALUES ('Chinese');
INSERT INTO cuisines (cuisine_name) VALUES ('Fast Food');
INSERT INTO cuisines (cuisine_name) VALUES ('French');
INSERT INTO cuisines (cuisine_name) VALUES ('Fusion');
INSERT INTO cuisines (cuisine_name) VALUES ('German');
INSERT INTO cuisines (cuisine_name) VALUES ('Greek');
INSERT INTO cuisines (cuisine_name) VALUES ('Indonesian');
INSERT INTO cuisines (cuisine_name) VALUES ('Irish');
INSERT INTO cuisines (cuisine_name) VALUES ('Italian');
INSERT INTO cuisines (cuisine_name) VALUES ('Japanese');
INSERT INTO cuisines (cuisine_name) VALUES ('Korean');
INSERT INTO cuisines (cuisine_name) VALUES ('Mexican');
INSERT INTO cuisines (cuisine_name) VALUES ('Russian');
INSERT INTO cuisines (cuisine_name) VALUES ('Spanish');
INSERT INTO cuisines (cuisine_name) VALUES ('Street');
INSERT INTO cuisines (cuisine_name) VALUES ('Thai');
INSERT INTO cuisines (cuisine_name) VALUES ('Vegan');

INSERT INTO meals (meal_name) VALUES ('Breakfast');
INSERT INTO meals (meal_name) VALUES ('Lunch');
INSERT INTO meals (meal_name) VALUES ('Dinner');
INSERT INTO meals (meal_name) VALUES ('Dessert');
INSERT INTO meals (meal_name) VALUES ('Snack');
INSERT INTO meals (meal_name) VALUES ('Drink');

INSERT INTO brands (brand_name) VALUES ('');
INSERT INTO brands (brand_name) VALUES ('Tamar Valley');
INSERT INTO brands (brand_name) VALUES ('Jalna');
INSERT INTO brands (brand_name) VALUES ('Aeroplane');
INSERT INTO brands (brand_name) VALUES ('Bulk Nutrients');
INSERT INTO brands (brand_name) VALUES ('EHP Labs');
INSERT INTO brands (brand_name) VALUES ('Old El Paso');
INSERT INTO brands (brand_name) VALUES ('Leggo''s');

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (0, 0, 0, 0, 0)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Water', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'ml'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (217.6, 52, 0.2, 13, 0.3)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Apple', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (656, 156, 1.5, 0, 16)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Avocado', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (196.6, 47, 0.9, 12, 0.1)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Orange', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (615, 147, 20, 3, 6.3)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Steak, Rump', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (150, 35.9, 3.6, 5, 0.1)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Milk, Skim', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'ml'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (217.6, 52, 11, 0.7, 0.2)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Egg White', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (0, 0, 0, 0, 0)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Salt', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (0, 0, 0, 0, 0)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, serving_size, serving_measurement_id, macros_id)
    VALUES ('Pepper', 100, (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (231, 55, 6.6, 6.7, 0.2)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Greek Yoghurt, 99.85% Fat Free',
            (SELECT brand_id FROM brands WHERE brand_name = 'Tamar Valley'),
            100,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (31, 5, 0, 1, 0)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Pride Pre-Workout',
            (SELECT brand_id FROM brands WHERE brand_name = 'EHP Labs'),
            9.35,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (25, 6, 0.5, 0.5, 0.5)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Jelly, Lite',
            (SELECT brand_id FROM brands WHERE brand_name = 'Aeroplane'),
            100,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (1636, 391, 74.8, 10.2, 5.4)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Whey Protein Concentrate, Raw',
            (SELECT brand_id FROM brands WHERE brand_name = 'Bulk Nutrients'),
            100,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (1160, 277, 7.5, 50.3, 3.3)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Taco Spice Mix',
            (SELECT brand_id FROM brands WHERE brand_name = 'Old El Paso'),
            100,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (690, 165, 31, 0, 3.6)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Chicken, Breast (Skinless)',
            (SELECT brand_id FROM brands WHERE brand_name = ''),
            100,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;

BEGIN;
    WITH m AS (
        INSERT INTO macros (energy, calories, protein, carbohydrates, fat)
        VALUES (296, 70.7, 3.3, 11.3, 0.5)
        RETURNING macros_id
    )
    INSERT INTO foods (food_name, brand_id, serving_size, serving_measurement_id, macros_id)
    VALUES ('Tomato Paste',
            (SELECT brand_id FROM brands WHERE brand_name = 'Leggo''s'),
            100,
            (SELECT measurement_id FROM measurements WHERE measurement_abbreviation = 'g'), 
            lastval());
COMMIT;


