package repository

import (
	"database/sql"
	"errors"

	"github.com/LeUrok/DS-lab1/internal/model"
	"github.com/LeUrok/DS-lab1/internal/service"
)

type psqlPersonRep struct {
	db *sql.DB
}

func NewPersonRep (db *sql.DB) service.PersonRepository {
	return &psqlPersonRep{db: db}
}

func (r *psqlPersonRep) Create(req model.PersonRequest) (int32, error) {
	var id int32
	err := r.db.QueryRow(
		`insert into persons (name, age, address, work)
		values ($1, $2, $3, $4) returning id`,
		req.Name, req.Age, req.Address, req.Work,
	).Scan(&id)

	return id, err
}

func (r *psqlPersonRep) GetByID(id int32) (*model.Person, error) {
	p := &model.Person{}

	err := r.db.QueryRow(
		`select id, name, age, address, work from persons where id = $1`, 
		id,
	).Scan(&p.Id, &p.Name, &p.Age, &p.Address, &p.Work)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, service.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *psqlPersonRep) GetAll() ([]model.Person, error) {
	rows, err := r.db.Query(
		`select id, name, age, address, work from persons`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	persons := []model.Person{}
	
	for rows.Next() {
		var p model.Person
		err := rows.Scan(&p.Id, &p.Name, &p.Age, &p.Address, &p.Work)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, rows.Err()
}

func (r *psqlPersonRep) Update(id int32, req model.PersonRequest) (*model.Person, error) {
	p, err := r.GetByID(id)

	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		p.Name = req.Name
	}
	if req.Age != 0 {
		p.Age = req.Age
	}
	if req.Address != "" {
		p.Address = req.Address
	}
	if req.Work != "" {
		p.Work = req.Work
	}
	_, err = r.db.Exec(
		`update persons set name = $1, age = $2, address = $3, work = $4 where id = $5`,
		p.Name, p.Age, p.Address, p.Work, id,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *psqlPersonRep) Delete(id int32) error {
	res, err := r.db.Exec(`delete from persons where id = $1`, id)
	if err != nil {
		return err
	}
	n, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if n == 0 {
		return service.ErrNotFound
	}
	
	return nil
}