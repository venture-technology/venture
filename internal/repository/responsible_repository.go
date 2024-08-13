package repository

type IResponsibleRepository interface {
	Create()
	Get()
	Update()
	Delete()
	SaveCard()
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(db *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: db,
	}
}

func (rr *ResponsibleRepository) Create() {

}

func (rr *ResponsibleRepository) Get() {
	
}

func (rr *ResponsibleRepository) Update() {
	
}

func (rr *ResponsibleRepository) Delete() {
	
}

func (rr *ResponsibleRepository) SaveCard() {
	
}