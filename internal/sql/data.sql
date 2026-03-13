-- Grocery Categories
INSERT INTO grocery_categories (name) VALUES
('Produce'),
('Meat & Seafood'),
('Dairy & Eggs'),
('Pasta & Grains'),
('Canned & Jarred'),
('Baking & Spices'),
('Condiments & Sauces'),
('Frozen'),
('Bread & Bakery');

-- Ingredients (Produce - category_id 1)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Garlic', 1),
('Onion', 1),
('Tomato', 1),
('Basil', 1),
('Parsley', 1),
('Lemon', 1),
('Lime', 1),
('Cilantro', 1),
('Bell Pepper', 1),
('Mushrooms', 1),
('Spinach', 1),
('Broccoli', 1),
('Carrots', 1),
('Celery', 1),
('Potatoes', 1),
('Sweet Potatoes', 1),
('Zucchini', 1),
('Cucumber', 1),
('Lettuce', 1),
('Romaine Lettuce', 1),
('Avocado', 1),
('Jalapeno', 1),
('Ginger', 1),
('Green Onions', 1),
('Cauliflower', 1),
('Asparagus', 1),
('Cherry Tomatoes', 1),
('Red Onion', 1);

-- Ingredients (Meat & Seafood - category_id 2)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Ground Beef', 2),
('Beef Steak', 2),
('Chicken Breast', 2),
('Chicken Thighs', 2),
('Pork Shoulder', 2),
('Pork Chops', 2),
('Bacon', 2),
('Ground Pork', 2),
('Shrimp', 2),
('Carne Asada', 2),
('Ground Turkey', 2);

-- Ingredients (Dairy & Eggs - category_id 3)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Parmesan Cheese', 3),
('Mozzarella Cheese', 3),
('Cheddar Cheese', 3),
('Heavy Cream', 3),
('Butter', 3),
('Milk', 3),
('Eggs', 3),
('Ricotta Cheese', 3),
('Cream Cheese', 3),
('Sour Cream', 3),
('Feta Cheese', 3);

-- Ingredients (Pasta & Grains - category_id 4)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Spaghetti', 4),
('Fettuccine', 4),
('Penne', 4),
('Lasagna Noodles', 4),
('White Rice', 4),
('Jasmine Rice', 4),
('Bread Crumbs', 4);

-- Ingredients (Canned & Jarred - category_id 5)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Marinara Sauce', 5),
('Diced Tomatoes', 5),
('Tomato Paste', 5),
('Tomato Sauce', 5),
('Chicken Broth', 5),
('Beef Broth', 5),
('Coconut Milk', 5),
('Black Beans', 5),
('Refried Beans', 5),
('Corn', 5),
('Olives', 5);

-- Ingredients (Baking & Spices - category_id 6)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Flour', 6),
('Sugar', 6),
('Salt', 6),
('Black Pepper', 6),
('Italian Seasoning', 6),
('Red Pepper Flakes', 6),
('Cumin', 6),
('Chili Powder', 6),
('Paprika', 6),
('Garlic Powder', 6),
('Onion Powder', 6),
('Oregano', 6),
('Thyme', 6),
('Bay Leaves', 6),
('Curry Powder', 6),
('Turmeric', 6),
('Baking Powder', 6),
('Cornstarch', 6);

-- Ingredients (Condiments & Sauces - category_id 7)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Olive Oil', 7),
('Vegetable Oil', 7),
('Soy Sauce', 7),
('Worcestershire Sauce', 7),
('Hot Sauce', 7),
('Buffalo Sauce', 7),
('Ranch Dressing', 7),
('Caesar Dressing', 7),
('Alfredo Sauce', 7),
('Teriyaki Sauce', 7),
('BBQ Sauce', 7),
('Ketchup', 7),
('Mustard', 7),
('Mayonnaise', 7);

-- Ingredients (Frozen - category_id 8)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Frozen Peas', 8),
('Frozen Peas and Carrots', 8),
('Frozen French Fries', 8);

-- Ingredients (Bread & Bakery - category_id 9)
INSERT INTO ingredients (name, grocery_category_id) VALUES
('Hamburger Buns', 9),
('Slider Buns', 9),
('Sandwich Bread', 9),
('Tortillas', 9),
('Flour Tortillas', 9),
('Corn Tortillas', 9),
('Pizza Dough', 9);

