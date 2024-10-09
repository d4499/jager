package jobapplication

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/d4499/jager/internal"
	"github.com/d4499/jager/internal/auth"
	"github.com/d4499/jager/internal/database/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v5/pgtype"
)

type JobApplicationRoutes struct {
	svc  JobApplicationService
	auth auth.AuthService
}

func NewJobApplicationRoutes(svc JobApplicationService, auth auth.AuthService) *JobApplicationRoutes {
	return &JobApplicationRoutes{
		svc:  svc,
		auth: auth,
	}
}

func (j *JobApplicationRoutes) Register(r *chi.Mux) {
	r.Route("/api/jobapplications", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(j.auth.SessionMiddleware(&j.auth))
			r.Post("/", j.handleCreateJobApplication)
			r.Get("/", j.handleGetAllJobApplications)
			r.Delete("/", j.handleDeleteJobApplication)
		})
	})
}

type CreateJobApplicationRequest struct {
	Title       string    `json:"title"`
	Company     string    `json:"company"`
	AppliedDate time.Time `json:"applied_date"`
}

func (j *JobApplicationRoutes) handleCreateJobApplication(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("session").(db.Session)
	if !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	var createJobApplicationRequest CreateJobApplicationRequest

	if err := json.NewDecoder(r.Body).Decode(&createJobApplicationRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	jobApp, err := j.svc.CreateJobApplication(db.CreateJobApplicationParams{
		ID:      internal.NewCUID(),
		Title:   createJobApplicationRequest.Title,
		Company: createJobApplicationRequest.Company,
		UserID:  session.UserID,
		AppliedDate: pgtype.Timestamp{
			Time:  createJobApplicationRequest.AppliedDate,
			Valid: true,
		},
	})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	render.JSON(w, r, jobApp)
}

func (j *JobApplicationRoutes) handleGetAllJobApplications(w http.ResponseWriter, r *http.Request) {
	session, ok := r.Context().Value("session").(db.Session)
	if !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	jobApps, err := j.svc.GetAllJobApplications(session.UserID)
	if err != nil {
		log.Printf("Unable to get job applications: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, jobApps)
}

type DeleteJobApplicationRequest struct {
	UserId string `json:"userId"`
}

func (j *JobApplicationRoutes) handleDeleteJobApplication(w http.ResponseWriter, r *http.Request) {
	_, ok := r.Context().Value("session").(db.Session)
	if !ok {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}

	var deleteJobApplicationRequest DeleteJobApplicationRequest

	if err := json.NewDecoder(r.Body).Decode(&deleteJobApplicationRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	err := j.svc.DeleteJobApplication(deleteJobApplicationRequest.UserId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusForbidden)
		return
	}

	render.JSON(w, r, "deleted job application")
}
