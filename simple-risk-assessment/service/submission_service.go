package service

import (
	"context"
	"fmt"
	"gin/simple-risk-assessment/entity"
)

// ISubmissionService mendefinisikan interface untuk layanan submission
type ISubmissionService interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
}

// ISubmissionRepository mendefinisikan interface untuk repository submission
type ISubmissionRepository interface {
	CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error)
	GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error)
	UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error)
	DeleteSubmission(ctx context.Context, id int) error
	GetAllSubmissions(ctx context.Context) ([]entity.Submission, error)
}

// submissionService adalah implementasi dari ISubmissionService yang menggunakan ISubmissionRepository
type submissionService struct {
	submissionRepo ISubmissionRepository
}

// NewSubmissionService membuat instance baru dari submissionService
func NewSubmissionService(submissionRepo ISubmissionRepository) ISubmissionService {
	return &submissionService{submissionRepo: submissionRepo}
}

// CreateSubmission membuat submission baru
func (s *submissionService) CreateSubmission(ctx context.Context, submission *entity.Submission) (entity.Submission, error) {
	// Memanggil CreateSubmission dari repository untuk membuat submission baru
	createdSubmission, err := s.submissionRepo.CreateSubmission(ctx, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("gagal membuat submission: %v", err)
	}
	return createdSubmission, nil
}

// GetSubmissionByID mendapatkan submission berdasarkan ID
func (s *submissionService) GetSubmissionByID(ctx context.Context, id int) (entity.Submission, error) {
	// Memanggil GetSubmissionByID dari repository untuk mendapatkan submission berdasarkan ID
	submission, err := s.submissionRepo.GetSubmissionByID(ctx, id)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("gagal mendapatkan submission berdasarkan ID: %v", err)
	}
	return submission, nil
}

// UpdateSubmission memperbarui data submission
func (s *submissionService) UpdateSubmission(ctx context.Context, id int, submission entity.Submission) (entity.Submission, error) {
	// Memanggil UpdateSubmission dari repository untuk memperbarui data submission
	updatedSubmission, err := s.submissionRepo.UpdateSubmission(ctx, id, submission)
	if err != nil {
		return entity.Submission{}, fmt.Errorf("gagal memperbarui submission: %v", err)
	}
	return updatedSubmission, nil
}

// DeleteSubmission menghapus submission berdasarkan ID
func (s *submissionService) DeleteSubmission(ctx context.Context, id int) error {
	// Memanggil DeleteSubmission dari repository untuk menghapus submission berdasarkan ID
	err := s.submissionRepo.DeleteSubmission(ctx, id)
	if err != nil {
		return fmt.Errorf("gagal menghapus submission: %v", err)
	}
	return nil
}

// GetAllSubmissions mendapatkan semua submission
func (s *submissionService) GetAllSubmissions(ctx context.Context) ([]entity.Submission, error) {
	// Memanggil GetAllSubmissions dari repository untuk mendapatkan semua submission
	submissions, err := s.submissionRepo.GetAllSubmissions(ctx)
	if err != nil {
		return nil, fmt.Errorf("gagal mendapatkan semua submission: %v", err)
	}
	return submissions, nil
}