-- Meals
INSERT INTO meals (name, description) VALUES
('Spaghetti', 'Classic spaghetti with marinara and ground beef'),
('Steak and Potatoes', 'Grilled steak served with roasted potatoes'),
('Steak and Frites with Mushroom Sauce', 'Steak with crispy fries and creamy mushroom sauce'),
('Grilled Chicken', 'Simple grilled chicken breast with herbs'),
('Chicken Fettuccine Alfredo', 'Creamy alfredo pasta with grilled chicken'),
('Chicken Parmesan', 'Breaded chicken with marinara and melted mozzarella'),
('Pulled Pork', 'Slow-cooked pulled pork with BBQ sauce'),
('Pork Chops', 'Pan-seared pork chops with herbs'),
('Orange Chicken and Fried Rice', 'Crispy orange chicken with savory fried rice'),
('Carne Asada Tacos', 'Grilled steak tacos with fresh toppings'),
('Burritos', 'Flour tortillas filled with rice, beans, and meat'),
('Fajitas', 'Sizzling grilled meat with peppers and onions'),
('Enchiladas', 'Corn tortillas filled with chicken and topped with sauce'),
('California Chicken Sandwiches', 'Grilled chicken sandwiches with avocado and bacon'),
('Homemade Mac and Cheese', 'Creamy baked macaroni and cheese'),
('Lasagna', 'Layered pasta with meat sauce and cheese'),
('Chicken Cacciatore', 'Italian hunter-style chicken with tomatoes and peppers'),
('Beef Stew', 'Hearty beef stew with vegetables'),
('Homemade Chicken Noodle Soup', 'Classic chicken soup with egg noodles'),
('Meatloaf', 'Classic meatloaf with ketchup glaze'),
('Yellow Curry with Rice', 'Thai yellow curry with chicken and jasmine rice'),
('Chicken Teriyaki', 'Grilled chicken glazed with teriyaki sauce'),
('Hamburger Sliders', 'Mini beef burgers on slider buns'),
('Breakfast For Dinner', 'Classic breakfast items for dinner'),
('Sheet-Pan Chicken and Veggies', 'Roasted chicken with seasonal vegetables'),
('Homemade Pizza', 'Hand-tossed pizza with custom toppings'),
('Stuffed Peppers', 'Bell peppers stuffed with rice and meat'),
('Stir Fry', 'Quick-cooked vegetables and meat with sauce'),
('Shrimp Pasta', 'Pasta with shrimp in garlic butter sauce'),
('Baked Potatoes', 'Loaded baked potatoes with toppings'),
('Buffalo Chicken', 'Spicy buffalo chicken with dipping sauce');

-- Meal Ingredients (using meal_ids 1-31 and ingredient_ids from above)
-- Spaghetti (meal_id 1)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(1, 18, 1, 'lb'),
(1, 33, 24, 'oz'),
(1, 45, 2, 'cloves'),
(1, 61, 1, 'tbsp'),
(1, 17, 1, 'tsp'),
(1, 55, 1, 'cup'),
(1, 41, 0.5, 'tsp');

-- Steak and Potatoes (meal_id 2)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(2, 30, 2, 'steaks'),
(2, 46, 2, 'lb'),
(2, 45, 3, 'cloves'),
(2, 61, 2, 'tbsp'),
(2, 66, 1, 'tsp'),
(2, 67, 0.5, 'tsp'),
(2, 41, 1, 'tsp');

-- Steak and Frites with Mushroom Sauce (meal_id 3)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(3, 30, 2, 'steaks'),
(3, 76, 1, 'lb'),
(3, 36, 8, 'oz'),
(3, 45, 2, 'cloves'),
(3, 61, 1, 'tbsp'),
(3, 63, 0.5, 'cup'),
(3, 55, 0.25, 'cup'),
(3, 41, 0.5, 'tsp'),
(3, 67, 0.25, 'tsp');

-- Grilled Chicken (meal_id 4)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(4, 31, 4, 'breasts'),
(4, 45, 3, 'cloves'),
(4, 61, 2, 'tbsp'),
(4, 66, 1, 'tsp'),
(4, 67, 1, 'tsp'),
(4, 73, 1, 'tsp'),
(4, 41, 1, 'tsp');

