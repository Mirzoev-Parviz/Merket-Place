package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type ReviewService struct {
	repo repository.Review
}

func NewReviewService(repo repository.Review) *ReviewService {
	return &ReviewService{repo: repo}
}

func (r *ReviewService) CreateReview(review models.Review) error {
	return r.repo.CreateReview(review)
}
func (r *ReviewService) CalculateProductRating(productID int) error {
	return r.repo.CalculateProductRating(productID)
}

func (r *ReviewService) GetReview(id int) (review models.Review, err error) {
	return r.repo.GetReview(id)
}

func (r *ReviewService) GetAllReviews() (reviews []models.Review, err error) {
	return r.repo.GetAllReviews()
}
func (r *ReviewService) GetMerchantProductReview(merchantID, merchantProductID int) (reviews []models.Review, err error) {
	return r.repo.GetMerchantProductReview(merchantID, merchantProductID)
}
func (r *ReviewService) UpdateReview(input models.Review) error {
	return r.repo.UpdateReview(input)
}

func (r *ReviewService) DeleteReview(id, userID int) error {
	return r.repo.DeleteReview(id, userID)
}
