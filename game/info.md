# Mining Post
An API-based multiplayer game where you purchase mining equipment and various resources to become rich.

## Market updates
All game world updates to stock prices, player inventory, etc., happen at the same time. This is based on an internal timer that triggers on a regular cadence. During a game world update, everything is locked and nothing can be purchased or sold.

During a game world update:
- Prices in the marketplace are updated.
- Player inventory is updated based on the player's items in their inventory.
- Player net worth is updated.

## Types of stock
- Commodity (Granite, Limestone, Gold, Sapphire, etc.)
- Equipment (Pickaxe, Mining Cart, Dynamite)
- Land (Claim, Mine)
- Employee (Worker, Specialist, Mining Engineer)

## REST APIs
TODO apis
