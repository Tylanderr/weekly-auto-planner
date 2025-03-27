package consts

type GroceryCategory struct {
	Name       string
	ItemsSlice []string
}

var Produce = GroceryCategory{
	Name:       "Produce",
	ItemsSlice: []string{"Apples", "Bananas", "Oranges", "Grapes", "Strawberries", "Blueberries", "Raspberries", "Broccoli", "Carrots", "Spinach", "Lettuce", "Tomatoes", "Potatoes", "Onions", "Garlic", "Bell Peppers", "Cucumbers", "Zucchini", "Avocados", "Mushrooms", "Salad"},
}

var MeatAndPoultry = GroceryCategory{
	Name:       "MeatAndPoultry",
	ItemsSlice: []string{"Chicken Breast", "Chicken Thighs", "Ground Beef", "Steak", "Pork Chops", "Sausage", "Bacon", "Turkey", "Ham", "Roast Beef", "Lamb Chops", "Duck", "Ground Turkey", "Pork Tenderloin", "Chicken Wings", "Filet Mignon"},
}

var Seafood = GroceryCategory{
	Name:       "Seafood",
	ItemsSlice: []string{"Salmon", "Tuna", "Shrimp", "Cod", "Tilapia", "Crab", "Lobster", "Oysters", "Mussels", "Clams", "Scallops", "Trout", "Catfish", "Swordfish", "Anchovies"},
}

var Dairy = GroceryCategory{
	Name:       "Dairy",
	ItemsSlice: []string{"Milk", "Cheese", "Yogurt", "Butter", "Eggs", "Cream", "Sour Cream", "Cottage Cheese", "Ice Cream", "Whipped Cream", "Cream Cheese", "Half-and-Half", "Mozzarella", "Cheddar", "Parmesan", "Provolone"},
}

var Bakery = GroceryCategory{
	Name:       "Bakery",
	ItemsSlice: []string{"Bread", "Rolls", "Bagels", "Croissants", "Muffins", "Cookies", "Cakes", "Pies", "Donuts", "Pastries", "Brownies", "Scones", "Biscuits", "Loaf Cakes", "Cupcakes"},
}

var FrozenFoods = GroceryCategory{
	Name:       "FrozenFoods",
	ItemsSlice: []string{"Frozen Vegetables", "Frozen Fruits", "Frozen Pizza", "Frozen Meals", "Ice Cream", "Frozen Yogurt", "Frozen Waffles", "Frozen Pancakes", "Frozen Chicken Nuggets", "Frozen Fish", "Frozen Pies", "Frozen Desserts", "Frozen Appetizers", "Frozen Smoothies", "Frozen Juice Concentrates"},
}

var PantryStaples = GroceryCategory{
	Name:       "PantryStaples",
	ItemsSlice: []string{"Rice", "Pasta", "Fettuccini noodles", "Spaghetti Noodles", "Spaghetti Sauce", "Alfredo Sauce", "Canned Vegetables", "Canned Fruits", "Beans", "Lentils", "Oats", "Flour", "Sugar", "Salt", "Pepper", "Cooking Oil", "Vinegar", "Canned Soups", "Broth", "Spices", "Herbs", "Cereals", "Peanut Butter", "Jelly"},
}

var Beverages = GroceryCategory{
	Name:       "Beverages",
	ItemsSlice: []string{"Water", "Juice", "Soda", "Coffee", "Tea", "Sports Drinks", "Energy Drinks", "Milk Alternatives", "Iced Tea", "Lemonade", "Sparkling Water", "Coconut Water", "Vegetable Juice", "Smoothies", "Hot Chocolate"},
}

var Snacks = GroceryCategory{
	Name:       "Snacks",
	ItemsSlice: []string{"Chips", "Pretzels", "Popcorn", "Crackers", "Nuts", "Seeds", "Trail Mix", "Candy", "Chocolate", "Granola Bars", "Fruit Snacks", "Yogurt Tubes", "Jerky", "Rice Cakes", "Pudding Cups"},
}

var HouseholdGoods = GroceryCategory{
	Name:       "HouseholdGoods",
	ItemsSlice: []string{"Cleaning Supplies", "Laundry Detergent", "Dish Soap", "Paper Towels", "Toilet Paper", "Trash Bags", "Light Bulbs", "Batteries", "Air Fresheners", "Sponges", "Aluminum Foil", "Plastic Wrap", "Food Storage Containers", "Cleaning Wipes", "Dishwasher Detergent"},
}

var PersonalCare = GroceryCategory{
	Name:       "PersonalCare",
	ItemsSlice: []string{"Shampoo", "Conditioner", "Soap", "Toothpaste", "Toothbrushes", "Deodorant", "Lotion", "Sunscreen", "Cosmetics", "Feminine Hygiene Products", "Hair Products", "Shaving Cream", "Razors", "Vitamins", "Supplements"},
}

var InternationalFoods = GroceryCategory{
	Name:       "InternationalFoods",
	ItemsSlice: []string{"Soy Sauce", "Teriyaki Sauce", "Salsa", "Taco Shells", "Rice Noodles", "Curry Paste", "Olive Oil", "Balsamic Vinegar", "Pita Bread", "Hummus", "Kimchi", "Sriracha", "Wasabi", "Seaweed", "Tortillas"},
}

var Deli = GroceryCategory{
	Name:       "Deli",
	ItemsSlice: []string{"Deli Meat", "Deli Cheese", "Sandwiches", "Salad", "Olives", "Pickles", "Prepared Meals", "Hummus", "Tzatziki", "Pasta Salad", "Potato Salad", "Coleslaw", "Quiche", "Soups", "Chicken Salad"},
}

var Floral = GroceryCategory{
	Name:       "Floral",
	ItemsSlice: []string{"Roses", "Lilies", "Tulips", "Carnations", "Sunflowers", "Orchids", "Daisies", "Chrysanthemums", "Gerberas", "Hydrangeas", "Greenery", "Potted Plants", "Bouquets", "Arrangements", "Seasonal Flowers"},
}
