package __

import (
	model "movieReviewSystem/movieReviewSystemApi/review/reviewModel"
)

func ReviewToModelReview(review *Review) *model.Review {
	return &model.Review{
		ID:               review.ReviewId,
		UserId:           review.UserId,
		BaseId:           review.BaseId,
		RootId:           review.RootId,
		Text:             review.Text,
		Status:           review.Status,
		Level:            review.Level,
		RootCommentCount: review.RootCommentCount,
		CreateAt:         review.CreateAt,
		UpdateAt:         review.UpdateAt,
		HeadId:           review.HeadId,
	}
}
func ModelReviewToReview(review *model.Review) *Review {
	return &Review{
		ReviewId:         review.ID,
		UserId:           review.UserId,
		BaseId:           review.BaseId,
		RootId:           review.RootId,
		Text:             review.Text,
		Status:           review.Status,
		Level:            review.Level,
		RootCommentCount: review.RootCommentCount,
		CreateAt:         review.CreateAt,
		UpdateAt:         review.UpdateAt,
		HeadId:           review.HeadId,
	}
}
