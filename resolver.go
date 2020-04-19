package gorecipe

import (
    "context"
    "github.com/kofoworola/gorecipe/models"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Resolver struct{
     DB *gorm.DB
}

func (r *Resolver) Mutation() MutationResolver {
    return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
    return &queryResolver{r}
}

//Resolver for mutations
type mutationResolver struct{ *Resolver }

//Create recipe mutation
func (r *mutationResolver) CreateItem(ctx context.Context, input *NewItem, ingredients []*NewIngredient) (*models.Item, error) {
    //Fetch Connection and close db
    db := models.FetchConnection()
    defer db.Close()

    //Create the recipe using the input structs
    recipe := models.Item{Name: input.Name, Procedure: *input.Procedure}

    //initialize the ingredients with the length of the input for ingredients
    recipe.Ingredients = make([]models.Ingredient,len(ingredients))
    //Loop and add all items
    for index,item := range ingredients{
        recipe.Ingredients[index] = models.Ingredient{Name: item.Name}
    }
    //Create by passing the pointer to the recipe
    db.Create(&recipe)
    return &recipe, nil
}

//Update recipe mutation
func (r *mutationResolver) UpdateItem(ctx context.Context, id *int, input *NewItem, ingredients []*NewIngredient) (*models.Item, error) {
    //Fetch Connection and close db
    db := models.FetchConnection()
    defer db.Close()

    var recipe models.Item
    //Find recipe based on ID and update
    db = db.Preload("Ingredients").Where("id = ?",*id).First(&recipe).Update("name",input.Name)
    if input.Procedure != nil{
        db.Update("procedure",*input.Procedure)
    }

    //Update Ingredients
    recipe.Ingredients = make([]models.Ingredient,len(ingredients))
    for index,item := range ingredients{
        recipe.Ingredients[index] = models.Ingredient{Name:item.Name}
    }
    db.Save(&recipe)
    return &recipe,nil
}

//Delete recipe mutation
func (r *mutationResolver) DeleteItem(ctx context.Context, id *int) ([]*models.Item, error) {
    //Fetch connection
    db := models.FetchConnection()
    defer db.Close()
    var recipe models.Item

    //Fetch based on ID and delete
    db.Where("id = ?",*id).First(&recipe).Delete(&recipe)

    //Preload and fetch all recipes
    var recipes []*models.Item
    db.Preload("Ingredients").Find(&recipes)
    return recipes,nil
}


//Query resolver
type queryResolver struct{ *Resolver }

//Get all recipes
func (r *queryResolver) Items(ctx context.Context) ([]*models.Item, error) {
    //Fetch a connection
    db := models.FetchConnection()
    //Defer closing the database
    defer db.Close()
    //Create an array of recipes to populate

    db.Preload("Ingredients").Find(&recipes)
    return recipes,nil
}
