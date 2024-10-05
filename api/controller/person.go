package controller

import (
	"fmt"
	"log"
	"net/http"

	errapi "github.com/beka-birhanu/GoCrudChallange/api/error"
	"github.com/beka-birhanu/GoCrudChallange/app/common/cqrs"
	personcmd "github.com/beka-birhanu/GoCrudChallange/app/person/commands"
	"github.com/beka-birhanu/GoCrudChallange/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PersonController struct {
	BaseController
	createHandler cqrs.Handler[*personcmd.CreatePersonCommand, *models.Person]
	updateHandler cqrs.Handler[*personcmd.UpdatePersonCommand, *models.Person]
	deleteHandler cqrs.Handler[uuid.UUID, bool]
	getHandler    cqrs.Handler[uuid.UUID, *models.Person]
	getAllHandler cqrs.Handler[struct{}, []*models.Person]
}

type NewPersonControllerParams struct {
	CreateHandler cqrs.Handler[*personcmd.CreatePersonCommand, *models.Person]
	UpdateHandler cqrs.Handler[*personcmd.UpdatePersonCommand, *models.Person]
	DeleteHandler cqrs.Handler[uuid.UUID, bool]
	GetHandler    cqrs.Handler[uuid.UUID, *models.Person]
	GetAllHandler cqrs.Handler[struct{}, []*models.Person]
}

func NewPersonController(params NewPersonControllerParams) *PersonController {
	return &PersonController{
		createHandler: params.CreateHandler,
		updateHandler: params.UpdateHandler,
		deleteHandler: params.DeleteHandler,
		getHandler:    params.GetHandler,
		getAllHandler: params.GetAllHandler,
	}
}

func (c *PersonController) RegisterPublic(route *gin.RouterGroup) {
	route.GET("/person", c.getAll)
	route.POST("/person", c.create)
	route.GET("/person/:id", c.get)
	route.PUT("/person/:id", c.update)
	route.DELETE("/person/:id", c.delete)
}

func (c *PersonController) RegisterProtected(route *gin.RouterGroup) {}

func (c *PersonController) RegisterPrivileged(route *gin.RouterGroup) {}

func (c *PersonController) create(ctx *gin.Context) {
	var dto CreatePersonDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		c.RespondError(ctx, errapi.NewBadRequest(err.Error()))
		log.Printf("CreatePerson: Error binding JSON: %v", err)
		return
	}

	command := &personcmd.CreatePersonCommand{
		Name:    dto.Name,
		Age:     dto.Age,
		Hobbies: dto.Hobbies,
	}

	person, err := c.createHandler.Handle(command)
	if err != nil {
		c.RespondError(ctx, errapi.Map(err))
		log.Printf("CreatePerson: Error creating person: %v", err)
		return
	}

	response := PersonResponseDTO{
		ID:      person.ID(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}
	baseURL := fmt.Sprintf("http://%s", ctx.Request.Host)
	resourceLocation := fmt.Sprintf("%s%s/%s", baseURL, ctx.Request.URL.Path, person.ID().String())
	c.RespondWithLocation(ctx, http.StatusCreated, response, resourceLocation)
	log.Printf("CreatePerson: Successfully created person with ID: %s", person.ID())
}

func (c *PersonController) update(ctx *gin.Context) {
	var dto CreatePersonDTO
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.RespondError(ctx, errapi.NewBadRequest(err.Error()))
		log.Printf("UpdatePerson: Invalid ID format: %v", err)
		return
	}
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		c.RespondError(ctx, errapi.NewBadRequest(err.Error()))
		log.Printf("UpdatePerson: Error binding JSON: %v", err)
		return
	}

	command := &personcmd.UpdatePersonCommand{
		ID:      id,
		Name:    dto.Name,
		Age:     dto.Age,
		Hobbies: dto.Hobbies,
	}

	person, updateErr := c.updateHandler.Handle(command)
	if updateErr != nil {
		c.RespondError(ctx, errapi.Map(updateErr))
		log.Printf("UpdatePerson: Error updating person with ID %s: %v", id, updateErr)
		return
	}

	response := PersonResponseDTO{
		ID:      person.ID(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}

	c.Respond(ctx, http.StatusOK, response)
	log.Printf("UpdatePerson: Successfully updated person with ID: %s", person.ID())
}

func (c *PersonController) delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.RespondError(ctx, errapi.NewBadRequest("Invalid ID format"))
		log.Printf("DeletePerson: Invalid ID format: %v", err)
		return
	}

	_, deleteErr := c.deleteHandler.Handle(id)
	if deleteErr != nil {
		c.RespondError(ctx, errapi.Map(deleteErr))
		log.Printf("DeletePerson: Error deleting person with ID %s: %v", id, deleteErr)
		return
	}

	c.Respond(ctx, http.StatusNoContent, nil)
	log.Printf("DeletePerson: Successfully deleted person with ID: %s", id)
}

func (c *PersonController) get(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		c.RespondError(ctx, errapi.NewBadRequest("Invalid ID format"))
		log.Printf("GetPerson: Invalid ID format: %v", err)
		return
	}

	person, getErr := c.getHandler.Handle(id)
	if getErr != nil {
		c.RespondError(ctx, errapi.Map(getErr))
		log.Printf("GetPerson: Error retrieving person with ID %s: %v", id, getErr)
		return
	}

	response := PersonResponseDTO{
		ID:      person.ID(),
		Name:    person.Name(),
		Age:     person.Age(),
		Hobbies: person.Hobbies(),
	}

	c.Respond(ctx, http.StatusOK, response)
	log.Printf("GetPerson: Successfully retrieved person with ID: %s", person.ID())
}

func (c *PersonController) getAll(ctx *gin.Context) {
	persons, err := c.getAllHandler.Handle(struct{}{})
	if err != nil {
		c.RespondError(ctx, errapi.Map(err))
		log.Printf("GetAllPersons: Error retrieving all persons: %v", err)
		return
	}

	var responses []PersonResponseDTO
	for _, person := range persons {
		responses = append(responses, PersonResponseDTO{
			ID:      person.ID(),
			Name:    person.Name(),
			Age:     person.Age(),
			Hobbies: person.Hobbies(),
		})
	}

	c.Respond(ctx, http.StatusOK, responses)
	log.Printf("GetAllPersons: Successfully retrieved %d persons", len(responses))
}