-- Chicken Fettuccine Alfredo (meal_id 5)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(5, 31, 2, 'breasts'),
(5, 56, 12, 'oz'),
(5, 55, 1, 'cup'),
(5, 54, 0.5, 'cup'),
(5, 45, 3, 'cloves'),
(5, 53, 0.5, 'cup'),
(5, 61, 2, 'tbsp'),
(5, 41, 0.5, 'tsp'),
(5, 67, 0.25, 'tsp');

-- Chicken Parmesan (meal_id 6)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(6, 31, 4, 'breasts'),
(6, 65, 1, 'cup'),
(6, 52, 1, 'cup'),
(6, 33, 16, 'oz'),
(6, 45, 2, 'cloves'),
(6, 61, 0.25, 'cup'),
(6, 76, 2, 'eggs'),
(6, 62, 0.5, 'cup'),
(6, 73, 1, 'tsp'),
(6, 41, 1, 'tsp');

-- Pulled Pork (meal_id 7)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(7, 34, 3, 'lb'),
(7, 84, 1, 'cup'),
(7, 45, 4, 'cloves'),
(7, 66, 1, 'tbsp'),
(7, 67, 1, 'tsp'),
(7, 69, 1, 'tsp'),
(7, 70, 1, 'tsp'),
(7, 87, 0.25, 'cup'),
(7, 41, 2, 'tsp');

-- Pork Chops (meal_id 8)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(8, 35, 4, 'chops'),
(8, 45, 2, 'cloves'),
(8, 61, 2, 'tbsp'),
(8, 66, 1, 'tsp'),
(8, 73, 1, 'tsp'),
(8, 41, 1, 'tsp'),
(8, 67, 0.5, 'tsp');

-- Orange Chicken and Fried Rice (meal_id 9)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(9, 31, 2, 'breasts'),
(9, 59, 2, 'cups'),
(9, 45, 3, 'cloves'),
(9, 48, 3, 'green onions'),
(9, 78, 2, 'tbsp'),
(9, 81, 2, 'tbsp'),
(9, 89, 0.5, 'cup'),
(9, 62, 0.25, 'cup'),
(9, 91, 0.5, 'cup'),
(9, 67, 0.5, 'tsp');

-- Carne Asada Tacos (meal_id 10)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(10, 39, 1.5, 'lb'),
(10, 97, 8, 'tortillas'),
(10, 25, 0.5, 'cup'),
(10, 16, 2, 'limes'),
(10, 15, 0.5, 'cup'),
(10, 47, 1, 'onion'),
(10, 45, 3, 'cloves'),
(10, 61, 2, 'tbsp'),
(10, 69, 1, 'tsp'),
(10, 41, 1, 'tsp');

-- Burritos (meal_id 11)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(11, 28, 1, 'lb'),
(11, 98, 4, 'tortillas'),
(11, 59, 2, 'cups'),
(11, 64, 1, 'can'),
(11, 52, 1, 'cup'),
(11, 25, 0.5, 'cup'),
(11, 16, 1, 'lime'),
(11, 47, 0.5, 'onion'),
(11, 45, 2, 'cloves');

-- Fajitas (meal_id 12)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(12, 30, 1.5, 'lb'),
(12, 36, 2, 'bell peppers'),
(12, 47, 2, 'onions'),
(12, 45, 3, 'cloves'),
(12, 98, 8, 'tortillas'),
(12, 57, 0.5, 'cup'),
(12, 61, 2, 'tbsp'),
(12, 69, 1, 'tsp'),
(12, 70, 1, 'tsp'),
(12, 16, 2, 'limes');

-- Enchiladas (meal_id 13)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(13, 31, 3, 'cups'),
(13, 99, 10, 'tortillas'),
(13, 65, 2, 'cups'),
(13, 52, 1, 'cup'),
(13, 64, 1, 'can'),
(13, 58, 16, 'oz'),
(13, 45, 2, 'cloves'),
(13, 69, 1, 'tsp'),
(13, 70, 1, 'tsp'),
(13, 61, 1, 'tbsp');

