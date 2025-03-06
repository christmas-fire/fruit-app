package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/christmas-fire/fruit-app/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) getAllFruits(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://www.fruityvice.com/api/fruit/all", nil)
	if err != nil {
		logrus.Errorf("error creating request: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := fetchFruits(c, &h.Client, req)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		logrus.Errorf("unexpected status code: %d", res.StatusCode)
		c.JSON(http.StatusBadGateway, gin.H{"error": "external service returned an unexpected status code"})
		return
	}

	var fruits []models.Fruit
	if err := json.NewDecoder(res.Body).Decode(&fruits); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode JSON response"})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

func (h *Handler) getFruitsByName(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		logrus.Error("error: query param 'name' not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param not found"})
		return
	}

	query := fmt.Sprintf("https://www.fruityvice.com/api/fruit/%s", name)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", query, nil)
	if err != nil {
		logrus.Errorf("error creating request: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := fetchFruits(c, &h.Client, req)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			logrus.Errorf("name not found: %d", res.StatusCode)
			c.JSON(http.StatusNotFound, gin.H{"error": "name not found"})
			return
		}
		logrus.Errorf("unexpected status code: %d", res.StatusCode)
		c.JSON(http.StatusBadGateway, gin.H{"error": "external service returned an unexpected status code"})
		return
	}

	var fruit models.Fruit
	if err := json.NewDecoder(res.Body).Decode(&fruit); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode JSON response"})
		return
	}

	c.JSON(http.StatusOK, fruit)
}

func (h *Handler) getFruitsByFamily(c *gin.Context) {
	family := c.Query("family")
	if family == "" {
		logrus.Error("error: query param 'family' not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param not found"})
		return
	}

	query := fmt.Sprintf("https://www.fruityvice.com/api/fruit/family/%s", family)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", query, nil)
	if err != nil {
		logrus.Errorf("error creating request: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := fetchFruits(c, &h.Client, req)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			logrus.Errorf("family not found: %d", res.StatusCode)
			c.JSON(http.StatusNotFound, gin.H{"error": "family not found"})
			return
		}
		logrus.Errorf("unexpected status code: %d", res.StatusCode)
		c.JSON(http.StatusBadGateway, gin.H{"error": "external service returned an unexpected status code"})
		return
	}

	var fruits []models.Fruit
	if err := json.NewDecoder(res.Body).Decode(&fruits); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode JSON response"})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

func (h *Handler) getFruitsByGenus(c *gin.Context) {
	genus := c.Query("genus")
	if genus == "" {
		logrus.Error("error: query param 'genus' not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param not found"})
		return
	}

	query := fmt.Sprintf("https://www.fruityvice.com/api/fruit/genus/%s", genus)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", query, nil)
	if err != nil {
		logrus.Errorf("error creating request: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := fetchFruits(c, &h.Client, req)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			logrus.Errorf("genus not found: %d", res.StatusCode)
			c.JSON(http.StatusNotFound, gin.H{"error": "genus not found"})
			return
		}
		logrus.Errorf("unexpected status code: %d", res.StatusCode)
		c.JSON(http.StatusBadGateway, gin.H{"error": "external service returned an unexpected status code"})
		return
	}

	var fruits []models.Fruit
	if err := json.NewDecoder(res.Body).Decode(&fruits); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode JSON response"})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

func (h *Handler) getFruitsByOrder(c *gin.Context) {
	order := c.Query("order")
	if order == "" {
		logrus.Error("error: query param 'order' not found")
		c.JSON(http.StatusBadRequest, gin.H{"error": "query param not found"})
		return
	}

	query := fmt.Sprintf("https://www.fruityvice.com/api/fruit/order/%s", order)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", query, nil)
	if err != nil {
		logrus.Errorf("error creating request: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	res := fetchFruits(c, &h.Client, req)
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		if res.StatusCode == http.StatusNotFound {
			logrus.Errorf("order not found: %d", res.StatusCode)
			c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
			return
		}
		logrus.Errorf("unexpected status code: %d", res.StatusCode)
		c.JSON(http.StatusBadGateway, gin.H{"error": "external service returned an unexpected status code"})
		return
	}

	var fruits []models.Fruit
	if err := json.NewDecoder(res.Body).Decode(&fruits); err != nil {
		logrus.Errorf("error decoding JSON: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode JSON response"})
		return
	}

	c.JSON(http.StatusOK, fruits)
}

func fetchFruits(c *gin.Context, client *http.Client, r *http.Request) *http.Response {
	res, err := client.Do(r)
	if err != nil {
		logrus.Errorf("error making request to Fruityvice API: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from external service"})
		return nil
	}

	return res
}
