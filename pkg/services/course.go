package services

import (
	"context"
	"net/http"

	"github.com/Bryan-BC/go-course-microservice/pkg/db"
	"github.com/Bryan-BC/go-course-microservice/pkg/models"
	"github.com/Bryan-BC/go-course-microservice/pkg/pb"
)

type Server struct {
	DBPointer *db.DB
}

func (s *Server) GetCourse(ctx context.Context, req *pb.GetCourseRequest) (*pb.GetCourseResponse, error) {
	var course models.Course

	if result := s.DBPointer.DataBase.Where(&models.Course{Id: req.Id}).First(&course); result.Error != nil {
		return &pb.GetCourseResponse{
			Status: http.StatusNotFound,
			Error:  "Course not found",
		}, nil
	}

	return &pb.GetCourseResponse{
		Status: http.StatusOK,
		Course: &pb.Course{
			Id:          course.Id,
			Name:        course.Name,
			Description: course.Description,
			Schedule:    course.Schedule,
		},
	}, nil
}

func (s *Server) CreateCourse(ctx context.Context, req *pb.CreateCourseRequest) (*pb.CreateCourseResponse, error) {
	var course models.Course

	if result := s.DBPointer.DataBase.Where(&models.Course{Id: req.Course.Id}).First(&course); result.Error == nil {
		return &pb.CreateCourseResponse{
			Status: http.StatusConflict,
			Error:  "Course already exists",
		}, nil
	}

	course.Id = req.Course.Id
	course.Name = req.Course.Name
	course.Description = req.Course.Description
	course.Schedule = req.Course.Schedule

	if result := s.DBPointer.DataBase.Create(&course); result.Error != nil {
		return &pb.CreateCourseResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error creating course",
		}, nil
	}

	return &pb.CreateCourseResponse{
		Status: http.StatusCreated,
		Course: &pb.Course{
			Id:          course.Id,
			Name:        course.Name,
			Description: course.Description,
			Schedule:    course.Schedule,
		},
	}, nil
}

func (s *Server) DeleteCourse(ctx context.Context, req *pb.DeleteCourseRequest) (*pb.DeleteCourseResponse, error) {
	var course models.Course

	if result := s.DBPointer.DataBase.Where(&models.Course{Id: req.Id}).First(&course); result.Error != nil {
		return &pb.DeleteCourseResponse{
			Status: http.StatusNotFound,
			Error:  "Course not found",
		}, nil
	}

	if result := s.DBPointer.DataBase.Delete(&course); result.Error != nil {
		return &pb.DeleteCourseResponse{
			Status: http.StatusInternalServerError,
			Error:  "Error deleting course",
		}, nil
	}

	return &pb.DeleteCourseResponse{
		Status: http.StatusNoContent,
	}, nil
}
