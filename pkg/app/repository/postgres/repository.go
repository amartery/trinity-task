package postgres

import (
	"database/sql"

	"github.com/amartery/trinity-task/pkg/app"
	"github.com/amartery/trinity-task/pkg/app/models"
)

type AppRepository struct {
	con *sql.DB
}

func NewAppRepository(con *sql.DB) app.Repository {
	return &AppRepository{
		con: con,
	}
}

func (r *AppRepository) GetTodayActivityLikes() ([]models.User, error) {
	// select distinct * from Users as u join likes as l on (u.id = l.user_id) where l.create_at >= 'today';
	query := `SELECT u.id, u.name FROM Users AS u
			   JOIN Likes AS l ON (u.id = l.user_id)
			   WHERE l.create_at >= 'today';`

	rows, err := r.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.User, 0)
	for rows.Next() {
		tmp := models.User{}
		if err := rows.Scan(
			&tmp.Id,
			&tmp.Name,
		); err != nil {
			return nil, err
		}

		result = append(result, tmp)
	}
	return result, nil
}

func (r *AppRepository) GetTodayActivityComments() ([]models.User, error) {
	// select distinct * from Users as u join comments as c on (u.id = c.user_id) where c.create_at >= 'today';
	query := `SELECT u.id, u.name FROM Users AS u
			   JOIN Comments AS c ON (u.id = c.user_id)
			   WHERE c.create_at >= 'today';`

	rows, err := r.con.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.User, 0)
	for rows.Next() {
		tmp := models.User{}
		if err := rows.Scan(
			&tmp.Id,
			&tmp.Name,
		); err != nil {
			return nil, err
		}

		result = append(result, tmp)
	}
	return result, nil
}

func (r *AppRepository) AddUser(u *models.User) error {
	query := `INSERT INTO Users (name) 
			  VALUES ($1) 
	          RETURNING id;`
	return r.con.QueryRow(query, u.Name).Scan(&u.Id)
}

func (r *AppRepository) AddComment(u *models.CommentRequest) error {
	query := `INSERT INTO Comments (video_id, text, user_id, create_at) 
			  VALUES ($1, $2, $3, $4) 
	          RETURNING id;`
	return r.con.QueryRow(query, u.VideoId, u.Text, u.UserId, u.CreateAt).Scan(&u.Id)
}

func (r *AppRepository) AddLike(u *models.Like) error {
	query := `INSERT INTO Likes (video_id, user_id, create_at) 
			  VALUES ($1, $2, $3);`
	_, err := r.con.Exec(query, u.VideoId, u.UserId, u.CreateAt)
	return err
}
