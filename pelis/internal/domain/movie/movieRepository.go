package movie

import (

	"gorm.io/gorm"
)



type MovieRepository struct{
	Db *gorm.DB
}



func(m *MovieRepository) GetAllMovies()([]Movie, error){
	var item []Movie
	e := m.Db.Find(&item)
	if(e.Error != nil){
		return nil, e.Error
	}
	return item, nil
}

func (m *MovieRepository) GetById(id int)(Movie, error){
	var movie Movie
	err := m.Db.First(&movie, id)
	if(err.Error != nil){
		return movie, err.Error
	}
	return movie, nil

}

func (m *MovieRepository) Save(movie *Movie)(error){
	e := m.Db.Create(movie)
	if(e.Error != nil){
		return e.Error
	}
	return nil
	
}