-- California Chicken Sandwiches (meal_id 14)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(14, 31, 4, 'breasts'),
(14, 94, 4, 'buns'),
(14, 20, 2, 'avocados'),
(14, 37, 8, 'slices'),
(14, 57, 0.25, 'cup'),
(14, 91, 0.25, 'cup'),
(14, 25, 0.5, 'cup'),
(14, 61, 1, 'tbsp'),
(14, 16, 1, 'lime');

-- Homemade Mac and Cheese (meal_id 15)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(15, 57, 16, 'oz'),
(15, 54, 4, 'tbsp'),
(15, 62, 0.25, 'cup'),
(15, 51, 2, 'cups'),
(15, 50, 1, 'cup'),
(15, 53, 1.5, 'cups'),
(15, 41, 1, 'tsp'),
(15, 67, 0.5, 'tsp'),
(15, 73, 0.5, 'tsp');

-- Lasagna (meal_id 16)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(16, 28, 1, 'lb'),
(16, 60, 12, 'sheets'),
(16, 52, 2, 'cups'),
(16, 50, 2, 'cups'),
(16, 55, 0.5, 'cup'),
(16, 33, 32, 'oz'),
(16, 47, 1, 'onion'),
(16, 45, 3, 'cloves'),
(16, 76, 1, 'egg'),
(16, 61, 2, 'tbsp'),
(16, 73, 1, 'tsp'),
(16, 41, 1, 'tsp');

-- Chicken Cacciatore (meal_id 17)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(17, 32, 4, 'thighs'),
(17, 47, 1, 'onion'),
(17, 36, 2, 'bell peppers'),
(17, 35, 8, 'oz'),
(17, 45, 4, 'cloves'),
(17, 58, 28, 'oz'),
(17, 61, 2, 'tbsp'),
(17, 74, 1, 'tsp'),
(17, 75, 1, 'tsp'),
(17, 41, 1, 'tsp'),
(17, 67, 0.5, 'tsp');

-- Beef Stew (meal_id 18)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(18, 29, 2, 'lb'),
(18, 46, 1.5, 'lb'),
(18, 44, 3, 'carrots'),
(18, 47, 2, 'onions'),
(18, 45, 3, 'cloves'),
(18, 63, 4, 'cups'),
(18, 61, 2, 'tbsp'),
(18, 74, 1, 'tsp'),
(18, 75, 1, 'tsp'),
(18, 76, 2, 'leaves'),
(18, 41, 1.5, 'tsp'),
(18, 67, 0.5, 'tsp');

-- Homemade Chicken Noodle Soup (meal_id 19)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(19, 32, 2, 'lb'),
(19, 44, 3, 'carrots'),
(19, 43, 3, 'stalks'),
(19, 47, 1, 'onion'),
(19, 45, 3, 'cloves'),
(19, 63, 8, 'cups'),
(19, 56, 8, 'oz'),
(19, 74, 1, 'tsp'),
(19, 75, 0.5, 'tsp'),
(19, 41, 1, 'tsp'),
(19, 67, 0.25, 'tsp');

-- Meatloaf (meal_id 20)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(20, 28, 2, 'lb'),
(20, 76, 1, 'egg'),
(20, 65, 0.5, 'cup'),
(20, 47, 1, 'onion'),
(20, 45, 2, 'cloves'),
(20, 61, 1, 'tbsp'),
(20, 88, 0.5, 'cup'),
(20, 87, 2, 'tbsp'),
(20, 73, 1, 'tsp'),
(20, 70, 1, 'tsp'),
(20, 41, 1.5, 'tsp'),
(20, 67, 0.5, 'tsp');

-- Yellow Curry with Rice (meal_id 21)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(21, 31, 1.5, 'lb'),
(21, 59, 2, 'cups'),
(21, 64, 14, 'oz'),
(21, 36, 1, 'bell pepper'),
(21, 45, 3, 'cloves'),
(21, 49, 1, 'tbsp'),
(21, 77, 2, 'tbsp'),
(21, 61, 1, 'tbsp'),
(21, 41, 1, 'tsp');

-- Chicken Teriyaki (meal_id 22)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(22, 31, 4, 'breasts'),
(22, 59, 2, 'cups'),
(22, 82, 0.5, 'cup'),
(22, 45, 2, 'cloves'),
(22, 48, 2, 'green onions'),
(22, 61, 1, 'tbsp'),
(22, 89, 1, 'tbsp'),
(22, 41, 0.5, 'tsp');

