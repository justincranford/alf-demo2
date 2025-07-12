# Texas Recipes API - Copilot Instructions

## Project Overview
This project implements a REST API for managing Texas recipes, following the OpenAPI 3.0 specification and built with Go.

## Constraints
1. **Texas-Only Recipes**
   - All recipes must be Texas-style or Texas-inspired cuisine
   - Valid categories: Tex-Mex, Texas BBQ, Chili (no beans), Southern Texas, Cowboy/Ranch-style, Gulf Coast

2. **Recipe Requirements**
   - Must include all required fields: id, name, ingredients, instructions
   - Should specify Texas cuisine type in the cuisine field
   - Should include realistic preparation and cooking times

## Code Guidelines
1. **API Implementation**
   - Follow RESTful principles for all endpoints
   - Implement proper error handling and status codes
   - Use idiomatic Go code with proper package structure
   - Ensure endpoints match the OpenAPI specification

2. **Data Validation**
   - Validate all incoming recipe data
   - Ensure recipes contain Texas-specific ingredients and techniques
   - Reject non-Texas recipes with appropriate error messages

## Example Recipe Structure
```json
{
  "id": "tx001",
  "name": "Texas-Style Beef Brisket",
  "ingredients": [
    "1 whole beef brisket (10-12 pounds)",
    "1/4 cup kosher salt",
    "1/4 cup black pepper",
    "2 tablespoons garlic powder",
    "2 tablespoons onion powder",
    "1 tablespoon cayenne pepper"
  ],
  "instructions": "1. Mix all spices together to create a rub. 2. Trim excess fat from brisket, leaving about 1/4 inch fat cap. 3. Apply rub generously to all sides of brisket. 4. Let sit for 1 hour at room temperature. 5. Prepare smoker to 225°F using oak or hickory wood. 6. Smoke brisket fat side up for 6 hours. 7. Wrap in butcher paper and continue cooking until internal temperature reaches 203°F (about 6 more hours). 8. Rest for 1-2 hours before slicing against the grain.",
  "cuisine": "Texas BBQ",
  "prepTime": 60,
  "cookTime": 720
}
```

## Texas Recipe Types to Consider
1. **Tex-Mex**
   - Fajitas, enchiladas, tacos, queso, chili con carne

2. **Texas BBQ**
   - Brisket, beef ribs, sausage, smoked turkey

3. **Texas Chili**
   - Texas Red (no beans), competition chili

4. **Southern Texas**
   - Chicken fried steak, cornbread, pecan pie

5. **Cowboy/Ranch-style**
   - Dutch oven recipes, chuck wagon meals

6. **Gulf Coast**
   - Gulf shrimp, redfish, oysters

## Implementation Notes
- Server runs on port 3000
- Swagger UI available at /swagger/
- OpenAPI specification defines all endpoints
- In-memory storage currently implemented
