package order_service

import (
	"fmt"
	"net/http"
	"sync"
	"testing"

	"github.com/shakhboznorbekov/mytasks/golang_crud/models"

	"github.com/bxcodec/faker/v3"
	"github.com/test-go/testify/assert"
)

func TestActor(t *testing.T) {
	s = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			id := CreateActor(t)
			DeleteActor(t, id)
		}()

	}

	wg.Wait()

	fmt.Println("s: ", s)
}

func CreateActor(t *testing.T) string {
	response := &models.Actor{}

	request := &models.CreateActor{
		First_name: faker.Name(),
		Last_name:  faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPost, "/actor", request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 201)
	}

	fmt.Println(response)

	return response.Id
}

func UpdateActor(t *testing.T, id string) string {
	response := &models.Actor{}
	request := &models.UpdateActor{
		First_name: faker.Name(),
		Last_name:  faker.Name(),
	}

	resp, err := PerformRequest(http.MethodPut, "/actor/"+id, request, response)

	assert.NoError(t, err)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 200)
	}

	fmt.Println(resp)

	return response.Id
}

func DeleteActor(t *testing.T, id string) string {

	resp, _ := PerformRequest(
		http.MethodDelete,
		fmt.Sprintf("/actor/%s", id),
		nil,
		nil,
	)

	assert.NotNil(t, resp)

	if resp != nil {
		assert.Equal(t, resp.StatusCode, 204)
	}

	return ""
}
