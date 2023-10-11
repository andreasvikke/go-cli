package confluence

import (
	"confluence-poc/src/helpers"
	"confluence-poc/src/models"
	"net/url"
	"strconv"
	"strings"
	"sync"
)

// GetContentByID retrieves content from the API by ID.
//
// It takes in the following parameters:
// - ids: a slice of strings representing the IDs of the content to retrieve.
// - query: a ContentQuery struct representing the query parameters for retrieving the content.
//
// It returns a *models.ContentResult pointer, which contains the retrieved content and any errors encountered.
func (a *API) GetContentByID(ids []string, query models.ContentQuery) *models.ContentResult {
	var wg sync.WaitGroup
	uniqueIDs := helpers.UniqueString(ids)

	results := make(chan *models.Content, len(uniqueIDs))
	errors := make(chan *models.ContentError, len(uniqueIDs))

	wg.Add(len(uniqueIDs))
	for _, id := range uniqueIDs {
		go a.getContentByID(id, query, results, errors, &wg)
	}

	go func() {
		wg.Wait()
		close(errors)
		close(results)
	}()

	output := &models.ContentResult{
		Data:   make([]models.Content, len(results)),
		Errors: make([]models.ContentError, len(errors)),
	}
	for val := range results {
		output.Data = append(output.Data, *val)
	}
	for err := range errors {
		output.Errors = append(output.Errors, *err)
	}

	return output
}

// getContentByID retrieves content by ID from the API.
//
// Parameters:
// - id: the ID of the content to retrieve.
// - query: the content query parameters.
// - output: a channel to send the retrieved content to.
// - errors: a channel to send any error result to.
// - wg: a wait group to synchronize the goroutines.
//
// Returns nothing. But it sends the retrieved content to the output channel and any errors encountered to the errors channel.
func (a *API) getContentByID(id string, query models.ContentQuery, output chan<- *models.Content, errors chan<- *models.ContentError, wg *sync.WaitGroup) {
	defer wg.Done()

	endpoint, err := url.ParseRequestURI(a.endPoint.String() + "/content/" + id)
	if err != nil {
		errors <- newContentError(id, query.Version, err.Error())
		return
	}
	endpoint.RawQuery = addContentQueryParams(query).Encode()

	content, err := a.SendContentRequest(endpoint, "GET", nil)
	if err != nil {
		errors <- newContentError(id, query.Version, err.Error())
		return
	}

	output <- content
}

// SetContent sets the content of the API.
//
// Parameters:
// - contents: a slice of `models.Content` representing the contents to set.
//
// Returns:
// - *models.ContentResult: a pointer to the `models.ContentResult` containing the results of the operation.
func (a *API) SetContent(contents []models.Content) *models.ContentResult {
	var wg sync.WaitGroup

	results := make(chan *models.Content, len(contents))
	errors := make(chan *models.ContentError, len(contents))

	wg.Add(len(contents))
	for _, c := range contents {
		go a.setContent(c, results, errors, &wg)
	}

	go func() {
		wg.Wait()
		close(errors)
		close(results)
	}()

	output := &models.ContentResult{
		Data:   make([]models.Content, len(results)),
		Errors: make([]models.ContentError, len(errors)),
	}
	for val := range results {
		output.Data = append(output.Data, *val)
	}
	for err := range errors {
		output.Errors = append(output.Errors, *err)
	}

	return output
}

func (a *API) setContent(c models.Content, output chan<- *models.Content, errors chan<- *models.ContentError, wg *sync.WaitGroup) {
	defer wg.Done()

	query := models.ContentQuery{
		Expand: []string{"version"},
	}

	endpoint, err := url.ParseRequestURI(a.endPoint.String() + "/content/" + c.Id)
	if err != nil {
		errors <- newContentError(c.Id, query.Version, err.Error())
		return
	}
	endpoint.RawQuery = addContentQueryParams(query).Encode()

	currentContent, err := a.SendContentRequest(endpoint, "GET", nil)
	if err != nil {
		errors <- newContentError(c.Id, query.Version, err.Error())
		return
	}

	c.Version = &models.Version{
		Number: currentContent.Version.Number + 1,
	}

	content, err := a.SendContentRequest(endpoint, "PUT", &c)
	if err != nil {
		errors <- newContentError(c.Id, query.Version, err.Error())
		return
	}

	output <- content
}

// addContentQueryParams generates URL query parameters based on the provided ContentQuery.
//
// Parameters:
// - query: the ContentQuery used to generate the query parameters.
//
// Returns a pointer to a url.Values containing the generated query parameters.
func addContentQueryParams(query models.ContentQuery) *url.Values {
	data := url.Values{}
	if len(query.Expand) != 0 {
		data.Set("expand", strings.Join(query.Expand, ","))
	}
	//get specific version
	if query.Version != 0 {
		data.Set("version", strconv.Itoa(query.Version))
	}
	return &data
}

// newContentError creates a new instance of the ContentError struct.
//
// Parameters:
//   - id: the ID of the content error.
//   - version: the version of the content error.
//   - errorMessage: the error message of the content error.
//
// Return:
//   - *models.ContentError: a pointer to the newly created ContentError struct.
func newContentError(id string, version int, errorMessage string) *models.ContentError {
	return &models.ContentError{
		Id:           id,
		Version:      version,
		ErrorMessage: errorMessage,
	}
}
