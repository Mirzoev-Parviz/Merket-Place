package repository

import (
	"market_place/config"
	"market_place/models"
	"math"

	"gorm.io/gorm"
)

type ReviewPostgres struct {
	db *gorm.DB
}

func NewReviewPostgres(db *gorm.DB) *ReviewPostgres {
	return &ReviewPostgres{db: db}
}

func (m *ReviewPostgres) CreateReview(review models.Review) error {
	if review.Rating > 5 {
		review.Rating = 5
	}

	if err := config.DB.Create(&review).Error; err != nil {
		return err
	}

	if err := m.CalculateProductRating(review.ProductID); err != nil {
		return err
	}

	return nil
}

func (r *ReviewPostgres) GetReview(id int) (review models.Review, err error) {
	err = config.DB.Where("id = ?").First(&review).Error
	if err != nil {
		return models.Review{}, err
	}

	return review, nil
}

func (r *ReviewPostgres) GetMerchantProductReview(merchantID, merchantProductID int) (reviews []models.Review, err error) {
	err = config.DB.Where("merhant_id = ? AND merchant_product_id = ?",
		merchantID, merchantProductID).Find(reviews).Error
	if err != nil {
		return []models.Review{}, err
	}

	return reviews, nil
}

func (r *ReviewPostgres) GetAllReviews() (reviews []models.Review, err error) {
	if err = config.DB.Find(&reviews).Error; err != nil {
		return []models.Review{}, err
	}

	return reviews, nil
}

func (r *ReviewPostgres) UpdateReview(input models.Review) error {
	return config.DB.Model(&models.Review{}).
		Updates(&input).Where("id = ?", input.ID).Error
}

func (r *ReviewPostgres) DeleteReview(id, userID int) error {
	return config.DB.Model(&models.Review{}).Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Review{}).Error

}

func (r *ReviewPostgres) CalculateProductRating(productID int) error {
	var totalRating float64
	var count int64
	if err := config.DB.Model(&models.Review{}).
		Where("product_id = ?", productID).
		Select("COALESCE(AVG(rating), 0)").
		Count(&count).Row().Scan(&totalRating); err != nil {
		return err
	}
	overallRating := math.Min(totalRating, 5.0)
	return config.DB.Model(&models.MerchantProduct{}).
		Where("id = ?", productID).
		UpdateColumn("rating", overallRating).Error
}
