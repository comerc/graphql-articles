package project

import (
	"context"
	"errors"
	"github.com/acelot/articles/internal/go-pg/condition"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

const projectNameConstraint string = "project_name_key"

var ProjectNameConstraintError = errors.New("project name constraint error")
var NoRowsAffectedError = errors.New("no rows affected error")

type Repository struct {
	primaryDB   *pg.DB
	secondaryDB *pg.DB
}

type FindFilterIDAnyOf []uuid.UUID
type FindOrderByName bool
type FindOrderByCreatedAt bool
type FindOrderByModifiedAt bool

func (uuids FindFilterIDAnyOf) Apply(query *pg.Query) {
	query.Where(`id = ANY(?)`, pg.Array(uuids))
}

func (isDesc FindOrderByName) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("name " + dir)
}

func (isDesc FindOrderByCreatedAt) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("created_at " + dir)
}

func (isDesc FindOrderByModifiedAt) Apply(query *pg.Query) {
	dir := "ASC"
	if isDesc {
		dir = "DESC"
	}

	query.Order("modified_at " + dir)
}

func NewRepository(primaryDB *pg.DB, secondaryDB *pg.DB) *Repository {
	return &Repository{
		primaryDB,
		secondaryDB,
	}
}

func (r *Repository) Find(ctx context.Context, conditions ...condition.Condition) ([]Project, error) {
	var items []Project

	query := r.secondaryDB.ModelContext(ctx, &items)

	condition.Apply(query, conditions...)

	err := query.Select()

	return items, err
}

func (r *Repository) FindOneByID(ctx context.Context, id uuid.UUID) (*Project, error) {
	entities, err := r.Find(ctx, FindFilterIDAnyOf{id})
	if err != nil {
		return nil, err
	}

	if len(entities) == 0 {
		return nil, nil
	}

	return &entities[0], nil
}

func (r *Repository) Count(ctx context.Context, conditions ...condition.Condition) (int, error) {
	query := r.secondaryDB.ModelContext(ctx, (*Project)(nil))

	condition.Apply(query, conditions...)

	return query.Count()
}

func (r *Repository) Insert(ctx context.Context, entity Project) error {
	_, err := r.primaryDB.ModelContext(ctx, &entity).Insert()

	return specifyError(err)
}

func (r *Repository) Update(ctx context.Context, entity Project) error {
	currentVersion := entity.Version
	entity.Version++

	res, err := r.primaryDB.ModelContext(ctx, &entity).
		WherePK().
		Where("version = ?", currentVersion).
		Update()

	if err != nil {
		return specifyError(err)
	}

	if res.RowsAffected() == 0 {
		return NoRowsAffectedError
	}

	return nil
}

func specifyError(err error) error {
	pgErr, ok := err.(pg.Error)
	if !ok {
		return err
	}

	constraint := pgErr.Field([]byte("n")[0])

	if constraint == projectNameConstraint {
		return ProjectNameConstraintError
	}

	return err
}