package postgres

import (
    "context"
    "fmt"
    "time"

    sq "github.com/Masterminds/squirrel"

    "github.com/018bf/example/pkg/log"

    "github.com/018bf/example/internal/domain/models"
    "github.com/018bf/example/internal/domain/repositories"

    "github.com/018bf/example/internal/domain/errs"
    "github.com/018bf/example/pkg/postgresql"
    "github.com/018bf/example/pkg/utils"
    "github.com/jmoiron/sqlx"
)


type DayDTO struct {
	ID          string         `db:"id,omitempty"`
    Name string `db:"name"`
    Repeat int `db:"repeat"`
    EquipmentID string `db:"equipment_id"`
	UpdatedAt   time.Time      `db:"updated_at,omitempty"`
	CreatedAt   time.Time      `db:"created_at,omitempty"`
}

func NewDayDTOFromModel(day *models.Day) *DayDTO {
	dto := &DayDTO{
		ID:          string(day.ID),
        Name: string(day.Name),
        Repeat: int(day.Repeat),
        EquipmentID: string(day.EquipmentID),
		UpdatedAt:   day.UpdatedAt,
		CreatedAt:   day.CreatedAt,
	}
	return dto
}

func (dto *DayDTO) ToModel() *models.Day {
	model := &models.Day{
		ID:          models.UUID(dto.ID),
        Name: string(dto.Name),
        Repeat: int(dto.Repeat),
        EquipmentID: string(dto.EquipmentID),
		UpdatedAt:   dto.UpdatedAt,
		CreatedAt:   dto.CreatedAt,
	}
	return model
}

type DayListDTO []*DayDTO

func (list DayListDTO) ToModels() []*models.Day {
    listDays := make([]*models.Day, len(list))
	for i := range list {
        listDays[i] = list[i].ToModel()
	}
	return listDays
}

type DayRepository struct {
    database *sqlx.DB
    logger   log.Logger
}

func NewDayRepository(
    database *sqlx.DB,
    logger log.Logger,
) repositories.DayRepository {
    return &DayRepository{
        database: database,
        logger:   logger,
    }
}

func (r *DayRepository) Create(
    ctx context.Context,
    day *models.Day,
) error {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    dto := NewDayDTOFromModel(day)
    q := sq.Insert("public.days").
        Columns(
            "name",
            "repeat",
            "equipment_id",
            "updated_at",
            "created_at",
        ).
        Values(
            dto.Name,
            dto.Repeat,
            dto.EquipmentID,
            dto.UpdatedAt,
            dto.CreatedAt,
        ).
        Suffix("RETURNING id")
    query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
    if err := r.database.QueryRowxContext(ctx, query, args...).StructScan(dto); err != nil {
        e := errs.FromPostgresError(err)
        return e
    }
    day.ID = models.UUID(dto.ID)
    return nil
}

func (r *DayRepository) Get(
    ctx context.Context,
    id models.UUID,
) (*models.Day, error ) {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    dto := &DayDTO{}
    q := sq.Select(
        "days.id",
        "days.name",
        "days.repeat",
        "days.equipment_id",
        "days.updated_at",
        "days.created_at",
    ).
        From("public.days").
        Where(sq.Eq{"id": id}).
        Limit(1)
    query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
    if err := r.database.GetContext(ctx, dto, query, args...); err != nil {
        e := errs.FromPostgresError(err).
            WithParam("day_id", string(id))
        return nil, e
    }
    return dto.ToModel(), nil
}

func (r *DayRepository) List(
    ctx context.Context,
    filter *models.DayFilter,
) ([]*models.Day, error ) {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    var dto DayListDTO
    const pageSize = uint64(10)
    if filter.PageSize == nil {
        filter.PageSize = utils.Pointer(pageSize)
    }
    q := sq.Select(
        "days.id",
        "days.name",
        "days.repeat",
        "days.equipment_id",
        "days.updated_at",
        "days.created_at",
    ).
        From("public.days").
        Limit(pageSize)
    if filter.Search != nil {
        q = q.Where(postgresql.Search{
            Lang:   "english",
            Query:  *filter.Search,
            Fields: []string{
                "name",
            },
        })
    }
    if filter.PageNumber != nil && *filter.PageNumber > 1 {
        q = q.Offset((*filter.PageNumber - 1) * *filter.PageSize)
    }
    q = q.Limit(*filter.PageSize)
    if len(filter.OrderBy) > 0 {
        q = q.OrderBy(filter.OrderBy...)
    }
    query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
    if err := r.database.SelectContext(ctx, &dto, query, args...); err != nil {
        e := errs.FromPostgresError(err)
        return nil, e
    }
    return dto.ToModels(), nil
}

func (r *DayRepository) Update(
    ctx context.Context,
    day *models.Day,
) error {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    dto := NewDayDTOFromModel(day)
    q := sq.Update("public.days").
        Where(sq.Eq{"id": day.ID}).
        Set("days.name", dto.Name).
        Set("days.repeat", dto.Repeat).
        Set("days.equipment_id", dto.EquipmentID).
        Set("updated_at", day.UpdatedAt)
    query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
    result, err := r.database.ExecContext(ctx, query, args...)
    if err != nil {
        e := errs.FromPostgresError(err).
            WithParam("day_id", fmt.Sprint(day.ID))
        return e
    }
    affected, err := result.RowsAffected()
    if err != nil {
        return errs.FromPostgresError(err).
            WithParam("day_id", fmt.Sprint(day.ID))
    }
    if affected == 0 {
        e := errs.NewEntityNotFound().
            WithParam("day_id", fmt.Sprint(day.ID))
        return e
    }
    return nil
}

func (r *DayRepository) Delete(
    ctx context.Context,
    id models.UUID,
) error {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    q := sq.Delete("public.days").Where(sq.Eq{"id": id})
    query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
    result, err := r.database.ExecContext(ctx, query, args...)
    if err != nil {
        e := errs.FromPostgresError(err).
            WithParam("day_id", fmt.Sprint(id))
        return e
    }
    affected, err := result.RowsAffected()
    if err != nil {
        e := errs.FromPostgresError(err).
            WithParam("day_id", fmt.Sprint(id))
        return e
    }
    if affected == 0 {
        e := errs.NewEntityNotFound().
            WithParam("day_id", fmt.Sprint(id))
        return e
    }
    return nil
}

func (r *DayRepository) Count(
    ctx context.Context,
    filter *models.DayFilter,
) (uint64, error ) {
    ctx, cancel := context.WithTimeout(ctx, time.Second)
    defer cancel()
    q := sq.Select("count(id)").From("public.days")
    query, args := q.PlaceholderFormat(sq.Dollar).MustSql()
    result := r.database.QueryRowxContext(ctx, query, args...)
    if err := result.Err(); err != nil {
        e := errs.FromPostgresError(err)
        return 0, e
    }
    var count uint64
    if err := result.Scan(&count); err != nil {
        e := errs.FromPostgresError(err)
        return 0, e
    }
    return count, nil
}
