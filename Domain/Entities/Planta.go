package Entities

type Plant struct {
	Id                           uint
	Nombre                       string
	GradoSupervivencia           float64
	CantidadAguaRequerida        float64
	CantidadAguaSistema          float64
	GradoHidratacion             uint
	CantidadNutrientesRequeridos float64
	CantidadNutrientesSistema    float64
	GradoNutricion               int
}

type IPlantaRepositorio interface {
	GetPlantas() []Plant
	GetPlant(id uint) Plant
	NewPlant(Plant Plant) bool
	UpdatePlant(Plant Plant) bool
	DeletePlant(id uint) bool
}

// implementar interfaz
func (s Plant) GetPlants() []Plant {
	plants := make([]Plant, 0)
	//students = metodoParaLeerBd
	return plants
}

func (s Plant) GetPlant(id uint) Plant {
	var plant Plant
	//students = metodoParaLeerBd
	return plant
}

func (s Plant) NewPlant(Plant Plant) bool {
	//var creada = metodoGuardarDB()
	//return creada
	return false
}
