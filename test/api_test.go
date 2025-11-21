package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	base := getBaseURL()
	t.Logf("Testing API at: %s", base)

	//=====================================================================

	t.Log("1. Create question test")
	var createdQ QuestionResponse
	doPost(
		t,
		fmt.Sprintf("%s/questions", base),
		CreateQuestionRequest{Text: "What is Golang?"},
		201,
		&createdQ,
	)
	require.NotZero(t, createdQ.ID)
	t.Logf("Created question with ID: %d", createdQ.ID)

	//=====================================================================

	t.Log("2. Questions list test")
	var listQ []QuestionResponse
	doGet(
		t,
		fmt.Sprintf("%s/questions", base),
		200,
		&listQ,
	)
	require.NotEmpty(t, listQ)
	t.Logf("Recieved %d questions", len(listQ))

	//=====================================================================

	t.Log("3. Get question test")
	var recievedQ QuestionResponse
	doGet(
		t,
		fmt.Sprintf("%s/questions/%d", base, createdQ.ID),
		200,
		&recievedQ,
	)
	require.Equal(t, "What is Golang?", recievedQ.Text)
	t.Log("Successfully recieved question")

	//=====================================================================

	t.Log("4. Create answer test")
	var createdAns AnswerResponse
	doPost(
		t,
		fmt.Sprintf("%s/questions/%d/answers", base, createdQ.ID),
		CreateAnswerRequest{
			UserID: "user-123",
			Text:   "Golang is a fast language",
		},
		201,
		&createdAns,
	)
	require.NotZero(t, createdAns.ID)
	t.Logf("Created answer with ID: %d", createdAns.ID)

	//=====================================================================

	t.Log("5. Get answer test")
	var recievedAns AnswerResponse
	doGet(
		t,
		fmt.Sprintf("%s/answers/%d", base, createdAns.ID),
		200,
		&recievedAns,
	)
	require.Equal(t, "Golang is a fast language", recievedAns.Text)
	t.Log("Successfully recieved answer")

	//=====================================================================

	t.Log("6. Delete answer test")
	doDelete(
		t,
		fmt.Sprintf("%s/answers/%d", base, createdAns.ID),
		204,
	)
	t.Log("Successfully deleted answer")

	//=====================================================================

	t.Log("7. Delete question test")
	doDelete(
		t,
		fmt.Sprintf("%s/questions/%d", base, createdQ.ID),
		204,
	)
	t.Log("Successfully deleted question")

	//=====================================================================

	t.Log("8. Deleted question test")
	doGet[any](
		t,
		fmt.Sprintf("%s/questions/%d", base, createdQ.ID),
		404,
		nil,
	)
	t.Log("Question not found")

	//=====================================================================

	t.Log("ALL TESTS PASSED!")
}
