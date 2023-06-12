# Software Application in GO (Golang)

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
| LastUpdate | Date of the last update.

## Posibles casos

Case 1:

Case 2:

Case 3:
