# Mining Post
An API-based multiplayer game where you purchase mining equipment and various resources to become rich.

## Instructions
Use the REST API to buy and sell commodities, equipment, land, and employees. Every world update will pay out your salary and trigger any special effects of resources you hold. The objective of the game is simple: become the richest miner in the world. By playing the stock market through buying and selling commodities, acquiring land and equipment, and expanding your corporate empire, you will grow your net worth and climb to the top of the charts!

## World updates
All world updates to market prices, player inventory, etc., happen at the same time. This is based on an internal timer that triggers on a regular cadence. During a world update, everything is locked and nothing can be purchased or sold.

During a world update:
- Prices in the marketplace are updated.
- Player inventory is updated based on the player's resources in their inventory.
- Player is paid their salary.
- Player net worth is calculated based on the player's current money and resources.

## Types of resources
There are different types of resources that can be purchased on the market:
- Commodity (Examples: Granite, Limestone, Gold, Sapphire)
- Equipment (Examples: Pickaxe, Mining Cart, Dynamite)
- Land (Examples: Claim, Mine)
- Employee (Examples: Worker, Specialist, Surveyor, Mining Engineer)

**Commodities** are raw resources and can be purchased or sold on the market. Selling these commodities is how miners make their money. To make money, you can buy low and sell high on the market, dig up commodities from the ground to sell, or run a full-scale mining operation to unearth the treasures hidden below.

**Equipment** covers anything that can assist in a mining operation. This can range from the lowly pickaxe all the way up to gargantuan dump trucks and refineries. Equipment is usually required to build a mine or increase the output of your mining operations.

**Land** is an area that can provide mining opportunities. Land falls under two categories: a claim or a mine. Prospecting allows you to discover new claims. Different types of claims can affect the rocks and minerals you dig up. A mine can be built on a claim. Once a mine is built, it will passively produce. Various types of equipment can increase the production and efficiency of your mines.

**Employee** is a category of laborers. This ranges from unskilled workers to highly specialized mining engineers. Hiring employees is necessary to build mines.

## Actions
In addition to money used to purchase resources, players also have action points they can spend. Action points are allocated based on player rank. Every world update refills the player's action points to the maximum for their rank and they cannot hold more than the maximum. Action points can be used to:
- Dig for rocks and minerals.
- Prospect for land.

**Dig** is an action that produces rocks and minerals from your claims. Unlike claims, mines are automated and do not need to be manually worked. The more action points you spend to dig in a single attempt, the better your results will be. Various equipment can improve your digging results.

**Prospect** is an action that allows you to discover new land. Discovered land is always a claim. Claims are necessary to dig or build mines, so prospecting is important. Various equipment can improve your prospecting results.

## Ranks
Players all start as amateurs, but can progress up through the ranks by increasing their net worth. Rank promotions come with a new title, higher salary, and more available action points. Higher ranks may also have additional requirements for promotion.

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
    "resource": "Limestone",
    "quantity": 2
}

200 OK
{
    "cost": "$19.04",
    "message": "Successfully purchased 2 of resource: Limestone, total cost: $19.04"
}
```

Sell resource:
```
POST /market/sell
{
    "player": "tstark",
    "resource": "Limestone",
    "quantity": 2
}

200 OK
{
    "profit": "$25.58",
    "message": "Successfully sold 2 of resource: Limestone, total profit: $25.58"
}
```
