package controller

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/joho/godotenv"
)

type ImageData struct {
	Image string `json:"image"`
}

type Analysis struct {
	LabelModelVersion string `json:"labelmodelversion"`
	Labels            []struct {
		Confidence float64 `json:"confidence"`
		Name       string  `json:"name"`
		Parents    []struct {
			Name string `json:"name"`
		}
	}
}

func ImageRecogniser(image string) Analysis {
	godotenv.Load()

	sess := session.New(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AccessKeyID"),
			os.Getenv("SecretAccessKey"),
			"",
		),
	})
	svc := rekognition.New(sess)
	decodedImage, err := base64.StdEncoding.DecodeString(image)
	if err != nil {
		fmt.Println(err)
	}

	input := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{
			Bytes: decodedImage,
		},
	}

	result, err := svc.DetectLabels(input)
	if err != nil {
		fmt.Println(err)
	}
	output, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(output)
	var responseData Analysis
	if err := json.Unmarshal(output, &responseData); err != nil {
		panic(err)
	}
	return responseData
}

func ImageController(w http.ResponseWriter, r *http.Request) {
	var imageData ImageData

	err := json.NewDecoder(r.Body).Decode(&imageData)
	if err != nil {
		fmt.Println(err, "Error in reading payload")
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var receivedData = ImageRecogniser(imageData.Image)
	json.NewEncoder(w).Encode(receivedData)
}