-- Hamburger Sliders (meal_id 23)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(23, 28, 1.5, 'lb'),
(23, 96, 12, 'buns'),
(23, 52, 6, 'slices'),
(23, 47, 1, 'onion'),
(23, 45, 2, 'cloves'),
(23, 88, 2, 'tbsp'),
(23, 87, 1, 'tbsp'),
(23, 73, 1, 'tsp'),
(23, 41, 1, 'tsp'),
(23, 67, 0.5, 'tsp');

-- Breakfast For Dinner (meal_id 24)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(24, 76, 6, 'eggs'),
(24, 37, 8, 'slices'),
(24, 95, 4, 'slices'),
(24, 61, 2, 'tbsp'),
(24, 46, 1, 'lb'),
(24, 45, 2, 'cloves'),
(24, 41, 0.5, 'tsp'),
(24, 67, 0.25, 'tsp');

-- Sheet-Pan Chicken and Veggies (meal_id 25)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(25, 32, 4, 'thighs'),
(25, 46, 1.5, 'lb'),
(25, 42, 2, 'cups'),
(25, 44, 2, 'carrots'),
(25, 45, 4, 'cloves'),
(25, 61, 3, 'tbsp'),
(25, 74, 1, 'tsp'),
(25, 41, 1, 'tsp'),
(25, 67, 0.5, 'tsp');

-- Homemade Pizza (meal_id 26)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(26, 100, 1, 'lb'),
(26, 33, 16, 'oz'),
(26, 52, 2, 'cups'),
(26, 38, 0.5, 'cup'),
(26, 37, 8, 'slices'),
(26, 36, 1, 'bell pepper'),
(26, 45, 2, 'cloves'),
(26, 61, 1, 'tbsp'),
(26, 73, 1, 'tsp'),
(26, 41, 0.5, 'tsp');

-- Stuffed Peppers (meal_id 27)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(27, 36, 6, 'bell peppers'),
(27, 28, 1, 'lb'),
(27, 59, 1, 'cup'),
(27, 33, 15, 'oz'),
(27, 47, 1, 'onion'),
(27, 45, 2, 'cloves'),
(27, 52, 1, 'cup'),
(27, 61, 1, 'tbsp'),
(27, 73, 1, 'tsp'),
(27, 41, 1, 'tsp');

-- Stir Fry (meal_id 28)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(28, 31, 1.5, 'lb'),
(28, 59, 2, 'cups'),
(28, 42, 2, 'cups'),
(28, 36, 2, 'bell peppers'),
(28, 45, 3, 'cloves'),
(28, 49, 1, 'tbsp'),
(28, 78, 3, 'tbsp'),
(28, 61, 2, 'tbsp'),
(28, 77, 1, 'tsp'),
(28, 41, 0.5, 'tsp');

-- Shrimp Pasta (meal_id 29)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(29, 40, 1, 'lb'),
(29, 56, 12, 'oz'),
(29, 45, 4, 'cloves'),
(29, 61, 4, 'tbsp'),
(29, 55, 0.25, 'cup'),
(29, 16, 1, 'lemon'),
(29, 51, 0.5, 'cup'),
(29, 41, 0.5, 'tsp'),
(29, 67, 0.25, 'tsp');

-- Baked Potatoes (meal_id 30)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(30, 46, 4, 'large potatoes'),
(30, 54, 4, 'tbsp'),
(30, 51, 1, 'cup'),
(30, 37, 8, 'slices'),
(30, 57, 1, 'cup'),
(30, 45, 2, 'cloves'),
(30, 25, 0.5, 'cup'),
(30, 41, 1, 'tsp');

-- Buffalo Chicken (meal_id 31)
INSERT INTO meal_ingredients (meal_id, ingredient_id, quantity, unit) VALUES
(31, 31, 2, 'lb'),
(31, 83, 1, 'cup'),
(31, 54, 4, 'tbsp'),
(31, 45, 3, 'cloves'),
(31, 61, 2, 'tbsp'),
(31, 58, 0.5, 'cup'),
(31, 85, 0.5, 'cup'),
(31, 41, 0.5, 'tsp');
