# Software Application in GO (Golang)

# VIVERO

Development of a server-side software application that provides an API to consume a series of services for a nursery.

The application simulates the behavior and growth of plants as they are fed and cared for in different ways. This allows for experimental testing without affecting actual plants. Each plant has unique properties and reacts differently to varying conditions.

## Plant Properties

Each plant has the following properties:

| Property | Description |
| --- | --- |
| Id | Plant identifier |
| Name | Plant name |
| DegreeSurvival | Represents the tolerance range around the plant's ideal value. It is calculated as a percentage relative to each subsequent item, allowing the plant to survive for a certain period of time. If this limit is exceeded, the plant dies. Unit: percentage (%) |
| AmountWaterRequired | Represents the amount of water required for the plant to have 100% water in the system and be fully satisfied. Unit: liters (l) |
| AmountWaterSystem | A value ranging from 0 to 100%, where the value decreases over time. This value decreases by 1/X for each unit of time elapsed, where X is the plant's hydration level. The 100% level is reached when the required amount of water is met. |
| DegreeHydration | A value between 20 and 100, where a lower value indicates faster dehydration of the plant, and a higher value indicates slower dehydration, allowing the plant to resist without requiring water for a longer time.
| AmountNutrientsRequired | Represents the amount of nutrients required for the plant to have 100% nutrients in the system and be fully satisfied. Unit: milligrams (mg) |
| AmountNutrientsSystem | A value ranging from 0 to 100%, where the value decreases over time. This value decreases by 1/Y for each unit of time elapsed, where Y is the plant's nutrition level. The 100% level is reached when the required amount of nutrients is met.
| DegreeNutrition | A value between 60 and 100, where a lower value indicates faster nutrient loss, and a higher value indicates slower nutrient loss, allowing the plant to resist without requiring nutrients for a longer time.
| Created | Date of the creation.
| LastUpdate | Date of the last update.

# Possible cases 

## Case 1:
In this case, the plant has an adequate supply of water and nutrients, but no additional water or nutrients are added. Every second, the plant loses 0.02% of water and 0.016% of nutrients. After 5000 seconds (approximately 83.3 minutes or 1.38 hours), the plant's water level reaches 0%. However, due to its 20% survival factor, the plant can withstand a little longer without water, acting as if it still has 20% of water in its system. With double the rate of loss, it would have 500 seconds (approximately 8.3 minutes) before the plant dies due to water depletion.

## Case 2:
In this case, an additional 30% of water is added to the plant, even though it already has 100% of its required amount (which is equivalent to 0.2 liters). The plant has a 20% survival factor. Therefore, it has 500 seconds (approximately 8.3 minutes) to eliminate the excess water and reach 100%. However, the added excess amount is 0.06 liters, which would actually take 750 seconds (approximately 12.5 minutes) to remove. Consequently, the plant fails to survive due to excessive water quantity.


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`DB_HOST`: host database

`DB_USER` : user database

`DB_PASSWORD`: password database

`DB_PORT`: database port

The credentials are for a postgres database.
You can find the sql in the root of the project.
