# Mining Post
An API-based multiplayer game where you purchase mining equipment and various resources to become rich.

## Instructions
Use the REST API to buy and sell commodities, equipment, land, and employees. Every world update will pay out your salary and trigger any special effects of items you hold. The objective of the game is simple: become the richest miner in the world. By playing the stock market through buying and selling commodities, acquiring land and equipment, and expanding your corporate empire, you will grow your net worth and climb to the top of the charts!

## World updates
All world updates to market prices, player inventory, etc., happen at the same time. This is based on an internal timer that triggers on a regular cadence. During a world update, everything is locked and nothing can be purchased or sold.

During a world update:
- Prices in the marketplace are updated.
- Player inventory is updated based on the player's items in their inventory.
- Player is paid their salary.
- Player net worth is updated.

## Types of resources
There are different types of resources that can be purchased on the market:
- Commodity (Granite, Limestone, Gold, Sapphire, etc.)
- Equipment (Pickaxe, Mining Cart, Dynamite)
- Land (Claim, Mine)
- Employee (Worker, Specialist, Mining Engineer)

## REST APIs
Here are some examples of the REST API in action. This is not a comprehensive list of all REST APIs.

Retrieve player inventory:
```
GET /player/tstark/inventory

200 OK
{
    "player": {
        "name": "tstark",
        "title": "Amateur Miner L1",
        "netWorth": "$116.18",
        "money": "$100.96",
        "salary": "$10.00",
        "inventory": [
            {
                "name": "Limestone",
                "description": "Limestone is a sedimentary rock composed mainly of calcium carbonate, often used in construction for its durability.",
                "type": "Commodity",
                "quantity": 2
            }
        ]
    }
}
```

Purchase resource:
```
POST /market/buy
{
    "player": "tstark",
    "item": "Limestone",
    "quantity": 2
}

200 OK
{
    "cost": 19.04,
    "message": "Successfully purchased 2 of item: Limestone, total cost: $19.04"
}
```

Sell resource:
```
POST /market/sell
{
    "player": "tstark",
    "item": "Limestone",
    "quantity": 2
}

200 OK
{
    "profit": 25.58,
    "message": "Successfully sold 2 of item: Limestone, total profit: $25.58"
}
```
