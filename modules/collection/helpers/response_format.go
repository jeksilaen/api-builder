package helpers

import (
	"github.com/jeksilaen/api-builder/modules/collection/models"
)

func ReturnFailedCreateResponse(message string) *models.FailedResponse {
	return &models.FailedResponse{
		Error:   "Create collection failed",
		Message: message,
		Links: []models.Link{
			{
				Rel:  "Create collection",
				Href: "/collection/v1",
			},
		},
	}
}

func ReturnFailedUpdateResponse(message string) *models.FailedResponse {
	return &models.FailedResponse{
		Error:   "Update collection failed",
		Message: message,
		Links: []models.Link{
			{
				Rel:  "Create collection",
				Href: "/collection/v1",
			},
		},
	}
}

func ReturnSucessCreateResponse(createdCollection *models.Collection) *models.SucessCreateResponse {
	return &models.SucessCreateResponse{
		Message: "Created Collection sucessfully",
		Data: models.CollectionResponse{
			ID:     createdCollection.ID,
			UserID: createdCollection.UserID,
			Name:   createdCollection.Name,
		},
		Links: []models.Link{
			{
				Rel:  "Create request",
				Href: "/request/v1",
			},
		},
	}
}

func ReturnSucessUpdateResponse(createdCollection *models.Collection) *models.SucessCreateResponse {
	return &models.SucessCreateResponse{
		Message: "Updated Collection sucessfully",
		Data: models.CollectionResponse{
			ID:     createdCollection.ID,
			UserID: createdCollection.UserID,
			Name:   createdCollection.Name,
		},
		Links: []models.Link{
			{
				Rel:  "Create request",
				Href: "/request/v1",
			},
		},
	}
}

func ReturnSucessGetResponse(collections []*models.Collection) *models.SucessGetResponse {
	var collectionResponses []models.CollectionResponse

	for _, collection := range collections {
		collectionResponses = append(collectionResponses, models.CollectionResponse{
			ID:     collection.ID,
			UserID: collection.UserID,
			Name:   collection.Name,
		})
	}

	return &models.SucessGetResponse{
		Message: "Get Collection sucessfully",
		Data:    collectionResponses,
		Links: []models.Link{
			{
				Rel:  "Get request",
				Href: "/request/collection/v1/:collection_id",
			},
		},
	}
}

func ReturnSucessDeleteResponse(message string) *models.SucessDeleteResponse {
	return &models.SucessDeleteResponse{
		Message: message,
		Links: []models.Link{
			{
				Rel:  "Create collection",
				Href: "/collection/v1",
			},
		},
	}
}
