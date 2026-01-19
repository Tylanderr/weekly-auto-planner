package consts

type GroceryCategory struct {
	Name       string
	ItemsSlice []string
}

var Produce = GroceryCategory{
	Name:       "Produce",
	ItemsSlice: []string{"apples", "bananas", "oranges", "grapes", "strawberries", "blueberries", "raspberries", "broccoli", "carrots", "spinach", "lettuce", "tomatoes", "potatoes", "onions", "sweet vidalia onion", "yellow onion", "green onion", "garlic", "bell peppers", "cucumbers", "zucchini", "avocados", "mushrooms", "salad"},
}

var MeatAndPoultry = GroceryCategory{
	Name:       "MeatAndPoultry",
	ItemsSlice: []string{"chicken breast", "chicken thighs", "ground beef", "steak", "pork chops", "sausage", "bacon", "turkey", "ham", "roast beef", "lamb chops", "duck", "ground turkey", "pork tenderloin", "chicken wings", "filet mignon"},
}

var Seafood = GroceryCategory{
	Name:       "Seafood",
	ItemsSlice: []string{"salmon", "tuna", "shrimp", "cod", "tilapia", "crab", "lobster", "oysters", "mussels", "clams", "scallops", "trout", "catfish", "swordfish", "anchovies"},
}

var Dairy = GroceryCategory{
	Name:       "Dairy",
	ItemsSlice: []string{"milk", "cheese", "yogurt", "butter", "eggs", "cream", "sour cream", "cottage cheese", "ice cream", "whipped cream", "cream cheese", "half-and-half", "mozzarella", "cheddar", "parmesan", "provolone", "sliced cheddar cheese"},
}

var Bakery = GroceryCategory{
	Name:       "Bakery",
	ItemsSlice: []string{"bread", "rolls", "bagels", "croissants", "muffins", "cookies", "cakes", "pies", "donuts", "pastries", "brownies", "scones", "biscuits", "loaf cakes", "cupcakes"},
}

var FrozenFoods = GroceryCategory{
	Name:       "FrozenFoods",
	ItemsSlice: []string{"frozen vegetables", "frozen fruits", "frozen pizza", "frozen meals", "ice cream", "frozen yogurt", "frozen waffles", "frozen pancakes", "frozen chicken nuggets", "frozen fish", "frozen pies", "frozen desserts", "frozen appetizers", "frozen smoothies", "frozen juice concentrates"},
}

var PantryStaples = GroceryCategory{
	Name:       "PantryStaples",
	ItemsSlice: []string{"alfredo sauce", "rice", "pasta", "fettuccini noodles", "spaghetti noodles", "spaghetti sauce", "canned vegetables", "canned fruits", "beans", "lentils", "oats", "flour", "sugar", "salt", "pepper", "cooking oil", "vinegar", "canned soups", "broth", "spices", "herbs", "cereals", "peanut butter", "jelly"},
}

var Beverages = GroceryCategory{
	Name:       "Beverages",
	ItemsSlice: []string{"water", "juice", "soda", "coffee", "tea", "sports drinks", "energy drinks", "milk alternatives", "iced tea", "lemonade", "sparkling water", "coconut water", "vegetable juice", "smoothies", "hot chocolate"},
}

var Snacks = GroceryCategory{
	Name:       "Snacks",
	ItemsSlice: []string{"chips", "pretzels", "popcorn", "crackers", "nuts", "seeds", "trail mix", "candy", "chocolate", "granola bars", "fruit snacks", "yogurt tubes", "jerky", "rice cakes", "pudding cups"},
}

var HouseholdGoods = GroceryCategory{
	Name:       "HouseholdGoods",
	ItemsSlice: []string{"cleaning supplies", "laundry detergent", "dish soap", "paper towels", "toilet paper", "trash bags", "light bulbs", "batteries", "air fresheners", "sponges", "aluminum foil", "plastic wrap", "food storage containers", "cleaning wipes", "dishwasher detergent"},
}

var PersonalCare = GroceryCategory{
	Name:       "PersonalCare",
	ItemsSlice: []string{"shampoo", "conditioner", "soap", "toothpaste", "toothbrushes", "deodorant", "lotion", "sunscreen", "cosmetics", "feminine hygiene products", "hair products", "shaving cream", "razors", "vitamins", "supplements"},
}

var InternationalFoods = GroceryCategory{
	Name:       "InternationalFoods",
	ItemsSlice: []string{"soy sauce", "teriyaki sauce", "salsa", "taco shells", "rice noodles", "curry paste", "olive oil", "balsamic vinegar", "pita bread", "hummus", "kimchi", "sriracha", "wasabi", "seaweed", "tortillas"},
}

var Deli = GroceryCategory{
	Name:       "Deli",
	ItemsSlice: []string{"deli meat", "deli cheese", "sandwiches", "olives", "pickles", "prepared meals", "hummus", "tzatziki", "pasta salad", "potato salad", "coleslaw", "quiche", "soups", "chicken salad"},
}
